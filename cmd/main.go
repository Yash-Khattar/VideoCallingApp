package main

import(
	"log"
	"video-calling-app/internal/server"
)

func main(){
	if err : = server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}