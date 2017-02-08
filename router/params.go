package router

type Params struct {
	Id     string
	Router RouterP
}

type RouterP struct {
	Path        string `json:"path"`
	UrlSrc      string `json:"urlSrc"`
	UrlDest     string `json:"urlDest"`
	Method      string `json:"method"`
	ReturnType  string `json:"returnType"`
	Description string `json:"description"`
}

type Param func(*Params)

func newParams(param ...Param) Params {
	params := Params{}
	for _, p := range param {
		p(&params)
	}
	return params
}
