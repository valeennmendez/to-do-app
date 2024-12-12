package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valeennmendez/to-do/connection"
	"github.com/valeennmendez/to-do/models"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "session error",
		})
		return
	}

	userID, ok := session.Values["userID"]
	/* 	fmt.Printf("Tipo de ID: %T\n", userID)
	   	fmt.Println("funcion del gettask", userID) */
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userid not found",
		})
		return
	}

	task.UserID = userID.(int)

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error to decode petition",
		})
		return
	}

	if err := connection.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "dont cant create task",
		})
		return
	}

	var value = task.ID

	c.JSON(http.StatusOK, gin.H{
		"message": "task created",
		"task":    value,
	})

}

func DeleteTask(c *gin.Context) {
	var task models.Task

	id := c.Param("id")

	if err := connection.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	if err := connection.DB.Unscoped().Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to delete task",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "task deleted successfully",
	})

}

func GetAllTask(c *gin.Context) {
	var task []models.Task

	connection.DB.Find(&task)

	c.JSON(http.StatusAccepted, &task)
}

func GetTask(c *gin.Context) {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "session error",
		})
		return
	}

	userID, ok := session.Values["userID"]
	/* 	fmt.Printf("Tipo de ID: %T\n", userID)
	   	fmt.Println("funcion del gettask", userID) */
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userid not found",
		})
		return
	}

	fmt.Println("el id del usuario es:", userID)
	var task []models.Task

	//PROBLEMA EN LA CONSULTA VERIFICAR ESTO, SE RECUPERA EL ID DEL USUARIO CORRECTAMENTE EN LAS DEMAS FUNCIONES.
	if err := connection.DB.Where("user_id = ?", userID.(int)).Find(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": task,
	})

}
