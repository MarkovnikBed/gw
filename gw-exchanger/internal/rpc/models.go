package rpc

type response struct {
	Rates rates `json:"rates"`
}

type rates struct {
	USD float32 `json:"USD"`
	EUR float32 `json:"EUR"`
}
