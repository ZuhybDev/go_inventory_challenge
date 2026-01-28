package main

import (
	"fmt"
	"strings"
)

func SummarizeCarWash(services []string, volunteers []string) string {
	// Write code here
	result := fmt.Sprintf("Service: [%s] - Volunteer: [%s]", strings.Join(services, ","), strings.Join(volunteers, ", "))
	return result
}

func main() {

	service := []string{"Basic Wash", "Premium Wash"}
	volunteer := []string{"Zuhaib", "Ahmed"}

	SummarizeCarWash(service, volunteer)

}
