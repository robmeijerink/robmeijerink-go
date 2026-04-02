package main

import (
    "log"
    "os"
    "strings"

    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/recover"
    // "github.com/gofiber/fiber/v3/middleware/static"
)

func viteCacheControl() fiber.Handler {
    return func(c fiber.Ctx) error {
        path := c.Path()
        if strings.HasPrefix(path, "/assets/") {
            c.Set(fiber.HeaderCacheControl, "public, max-age=31536000, immutable")
        } else {
            c.Set(fiber.HeaderCacheControl, "no-cache")
        }
        return c.Next()
    }
}

func setupSolvalutionsApp() *fiber.App {
    app := fiber.New()

    // Temporary Catch-All Route for testing
    app.Use(func(c fiber.Ctx) error {
        // Resultaat: "Dit is solvalutions.nl/over-ons"
        return c.SendString("Dit is solvalutions.nl" + c.Path())
    })

    return app
}

func setupRobMeijerinkApp() *fiber.App {
    app := fiber.New()

    app.Get("/workstation-setup", func(c fiber.Ctx) error {
        c.Set(fiber.HeaderStrictTransportSecurity, "max-age=31536000; includeSubDomains")
        return c.Redirect().Status(fiber.StatusFound).To("https://raw.githubusercontent.com/robmeijerink/workstation-setup/main/workstation-setup.sh")
    })

    // Temporary Catch-All Route
    app.Use(func(c fiber.Ctx) error {
        return c.SendString("Dit is robmeijerink.nl" + c.Path())
    })

    return app
}

func main() {
    gateway := fiber.New()
    gateway.Use(recover.New())

    solApp := setupSolvalutionsApp()
    robApp := setupRobMeijerinkApp()

    gateway.Use(func(c fiber.Ctx) error {
        host := strings.TrimPrefix(c.Hostname(), "www.")

        if host == "solvalutions.nl" {
            solApp.Handler()(c.RequestCtx())
            return nil
        }

        // robmeijerink.nl is the default fallback
        robApp.Handler()(c.RequestCtx())
        return nil
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(gateway.Listen(":" + port))
}
