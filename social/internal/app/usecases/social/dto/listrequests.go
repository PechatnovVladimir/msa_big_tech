package dto

type ListRequestsIN struct {
	UserID string
}

type Request struct {
	RequestID string
	Status    StatusRequest
}

type ListRequestsOUT struct {
	Requests []Request
}
