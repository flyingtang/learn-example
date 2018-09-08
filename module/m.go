package main

// go 1.11 新特效module 尝试

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.Run(":3000")
}