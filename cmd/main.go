package main

import (
	"fmt"
	"log"

	"github.com/Blue-Onion/MahilAi/handler/config"
)

func main(){
	camera,err:=config.LoadConfig()
	if err!=nil{
		log.Fatal(err.Error())
	}
	fmt.Println(camera)
	
}