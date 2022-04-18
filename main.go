package main

import (
  "fmt"
  "reflect"
  
  "github.com/huelet/ethereal/src/utils"
  "github.com/huelet/ethereal/src/routes"
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()
  
  utils.InitDB()

  routes.AddAdvert(app)

  app.Listen(":3000")
}