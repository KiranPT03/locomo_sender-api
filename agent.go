package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type applicationWindow struct {
	SourceType        string `json:"source_type"`
	SourceApplication string `json:"source_application"`
	SourceURL         string `json:"source_url"`
	SourceSection     string `json:"source_section"`
}

var (
	router = gin.Default()
)

func runApplicationHandler(ctx *gin.Context) {

	var runApplication applicationWindow

	if err := ctx.ShouldBindJSON(&runApplication); err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	fmt.Println(runApplication)
	if runApplication.SourceType == "web" {
		cmd := exec.Command("C:/Program Files (x86)/Microsoft/Edge/Application/msedge", runApplication.SourceURL)
		err := cmd.Start()

		if err != nil {
			log.Fatal(err)
		}
	} else if runApplication.SourceType == "application" {

		command := "C:/Program Files/Microsoft Office/root/Office16/" + runApplication.SourceApplication
		cmd := exec.Command(command)
		err := cmd.Start()

		if err != nil {
			log.Fatal(err)
		}

	}
	ctx.Status(http.StatusOK)

}

func mapUrls() {
	router.POST("/v1/controller/run-application/", runApplicationHandler)
}

func main() {
	mapUrls()
	router.Run(":8000")
}
