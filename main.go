package main

import (
	"fmt"
	"net/http"
	"stibee/controllers"
)

func main() {
	r := controllers.InitRouter()
	err := startStage("1")
	if err != nil {
		panic(err)
	}
	r.Run(":8090")
}

func startStage(level string) error {
	attackServerUrl := "http://attack.stibee.kr:8090/stage?id=ninehundred@outlook.kr"

	_, err := http.NewRequest("POST", attackServerUrl+"&level="+level, nil)
	if err != nil {
		fmt.Printf("stage%s can't start", level)
		return err
	}
	return nil
}
