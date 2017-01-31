package registry

type Params struct {
	Id     string
	router Router
}

type Router struct {
	Path        string `json:"path"`
	UrlSrc      string `json:"urlSrc"`
	UrlDest     string `json:"urlDest"`
	Method      string `json:"method"`
	ReturnType  string `json:"returnType"`
	Description string `json:"description"`
}
