package dto

type Status int

const (
	UNKNOWN Status = iota
	PENDING
	ACCEPTED
	DECLINED
)

type ListRequestsIN struct {
	UserID string
}

type Request struct {
	RequestID string
	Status    Status
}

type ListRequestsOUT struct {
	Requests []Request
}
