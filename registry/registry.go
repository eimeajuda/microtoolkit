package registry

import (
	"sync"

	"github.com/DanielDanteDosSantosViana/microtoolkit/discovery"
	"github.com/DanielDanteDosSantosViana/microtoolkit/module"
	"github.com/DanielDanteDosSantosViana/microtoolkit/router"
	"github.com/fatih/color"
)

func Register(params ...module.ParamM) {
	module := module.NewModule(params...)
	results := make(chan *router.Router)
	routes := module.Params.Module.Routers

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(routes))

	var nameModule = module.Params.Module.Name

	for _, route := range routes {
		go func(route router.Router) {
			_, err := discovery.FindModule(nameModule)
			if err != nil {
				color.Red("Router : %s to module: %s - Error ('%s') ", route.Params.Router.Path, nameModule, err.Error())
			}

			/*
				_, err = discovery.FindRouter(nameModule, route.Params.Router.Path)
				if err != nil {
					color.Red("Router : %s to module: %s - Error ('%s') ", route.Params.Router.Path, nameModule, err.Error())
					return
				}

				log.Println("Router : " + route.Params.Router.Path + " to Module " + " - OK ( EXISTIS )")
				result, err := discovery.CreateRouter(nameModule, route)
				if err != nil {
					color.Red("Router : %s to module: %s - Error ('%s') ", route.Params.Router.Path, nameModule, err.Error())
				}

				if result.StatusCode != 200 {
					color.Red("Error to create Router : %s MSG: %s", route.Params.Router.Path, result.MSG)
				}*/
			color.Blue("Router : %s to module: %s - OK (new)", route.Params.Router.Path)

			waitGroup.Done()
		}(route)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

}

func NameModule(nameModule string) module.ParamM {
	return func(o *module.ParamsM) {
		o.Module.Name = nameModule
	}
}

func Routers(routers []router.Router) module.ParamM {
	return func(o *module.ParamsM) {
		o.Module.Routers = routers
	}
}
