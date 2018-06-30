package main

import (
	"net/http"
	"fmt"
	"./modules/frontend/src"
	"./modules/admin/src"
	"./modules/user/src"
)

func main()  {

	frontend.Setup()
	admin.Setup()
	user.Setup()

	fmt.Println("start")
	http.ListenAndServe(":80", nil)

}