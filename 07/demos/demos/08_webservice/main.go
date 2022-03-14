package main

import (
	"net/http"

	"github.com/fjinss/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
