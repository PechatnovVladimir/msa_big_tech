package social

import "time"

type Friend struct {
	FromUserID string
	ToUserID   string
	CreatedAt  time.Time
}
