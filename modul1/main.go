package main // untuk pengelompokan package, file termasuk paket mana. dampaknya : bisa panggil function dalam package yg sama. klo beda gabisa.  
 
import (
    "log"
 
    "github.com/gofiber/fiber/v2" // untuk import package fiber agar bisa pakai framework fiber untuk golang
)
 
func main() {
    app := fiber.New() // untuk membuat aplikasi/server fiber baru
 
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!") // untuk print hallo word, sebagai tanda berhasil atau gak.
    })
 
    log.Fatal(app.Listen(":3000"))
}

