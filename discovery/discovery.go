package discovery

import (
	"fmt"
	"net/http"
)

type Result struct {
	StatusCode int
	MSG        string
}

func FindModule(nameModule string) (Result, error) {
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

func FindRouter(nameModule string, pathRouter string) (Result, error) {
	resp, err := http.Get("")
	if err != nil {
		return Result{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Result{resp.StatusCode, "Router not found"}, fmt.Errorf("Not found router, return status code %d\n", resp.StatusCode)
	}
	return Result{resp.StatusCode, "Router found"}, nil

}

func CreateModule(module Module) (Result, error) {
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

func CreateRouter(nameModule string, router Router) (Result, error) {
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