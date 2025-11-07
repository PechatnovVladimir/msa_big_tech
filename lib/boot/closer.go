package boot

import "github.com/PechatnovVladimir/msa_big_tech/lib/closer"

func (app *App) Closer() (*closer.Closer, error) {
	cl := closer.New()
	app.Cl = cl
	return cl, nil
}
