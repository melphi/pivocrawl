package testit

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melphi/pivocrawl/infra/deps"
)

type TestContext struct {
	API *fiber.App
}

func StartTestContext() *TestContext {
	deps := deps.Init()
	return &TestContext{
		API: deps.ApiService.GetAPI(),
	}
}
