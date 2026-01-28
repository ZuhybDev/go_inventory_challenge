import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type FirstAid struct {
	Essentials string
	Aid        string
	ExpireDate time.Time
	Quantity   int
}

func evaluateFirstAidKit(supplies []string, householdSize int, currentDate string) bool {

	suppliesData := make([]FirstAid, 0)

	suppliesParts := make([]string, 0)

	layout := "2006-01-02"
	formattedDate, _ := time.Parse(layout, currentDate)

	for _, v := range supplies {
		suppliesEntry := strings.Split(v, ",")

		suppliesParts = append(suppliesParts, suppliesEntry...)

	}

	for _, v := range suppliesParts {
		parts := strings.Split(v, ":")

		essen := parts[0]
		aid := parts[1]
		expireDate, _ := time.Parse(layout, parts[2])
		quantity, _ := strconv.Atoi(parts[3])

		suppliesData = append(suppliesData,
			FirstAid{
				Essentials: essen,
				Aid:        aid,
				ExpireDate: expireDate,
				Quantity:   quantity,
			},
		)
	}

	if len(suppliesData) < 3 {
		return false
	}

	limit := 3

	if len(suppliesData) < limit {
		limit = len(suppliesData)
	}

	for _, v := range suppliesData[:limit] {

		if householdSize == 0 {
			fmt.Println("ERROR: house hold size must be greater then Zero")
		}

		if v.Quantity%householdSize != 0 {
			return false
		}

		// expire check
		if v.ExpireDate.Before(formattedDate) {
			return false
		}

	}

	return true
}
