// /
// go:generate scripts/generate-openapi.sh
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	port             *int
	postgresUser     = "app"
	postgresPassword *string
	postgresHost     *string
	postgresPort     = "5432"
)

func main() {
	port = flag.Int("port", 8081, "Port for test HTTP server")
	postgresPassword = flag.String("postgresPassword", "", "postgres Password")
	postgresHost = flag.String("postgresHost", "localhost", "postgres hostname")
	flag.Parse()

	//Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=movie port=%s", *postgresHost, postgresUser, *postgresPassword, postgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(Movie{})

	swagger, err := GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	ServerImplem := &ServerImplementation{
		DB: db,
	}

	e := echo.New()
	e.Use(echomiddleware.Logger())

	e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	RegisterHandlers(e, ServerImplem)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
