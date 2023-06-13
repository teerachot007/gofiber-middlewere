// package main

// import (
// 	"encoding/json"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// 	"github.com/gofiber/fiber/v2/middleware/requestid"
// 	"github.com/gofiber/fiber/v2/utils"
// 	middleauth "github.com/teerachot007/gofiber-middlewere/middle-auth"
// )

// type dataclaim struct {
// 	Content struct {
// 		Type string `json:"type"`
// 		User string `json:"user"`
// 	} `json:"context"`
// 	Exp int `json:"exp"`
// }

// func main() {
// 	app := fiber.New()
// 	app.Use(requestid.New())
// 	app.Use(requestid.New(requestid.Config{
// 		Header: "Test-Service-Header",
// 		Generator: func() string {
// 			return utils.UUID()
// 		},
// 	}))
// 	app.Use(logger.New(logger.Config{
// 		Format:     "${pid} ${status} - ${method} ${path}\n",
// 		TimeFormat: "02-Jan-2006",
// 		TimeZone:   "Asia/Bangkok",
// 	}))
// 	// privatekey := middleauth.GenKey()
// 	ACCESS_TOKEN_SECRET := "EiKf9vBVMW0Qiu6EWgzwU7PyCdD0BLxv7ks4kTe4fXvGPDYsS3QT3wugV4ReGopt"
// 	REFRESH_TOKEN_SECRET := "0ueUlWRDDjvu7188rORSqZVuwWUVvJSyPGWw84J3HxgWmW9VKRP4RFzW2Imvb1Jr"

// 	app.Post("/login", func(c *fiber.Ctx) error {
// 		content := map[string]interface{}{
// 			"user": "join wick",
// 			"type": "action",
// 		}
// 		// uid := "144479bd-fcdc-4c9f-b116-f2a08807a4c3" //utils.UUID()
// 		token := middleauth.GenerateToken(content, time.Minute*10, time.Hour*24, ACCESS_TOKEN_SECRET, REFRESH_TOKEN_SECRET)

// 		c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"access_token":  token.AccessToken,
// 			"refresh_token": token.RefreshToken,
// 		})

// 		return nil
// 	})
// 	app.Post("/refresh", func(c *fiber.Ctx) error {
// 		refreshToken := struct {
// 			RefreshToken string `json:"refresh_token"`
// 		}{}
// 		if err := c.BodyParser(&refreshToken); err != nil {
// 			return err
// 		}

// 		dd_cl := dataclaim{}
// 		err := middleauth.RefreshToken(refreshToken.RefreshToken, REFRESH_TOKEN_SECRET, &dd_cl)
// 		if err != nil {
// 			return err
// 		}
// 		content := map[string]interface{}{
// 			"user": "join wick",
// 			"type": "action",
// 		}
		
// 		if dd_cl.Content.User == "join wick" {
// 			token := middleauth.GenerateToken(content, time.Minute*10, time.Hour*24, ACCESS_TOKEN_SECRET, REFRESH_TOKEN_SECRET)
// 			c.Status(fiber.StatusOK).JSON(fiber.Map{
// 				"access_token":  token.AccessToken,
// 				"refresh_token": token.RefreshToken,
// 			})
// 		}
// 		return nil
// 	})

// 	app.Use(middleauth.MiddleWare(ACCESS_TOKEN_SECRET))

// 	app.Get("/getuser", user)

// 	err := app.Listen(":5000")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func user(c *fiber.Ctx) error {
// 	cl := middleauth.Getclaims(c)

// 	ff := dataclaim{}
// 	err := json.Unmarshal([]byte(string(cl)), &ff)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c.JSON(ff)
// }
