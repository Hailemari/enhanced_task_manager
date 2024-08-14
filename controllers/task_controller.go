package controllers

import (
	"net/http"
    "io"
    "bytes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/Hailemari/enhanced_task_manager/data"
	"github.com/Hailemari/enhanced_task_manager/models"
)

func GetTasks(ctx *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, found := data.GetTaskByID(id)
	if found {
		ctx.JSON(http.StatusOK, task)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
}

func DeleteTask(ctx *gin.Context) {
    id := ctx.Param("id")
    err := data.DeleteTask(id)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete task"})
        }
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}


func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task

	// Read the request body
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Check if the body is empty
	if len(bodyBytes) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Empty request body"})
		return
	}

	// Reassign the body to the request so it can be read again by ShouldBindJSON
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Bind the JSON body to the updatedTask struct
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the data layer to update the task in the database
	if err := data.UpdateTask(id, updatedTask); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}



func AddTask(ctx *gin.Context) {
	var newTask models.Task

	// Read the request body
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Check if the body is empty
	if len(bodyBytes) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Empty request body"})
		return
	}

	// Reassign the body to the request so it can be read again by ShouldBindJSON
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := data.AddTask(newTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

