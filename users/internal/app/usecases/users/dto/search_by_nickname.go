package dto

import "time"

type Query struct {
	IDs         []string
	Email       *string
	Nickname    *string
	CreatedFrom *time.Time
	CreatedTo   *time.Time
}

type SearchByNickname struct {
	Query Query
	Limit uint64
}
