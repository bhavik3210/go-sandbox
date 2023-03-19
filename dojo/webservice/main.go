package main

import (
	"net/http"

	"github.com/bhavik3210/go-project/webservice/controllers"
)

func main() {
	// user := models.User{
	// 	ID:        1,
	// 	FirstName: "Arthur",
	// 	LastName:  "Morgan",
	// }

	// fmt.Println(user)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
