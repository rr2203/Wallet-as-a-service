package error_handling

import (
	"log"
)

func HandleError(httpStatusCode int, userResponse string, errorMessage string) string {
	log.Println(errorMessage)
	return string(httpStatusCode) + " " + userResponse
}