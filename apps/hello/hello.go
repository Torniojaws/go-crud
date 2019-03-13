package hello

import "github.com/gin-gonic/gin"

// Create will create hellos
func Create(ctx *gin.Context) {
    ctx.JSON(201, gin.H {
        "message": "Hello added!",
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
