// https://pkg.go.dev/time#Time.Format

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println(currentTime)

	formattedTime := currentTime.Format("02-01-2006") // Day, Month and Year values must be 02 01 and 2006 respectively

	fmt.Println(formattedTime)

	fullFormat := currentTime.Format("02-01-2006 15:04:05 Monday")

	fmt.Println(fullFormat)

	createdDate := time.Date(2024, time.January, 8, 9, 22, 37, 0, time.Local)

	fmt.Println(createdDate)
}
