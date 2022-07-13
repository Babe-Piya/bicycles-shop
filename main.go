package main

import (
	"bicycles-shop/model"
	"fmt"
	"log"
)

func main() {
	if err := model.ConnectDB(); err != nil {
		log.Panic("Cannot connect database ", err)
	}

	fmt.Println("Testt")

}
