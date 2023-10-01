package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID      string `json:"id"`
	Title   string ` json:"title"`
	Content string `json:"content"`
}

var tasks = []task{
	{
		ID:      "1",
		Title:   "tasks one",
		Content: "task content",
	},
}

func getTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)

}
func getTaskID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return

		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task no found"})
}

func postTasks(c *gin.Context) {
	var newTasks task
	if err := c.BindJSON(&newTasks); err != nil {
		return
	}
	tasks = append(tasks, newTasks)
	c.IndentedJSON(http.StatusCreated, newTasks)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTask)
	router.GET("tasks/:id", getTaskID)

	router.POST("/tasks", postTasks)
	router.Run("localhost:8000")
}
