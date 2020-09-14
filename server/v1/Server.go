package v1

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

type Svr http.Server

func (s *Svr)Server() error {
	var err error

	// create new echo server
	e := echo.New()
	err = e.StartServer((*http.Server)(s))

	// set cors
	e.Use(middleware.CORS())

	// set logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf(
			"%v - %v: method=%v, uri=%v, status=%v\n",
		"${remote_ip}", time.Now(), "${method}", "${uri}",
		"${status}"),
	}))

	e.Logger.Fatal(err)
	return err
}