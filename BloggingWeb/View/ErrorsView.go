package view

type ErrResp struct {
	ErrMsg string
	Error  interface{}
}

type SuccessResp struct {
	SuccessMsg string
	Response   interface{}
}
