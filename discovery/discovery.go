package discovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/DanielDanteDosSantosViana/microtoolkit/module"
	"github.com/DanielDanteDosSantosViana/microtoolkit/router"
)

type routeSend struct {
	NameModule string         `json:"nameModule"`
	Router     router.RouterP `json:"router"`
}

var hostDiscovery = os.Getenv("MICRO_DISCOVERY")

func FindModule(nameModule string) (error, []byte) {
	resp, err := http.Get(hostDiscovery + "/servicediscovery/module/search?q=" + nameModule)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Not found module, return status code %d\n", resp.StatusCode), nil
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error when read response : %v", err), nil
	}

	return nil, result
}

func FindRouter(nameModule string, pathRouter string) (error, []byte) {
	resp, err := http.Get(hostDiscovery + "/servicediscovery/router/search?q=" + nameModule + pathRouter)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Not found router, return status code %d\n", resp.StatusCode), nil
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error when read response : %v", err), nil
	}

	return nil, result

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
	s := routeSend{nameModule, router}
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
