package main

import (
	"github.com/haksunkim/gobooks/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// load the template files
	app.RegisterView(iris.HTML("./web/views", ".html"))

	app.StaticWeb("/static", "./web/assets")

	// serve our controllers
	mvc.New(app.Party("/")).Handle(new(controllers.MainController))

	app.Run(
		// start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// disable updates;
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
