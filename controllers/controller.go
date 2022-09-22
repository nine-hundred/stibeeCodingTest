package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.POST("/stage1", Stage1)
	r.POST("/stage2", Stage2)
	r.POST("/stage3", Stage3)
	r.POST("/stage/result", StageResult)
	return r
}

type Result struct {
	Id          string `json:"id"`
	Level       int    `json:"level"`
	Success     bool   `json:"success"`
	Want        int    `json:"want"`
	Actual      int    `json:"actual"`
	ElapsedTime int    `json:"elapsedTime"`
}

func Stage1(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body)
	for i, _ := range buf {
		buf[i]++
	}
	c.String(http.StatusOK, string(buf))
}

func Stage2(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body)
	for i, b := range buf {
		if b == 10 {
			continue
		}
		buf[i]++
	}
	c.String(http.StatusOK, string(buf))
}

func Stage3(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(buf))
	fmt.Println("===end==")
	for i, b := range buf {
		if b == 10 {
			continue
		}
		buf[i]++
	}
	c.String(http.StatusOK, string(buf))
}

func StageResult(c *gin.Context) {
	result := Result{}
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Printf("level:%d\n", result.Level)
	fmt.Printf("success:%t\n", result.Success)
	fmt.Printf("Actual:%d\n", result.Actual)
}
