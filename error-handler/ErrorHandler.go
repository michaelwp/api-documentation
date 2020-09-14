package error_handler

import (
	"fmt"
	"log"
)

func ErrorLogger(e error)  {
	if e != nil {
		log.Println(fmt.Sprintf("Error: %v", e))
	}
}