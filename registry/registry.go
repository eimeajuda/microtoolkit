package registry

import (
	"log"
	"sync"

	"github.com/DanielDanteDosSantosViana/microtoolkit/discovery"
)

func Register(params ...discovery.ParamM) {
	module := discovery.NewModule(params...)
	results := make(chan *discovery.Router)
	routes := module.Params.Module.Routers

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(routes))

	var nameModule = module.Params.Module.Name

	for _, route := range routes {
		go func(route discovery.Router) {
			_, err := discovery.FindModule(nameModule)
			if err != nil {
				log.Println("Router : " + route.Params.Router.Path + " to Module " + " - Error (" + err.Error() + " )")
				continue
			}
			result,err = discovery.FindRouter(nameModule, route.Params.Router.Path)
			if err != nil {
				log.Println("Router : " + route.Params.Router.Path + " to Module " + " - Error (" + err.Error() + " )")
				continue
			}

			log.Println("Router : " + route.Params.Router.Path + " to Module " + " - OK ( EXISTIS )")
				} else {
					result = discovery.CreateRouter(nameModule, route)
					if result.Created {
						log.Println("Router : " + route.Params.Router.Path + " to Module " + " - OK (NEW )")
					}
				}
			}
			waitGroup.Done()
		}(route)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

}
