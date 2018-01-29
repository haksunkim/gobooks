package controllers

import (
	"errors"

	"github.com/kataras/iris/mvc"
)

type MainController struct{}

var mainView = mvc.View{
	Name: "main/index.html",
	Data: map[string]interface{}{
		"Title":     "Main Page",
		"MyMessage": "Welcome to GoBooks",
	},
}

func (c *MainController) Get() mvc.Result {
	return mainView
}

var errBadName = errors.New("bad name")

var badName = mvc.Response{Err: errBadName, Code: 400}
