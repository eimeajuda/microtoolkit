package registry

type Router struct {
	params Params
	init   chan bool
}

func newRouter(param ...Param) *Router {
	params := newParams(param...)
	return &Router{params, make(chan bool)}
}
