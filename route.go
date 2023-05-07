package main

import (
	"encoding/json"
	"fmt"
	"ids/config"

	"github.com/gofiber/fiber/v2"
)

func initRoute(app *fiber.App) {
	app.Get("/snowflake/next", SnowflakeNextHandle)
	app.Get("/snowflake/batch", SnowflakeBatchHandle)
}

// SnowflakeNextHandle generate snowflak ids
func SnowflakeNextHandle(c *fiber.Ctx) error {
	id := config.SnowflakeNode.Generate()
	response := SnowflakNextResponse{
		Id: id.Int64(),
	}
	return sendJson(c, response)
}

// SnowflakeBatchHandle generate snowflak ids batch
func SnowflakeBatchHandle(c *fiber.Ctx) error {
	size := c.QueryInt("size", 1)
	if size <= 0 {
		size = 1
	}

	ids := []int64{}
	for i := 0; i < size; i++ {
		id := config.SnowflakeNode.Generate()
		ids = append(ids, id.Int64())
	}

	response := SnowflakBatchResponse{ids}
	return sendJson(c, response)
}

func sendJson(c *fiber.Ctx, data any) error {
	c.Response().Header.Add("Content-Type", "application/json")
	body, err := json.Marshal(data)
	if err != nil {
		c.Status(500)
		c.SendString(fmt.Sprintf(`{"code": "%s", "msg": "%s", "data": null}`, "500", "marshal json error: "+err.Error()))
	}
	return c.Send(body)
}
