package main

import (
	"fmt"

	"github.com/AlejandroGleles/academichub/config"
	"github.com/AlejandroGleles/academichub/router"
)

func main() {

	//Initialize config
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	//initialize router
	router.Initialize()
}
