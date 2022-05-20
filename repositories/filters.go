package repositories

import "net/url"

type QuestionFilters struct {
	Filters map[string]interface{}
}

func NewFilters(queryParamaters url.Values) (QuestionFilters, error) {

	filters := make(map[string]interface{})

	for key, val := range queryParamaters {

		// the value is an array of data, only use if something is there
		if len(val) == 0 {
			continue
		}

		//exact filters
		if key == "id" {
			filters["id"] = val[0]
		}

		if key == "category" {
			filters["category"] = val[0]
		}

		if key == "level" {
			filters["level"] = val[0]
		}
	}

	return QuestionFilters{
		Filters: filters,
	}, nil
}
