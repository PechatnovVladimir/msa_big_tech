package dto

type StatusRequest int

const (
	UNKNOWN StatusRequest = iota
	PENDING
	ACCEPTED
	DECLINED
)
