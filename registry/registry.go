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
			_, err := discovery.FindModule(nameModule)
			if err != nil {
				color.Red("Module %s - Error ('%s') ", nameModule, err.Error())
				_, err = discovery.CreateModule(module)
				if err != nil {
					color.Red("Module %s - Create - Error (%s) ", nameModule, err)
					return
				}
				color.Blue("Module %s - Create - OK (new) ", nameModule)
			}

			_, err = discovery.FindRouter(nameModule, route.Path)
			if err != nil {
				color.Red("Router : %s to module: %s - Error ('%s') ", route.Path, nameModule, err.Error())
				return
			}
			/*

				log.Println("Router : " + route.Params.Router.Path + " to Module " + " - OK ( EXISTIS )")
				result, err := discovery.CreateRouter(nameModule, route)
				if err != nil {
					color.Red("Router : %s to module: %s - Error ('%s') ", route.Params.Router.Path, nameModule, err.Error())
				}

				if result.StatusCode != 200 {
					color.Red("Error to create Router : %s MSG: %s", route.Params.Router.Path, result.MSG)
				}*/
			//			color.Blue("Router : %s to module: %s - OK (new)", route.Params.Router.Path, nameModule)

			waitGroup.Done()
		}(route)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

}
