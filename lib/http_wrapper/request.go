package http_wrapper

import (
	"net/http"
	"net/url"
)

type Request struct {
	Params map[string]string
	Header http.Header
	Query  url.Values
	Body   interface{}
}

func NewRequest(req *http.Request, body interface{}) (ret *Request) {
	ret = &Request{}
	ret.Query = req.URL.Query()
	ret.Header = req.Header
	ret.Body = body
	ret.Params = make(map[string]string)
	return
}
