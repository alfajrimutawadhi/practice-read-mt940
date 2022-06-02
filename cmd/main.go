package main

import (
	"belajar-convert-mt940/helper"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(helper.ConvertMT940("dir/BTN IDR_Mock.txt"))
	})

	app.Listen(":3000")


	// const FTP_HOST = "files.000webhost.com:21"
	// const FTP_USERNAME = "akuanakpintar"
	// const FTP_PASSWORD = "akuanakpintar"

	// conn, err := ftp.Dial(FTP_HOST)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = conn.Login(FTP_USERNAME, FTP_PASSWORD)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // read file mt940-npp-sample-file.txt
	// ftpResponse, err := conn.Retr("./htdocs/BTN IDR_Mock.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// response := helper.ConvertMT940(*ftpResponse)
	// fmt.Println(response)
}
