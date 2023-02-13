package main

// Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.
import (
	"hanbit-react/rest"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Go Webframework Echo Init...")
	logrus.Fatal(rest.RunAPI(":8080"))
}
