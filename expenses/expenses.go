package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

func Filter(records []Record, predicate func(Record) bool) []Record {
	var filtered []Record
	for _, record := range records {
		if predicate(record) {
			filtered = append(filtered, record)
		}
	}
	return filtered
}

// 2. Filter records within a period of time
func ByDaysPeriod(period DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= period.From && r.Day <= period.To
	}
}

// 3. Filter records by category
func ByCategory(category string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == category
	}
}

// 4. Calculate the total amount of expenses in a period
func TotalByPeriod(records []Record, period DaysPeriod) float64 {
	filtered := Filter(records, ByDaysPeriod(period))
	var total float64
	for _, record := range filtered {
		total += record.Amount
	}
	return total
}

// 5. Calculate the total expenses for records of a category in a period
func CategoryExpenses(records []Record, period DaysPeriod, category string) (float64, error) {
	// First, check if the category exists in the records
	var categoryExists bool
	for _, record := range records {
		if record.Category == category {
			categoryExists = true
			break
		}
	}
	if !categoryExists {
		return 0, errors.New("unknown category " + category)
	}

	// Filter by both category and period
	filtered := Filter(records, func(r Record) bool {
		return r.Category == category && r.Day >= period.From && r.Day <= period.To
	})

	var total float64
	for _, record := range filtered {
		total += record.Amount
	}

	return total, nil
}
