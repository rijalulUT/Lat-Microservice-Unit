package main

import (
	// "pegawaimicroservice/pegawai"

	"fmt"
	"os"
	"unitmicroservice/config"
	"unitmicroservice/unit"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	r := gin.Default()
	db := config.DBInit()
	unit := unit.Units{DB: db}

	port := os.Getenv("PORT")

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	config.RegisterConsul()
	config.RegisterZipkin()

	r.GET("/unit", unit.GetUnit)
	r.GET("/unit/:id", unit.GetUnitById)
	r.GET("/healthcheck", config.Healthcheck)
	r.Run(":" + port)
}
