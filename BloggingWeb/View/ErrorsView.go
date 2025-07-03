package view

type ErrResp struct {
	ErrMsg string
	Error  error
}

type SuccessResp struct {
	SuccessMsg string
	Response   interface{}
}
