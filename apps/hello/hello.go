package hello

import "github.com/gin-gonic/gin"
import "strconv"

// Post is the expected payload in the POST /hello endpoint
type Post struct {
    Message string `form:"message" json:"message" binding:"required"`
    Number int `form:"number" json:"number" binding:"required"`
}

// Create will handle the POST request
func Create(ctx *gin.Context) {
    var json Post
    var response string

    if ctx.bindJSON(&json) == nil {
        response += "Received message: " + json.Message
        response += "And also a value: " + strconv.Itoa(json.Number)
    }

    ctx.JSON(201, gin.H {
        "message": response,
    })
}

// Read will get you info
func Read(ctx *gin.Context) {
    ctx.JSON(200, gin.H {
        "message": "Hello there!",
    })
}

// Update will do stuff
func Update(ctx *gin.Context) {
    ctx.JSON(200, gin.H {
        "message": "Updated hello!",
    })
}

// Delete will be gone
func Delete(ctx *gin.Context) {
    ctx.JSON(204, gin.H {})
}
