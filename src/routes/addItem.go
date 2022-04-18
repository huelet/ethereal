package routes

import (
  "github.com/gofiber/fiber/v2"
)

func CreateCampaign(app *fiber.App) {
  app.Post("/api/campaign", func(res *fiber.Ctx) error {
    var campaignName string = res.Query("name")
    return res.SendString(campaignName)
  })
}
func AddAdvert(app *fiber.App) {
  app.Post("/api/ad", func(res *fiber.Ctx) error {
    var adName string = res.Query("name")
    return res.SendString(adName)
  })
}