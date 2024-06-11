package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(result, paging, filter interface{}) *successRes {
	return &successRes{Data: result, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(result interface{}) *successRes {
	return &successRes{Data: result, Paging: nil, Filter: nil}
}
