package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var config map[string]string

// Response is represent the response structure
type Response struct {
	CurrentTime string `json:"current_time"`
	IpAddress   string `json:"ip_address"`
	Hostname    string `json:"hostname"`
	City        string `json:"city"`
	Region      string `json:"region"`
	Country     string `json:"country"`
}

// loadEnv loads the environmental variables
func loadEnv() {
	config = make(map[string]string, 0)
	fmt.Println("Loading the configuration...")
	for _, env := range os.Environ() {
		envPair := strings.SplitN(env, "=", 2)
		key := envPair[0]
		value := envPair[1]
		config[key] = value
	}
}

// main the entry point of the code
func main() {
	loadEnv()
	r := gin.Default()

	r.GET("/pucsd", myResponse) // endpoint

	r.Run(":8080") // server is starting at 8080 port
}

func myResponse(ctx *gin.Context) {
	loadEnv()

	// get the current time of IST
	currentTime := time.Now()
	timeFormat := currentTime.Format(time.RFC1123)

	// assigning the value to the structure
	rsp := &Response{
		CurrentTime: timeFormat,
		IpAddress:   config["IP_ADDRESS"],
		Hostname:    config["HOSTNAME"],
		City:        config["CITY"],
		Region:      config["REGION"],
		Country:     config["COUNTRY"],
	}
	if rsp == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "responsed value are not created"})
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}
