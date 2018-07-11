package handler

type BaseResp struct {
	Code int
	Msg string
}

type SetResp struct {
	BaseResp
	Data string
}

type GetResp struct {
	BaseResp
	Data string
}

type ExistResp struct {
	BaseResp
	Data bool
}

type KeysResp struct {
	BaseResp
	Data []string
}
