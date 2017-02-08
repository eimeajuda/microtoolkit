package discovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DanielDanteDosSantosViana/microtoolkit/module"
	"github.com/DanielDanteDosSantosViana/microtoolkit/router"
)

type Result struct {
	StatusCode int
	MSG        string
}

var hostDiscovery = os.Getenv("MICRO_DISCOVERY")

func FindModule(nameModule string) (Result, error) {
	resp, err := http.Get(hostDiscovery + "/servicediscovery/module/search?q=" + nameModule)
	if err != nil {
		return Result{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Result{resp.StatusCode, "Module not found"}, fmt.Errorf("Not found module, return status code %d\n", resp.StatusCode)
	}
	return Result{resp.StatusCode, "Module found"}, nil
}

func FindRouter(nameModule string, pathRouter string) (Result, error) {
	resp, err := http.Get(hostDiscovery + "/servicediscovery/router/search?q=" + nameModule + pathRouter)
	if err != nil {
		return Result{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Result{resp.StatusCode, "Router not found"}, fmt.Errorf("Not found router, return status code %d\n", resp.StatusCode)
	}
	return Result{resp.StatusCode, "Router found"}, nil

}

func CreateModule(moduleP module.Module) (Result, error) {
	module := moduleP.Params.Module
	moduleJ, err := json.Marshal(module)
	if err != nil {
		return Result{}, err
	}
	jose := string(moduleJ)
	log.Print(jose)
	resp, err := http.Post(hostDiscovery+"/serviceregistry/module", "application/json", bytes.NewBuffer(moduleJ))
	if err != nil {
		return Result{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Result{resp.StatusCode, "Module not created"}, fmt.Errorf("Not found module, return status code %d\n", resp.StatusCode)
	}
	return Result{resp.StatusCode, "Module found"}, nil

}

func CreateRouter(nameModule string, router router.RouterP) (Result, error) {

	resp, err := http.Get("")
	if err != nil {
		return Result{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Result{resp.StatusCode, "Module not found"}, fmt.Errorf("Not found module, return status code %d\n", resp.StatusCode)
	}
	return Result{resp.StatusCode, "Module found"}, nil

}
