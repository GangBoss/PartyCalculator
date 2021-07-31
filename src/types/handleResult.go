package types

type Requisites map[string]string
type HandleResult struct {
	Request     *Request
	Response    *Response
	Requisites  Requisites
	Error       error
	Continuable bool
}

func (result *HandleResult) Success() bool {
	return result.Error == nil
}

func (result *HandleResult) End() *HandleResult {
	result.Continuable = false
	return result
}

func (result *HandleResult) HandelError(err error) *HandleResult {
	result.Error = err
	result.Continuable = false
	return result
}

func NewHandleResult() *HandleResult {
	return &HandleResult{Requisites: make(map[string]string), Continuable: true}
}
