package main

import (
	"fmt"
	"net/http"

	"github.com/mitch-foster/pluralsight_go_getting_started.git/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("Server up and running!")
	http.ListenAndServe(":3000", nil)
}
