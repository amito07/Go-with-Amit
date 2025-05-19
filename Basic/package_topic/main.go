package main

import (
	"fmt"

	"github.com/amito07/package_topic/auth"
	"github.com/fatih/color"
)

func main() {
	result, session := auth.LoginWithCredentials("admin", "1234")
	fmt.Println("Login successful:", result, session)
	color.Red(session)
}