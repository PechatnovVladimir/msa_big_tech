package closer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

var globalCloser *Closer

func init() {
	globalCloser = New(syscall.SIGINT, syscall.SIGTERM)
}

func Add(f ...CloseFunc) {
	globalCloser.Add(f...)
}

func CloseAll(ctx context.Context) error {
	return globalCloser.CloseAll(ctx)
}

type Closer struct {
	mu    sync.Mutex
	once  sync.Once
	funcs []CloseFunc
}

// CloseFunc - smth that close obj
type CloseFunc func(ctx context.Context) error

// New - returns closer
func New(sig ...os.Signal) *Closer {
	c := &Closer{}
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll(context.Background())
		}()
	}
	return c
}

// Add добавляем функцию, которую вызовем
// при получения сигнала о завершении работы
func (c *Closer) Add(f ...CloseFunc) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, f...)
}

// CloseAll закрывает все что передали в closer в обратном порядке
func (c *Closer) CloseAll(ctx context.Context) error {
	var closeErr error
	c.once.Do(func() {
		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		errs := make(chan error, len(funcs))

		// В обратном порядке и последовательно!
		for i := len(funcs) - 1; i >= 0; i-- {
			select {
			case <-ctx.Done():
			case errs <- funcs[i](ctx):
			}
		}

		msgs := make([]string, 0, len(funcs))
	Loop:
		for {
			select {
			case err, ok := <-errs:
				if !ok {
					break Loop
				}
				if err != nil {
					msgs = append(msgs, fmt.Sprintf("[!] %v", err))
				}
			case <-ctx.Done():
				break Loop
			}
		}

		if len(msgs) > 0 {
			closeErr = fmt.Errorf(
				"shutdown finished with error(s): \n%s",
				strings.Join(msgs, "\n"),
			)
		}
	})

	return closeErr
}
