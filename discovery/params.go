package discovery

type Params struct {
	Id     string
	Router router
}

type ParamsM struct {
	Id     string
	Module module
}

type router struct {
	Path        string `json:"path"`
	UrlSrc      string `json:"urlSrc"`
	UrlDest     string `json:"urlDest"`
	Method      string `json:"method"`
	ReturnType  string `json:"returnType"`
	Description string `json:"description"`
}

type module struct {
	Name     string   `json:"name"`
	HostName string   `json:"hostname"`
	Port     string   `json:"port"`
	Routers  []Router `json:"routers"`
}

type Param func(*Params)
type ParamM func(*ParamsM)

func newParams(param ...Param) Params {
	params := Params{}
	for _, p := range param {
		p(&params)
	}
	return params
}
func newParamsM(param ...ParamM) ParamsM {
	params := ParamsM{}
	for _, p := range param {
		p(&params)
	}
	return params
}
