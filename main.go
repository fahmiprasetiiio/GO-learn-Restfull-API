package main

import (
	"os"

	"github.com/fahmiprasetiiio/restAPI/handler"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	envFile := os.Getenv("ENV")
	if envFile == "" {
		envFile = ".env"
	}

	// Load env data
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	// // CREATING Table
	// table := new(handler.TableFahmi)

	// err = table.CreateTable(db)
	// if err != nil {
	// 	panic(err)
	// }

	// ENDPOINT
	e.GET("/foo/:params/", handler.Bar)
	e.POST("/post", handler.PostSomething)

	// TENTANG TABLE FAHMI
	e.POST("/tablefahmi", handler.CreateData)
	e.GET("/tablefahmi", handler.GetData)
	e.GET("/tablefahmis", handler.GetAllData)

	// Start echo on desired port
	e.Start(":" + os.Getenv("PORT"))

	// Start echo on desired port
	// e.Logger.Fatal(e.Start(":5050"))
}
