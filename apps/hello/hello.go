package hello

import "github.com/gin-gonic/gin"
import "strconv"

// Post is the expected payload in the POST /hello endpoint
type Post struct {
    Message string `form:"message" json:"message" binding:"required"`
    Number int `form:"number" json:"number" binding:"required"`
}

// Put is the expected payload in the PUT /hello/:id endpoint
type Put struct {
    Something string `form:"some" json:"some" binding:"required"`
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
func ReadOne(ctx *gin.Context) {
    var response string
    maybeID := ctx.Param("id")
    if id, err := strconv.Atoi(maybeID); err == nil {
        response += "This would the data of Hello ID: " + strconv.Itoa(id)   
    } else {
        response += "No such ID found: " + maybeID   
    }
    ctx.JSON(200, gin.H {
        "message": response,
    })
}

// ReadAll will get you everything you wanted
func ReadAll(ctx *gin.Context) {
    ctx.JSON(200, gin.H {
        "message": "Hello all!",
    })
}

// Update will do stuff
func Update(ctx *gin.Context) {
    maybeID := ctx.Param("id")
    var json Put
    
    var response string
    var code int
    
    if id, err := strconv.Atoi(maybeID); err == nil {
        if ctx.BindJSON(&json) == nil {
            response += "Updating ID: " + strconv.Itoa(id)
            response += "New values: " + json.Something + " " + json.Message + " " + strconv.Itoa(json.Number)
            code = 200
        } else {
            response += "Invalid JSON"
            code = 400
        }
    } else {
        response += "No such ID found: " + maybeID
        code = 404
    }

    ctx.JSON(code, gin.H {
        "message": response,
    })
}

// Delete will be gone
func Delete(ctx *gin.Context) {
    var maybeID := ctx.Param("id")
    
    var code int
    if id, err := strconv.Atoi(maybeID); err == nil {
        fmt.Println("Removing ID: " + id)
        code = 204
    } else {
        code = 400
    }
    
    ctx.JSON(code, gin.H {})
}
