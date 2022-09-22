package main

import (
	"fmt"
	"io/ioutil"
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

	req, err := http.NewRequest("POST", attackServerUrl+"&level="+level, nil)
	if err != nil {
		fmt.Printf("stage%s can't start", level)
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
	fmt.Println("stage" + level + "====================")
	return nil
}
