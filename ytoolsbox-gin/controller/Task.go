package controller

import (
	"log"
	"main/service"

	"github.com/gin-gonic/gin"
)

func PostNewTask(ctx *gin.Context) {
	log.Println("You Create a New Task!")

	go service.CreateNewTaskService()

}
