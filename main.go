package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"log"
	"os/exec"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		command := "shutdown"
		arg1 := "-h"
		arg2 := "now"

		cmd := exec.Command(command, arg1, arg2)
		_, err := cmd.Output()
		if err != nil {
			return ctx.SendString(err.Error())
		}

		return ctx.SendString("ok")
	})

	err := app.Listen(":8080")

	log.Println(err)
}
