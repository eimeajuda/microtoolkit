package registry

import (
	"sync"

	"github.com/DanielDanteDosSantosViana/microtoolkit/discovery"
	"github.com/DanielDanteDosSantosViana/microtoolkit/module"
	"github.com/DanielDanteDosSantosViana/microtoolkit/router"
	"github.com/fatih/color"
)

func Register(module module.Module) {
	results := make(chan *router.Router)
	routes := module.Params.Module.Routers

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(routes))

	var nameModule = module.Params.Module.Name

	for _, route := range routes {
		go func(route router.RouterP) {
			err := discovery.FindModule(nameModule)
			if err != nil {
				color.Red("Module %s - Error ('%s') ", nameModule, err.Error())
				err = discovery.CreateModule(module)
				if err != nil {
					color.Red("Module %s - Create - Error (%s) ", nameModule, err)
					waitGroup.Done()
					return
				}
				color.Blue("Module %s - Create - OK (new) ", nameModule)
			}

			err = discovery.FindRouter(nameModule, route.Path)
			if err != nil {
				color.Red("Router : %s to module: %s - Error ('%s') ", route.Path, nameModule, err.Error())
				err = discovery.CreateRouter(nameModule,route)
				if err!=nil{
					color.Red("Route %s - Create - Error (%s) ", route.Path, err)
					waitGroup.Done()
					return
				}
				color.Blue("Route %s - Create - OK (new) ", route.Path)
				waitGroup.Done()
				return
			}

		}(route)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

}
