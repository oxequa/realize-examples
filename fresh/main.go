package main

import (
	"errors"
	"github.com/tockins/fresh"
	"net/http"
	"os"
)

type Todo struct {
	Title string `json:"title"`
}

func main() {
	f := fresh.New()
	f.Config().SetPort(8080)
	// group
	g := f.Group("/group/").Before(filter1, filter1).After(filter1, filter1)
	g.GET("/test/", list).Before(filter2, filter3).After(filter2)
	g.GET("test/{id}", single)
	g.GET("test/alpha", single)
	f.Run()
}

func list(f fresh.Context) error {
	data := []Todo{{Title: "Buy milk: list"}}
	dir, _ := os.Getwd()
	return f.Response().Download(http.StatusOK, dir+"/listino.pdf")
	return f.Response().JSON(http.StatusOK, data)
}

func single(f fresh.Context) error {
	data := Todo{Title: "Buy milk single"}
	return errors.New("Error")
	return f.Response().JSON(http.StatusOK, data)
}

func filter1(f fresh.Context) error {
	return nil
}
func filter2(f fresh.Context) error {
	return nil
}
func filter3(f fresh.Context) error {
	return nil
}
