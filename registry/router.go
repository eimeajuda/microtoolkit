package registry

type Router struct {
	params Params
	init   chan bool
}

func NewRouter(param ...Param) *Router {
	params := newParams(param...)
	return &Router{params, make(chan bool)}
}
