package registry

import (
	"sync"

	"github.com/DanielDanteDosSantosViana/microtoolkit/discovery"
	"github.com/DanielDanteDosSantosViana/microtoolkit/module"
	"github.com/DanielDanteDosSantosViana/microtoolkit/router"
	"github.com/fatih/color"
)

type Result struct {
	Color *color.Color
	MSG   string
}

func Register(module module.Module) {
	results := make(chan *Result)
	routes := module.Params.Module.Routers
	//colors
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgGreen)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(routes))

	var nameModule = module.Params.Module.Name

	for _, route := range routes {
		go func(route router.RouterP) {
			err, _ := discovery.FindModule(nameModule)
			if err != nil {

				result := &Result{red, "Module " + nameModule + " - Error ('" + err.Error() + "') "}
				results <- result
				err = discovery.CreateModule(module)
				if err != nil {
					result = &Result{red, "Module " + nameModule + " - Create - Error (" + err.Error() + ")"}
					results <- result
					waitGroup.Done()
					return
				}
				result = &Result{blue, "Module " + nameModule + " -  Create - OK (new)"}
				results <- result

			}

			err, _ = discovery.FindRouter(nameModule, route.Path)
			if err != nil {
				result := &Result{red, "Router : " + route.Path + " to module: " + nameModule + " - Error ('" + err.Error() + "') "}
				results <- result
				err = discovery.CreateRouter(nameModule, route)
				if err != nil {
					result = &Result{red, "Route " + route.Path + " - Create - Error (" + err.Error() + ") "}
					results <- result
					waitGroup.Done()
					return
				}
				result = &Result{blue, "Route " + route.Path + " - Create - OK (new)"}
				results <- result
				waitGroup.Done()
				return
			}

			result := &Result{green, "Authenticated '" + route.Path + "' route in module '" + nameModule + "' - (OK)"}
			results <- result
			//color.Green("Authenticated '" + route.Path + "' route in module '" + nameModule + "' - (OK) \n")
			waitGroup.Done()
		}(route)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	log(results)
}

func log(results chan *Result) {
	for result := range results {
		result.Color.Println(result.MSG)
	}
}
