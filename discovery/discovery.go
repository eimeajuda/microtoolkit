package discovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/DanielDanteDosSantosViana/microtoolkit/module"
	"github.com/DanielDanteDosSantosViana/microtoolkit/router"
)

type Result struct {
	StatusCode int
	MSG        string
}

type routeSend struct {
	NameModule string `json:"nameModule"`
	Router router.RouterP `json:"router"`
}


var hostDiscovery = os.Getenv("MICRO_DISCOVERY")

func FindModule(nameModule string) error {
	resp, err := http.Get(hostDiscovery + "/servicediscovery/module/search?q=" + nameModule)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Not found module, return status code %d\n", resp.StatusCode)
	}
	return nil
}

func FindRouter(nameModule string, pathRouter string) error {
	resp, err := http.Get(hostDiscovery + "/servicediscovery/router/search?q=" + nameModule + pathRouter)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Not found router, return status code %d\n", resp.StatusCode)
	}
	return nil

}

func CreateModule(moduleP module.Module) error {
	module := moduleP.Params.Module
	moduleJ, err := json.Marshal(module)
	if err != nil {
		return err
	}
	resp, err := http.Post(hostDiscovery+"/serviceregistry/module", "application/json", bytes.NewBuffer(moduleJ))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Not found module, return status code %d\n", resp.StatusCode)
	}
	return nil

}

func CreateRouter(nameModule string, router router.RouterP) error {
	s:= routeSend{nameModule,router}
	routeJ, err := json.Marshal(s)
	if err != nil {
		return err
	}
	resp, err := http.Post(hostDiscovery+"/serviceregistry/router", "application/json", bytes.NewBuffer(routeJ))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Not found module, return status code %d\n", resp.StatusCode)
	}
	return nil

}
