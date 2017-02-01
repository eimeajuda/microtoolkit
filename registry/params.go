package registry

type Params struct {
	Id     string
	router router
}

func newParams(param ...Param) Params {
	params := Params{}
	for _, p := range param {
		p(&params)
	}
	return params
}
