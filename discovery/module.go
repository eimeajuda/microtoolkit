package discovery

type Module struct {
	Params ParamsM
	init   chan bool
}

func NewModule(param ...ParamM) Module {
	params := newParamsM(param...)
	return Module{params, make(chan bool)}
}
