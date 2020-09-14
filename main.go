package main

import (
	error_handler "github.com/michaelwp/api-documentation/error-handler"
	v1 "github.com/michaelwp/api-documentation/server/v1"
	"time"
)

func main() {
	// server options
	s := v1.Svr{
		Addr:         ":8093",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	// call server
	err := s.Server()
	error_handler.ErrorLogger(err)
}
