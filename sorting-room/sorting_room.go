package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	converted := strconv.FormatFloat(f, 'f', 1, 64)

	return fmt.Sprintf("This is the number %s", converted)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if fancy, ok := fnb.(FancyNumber); ok {
		// Try converting the value to an integer
		num, err := strconv.Atoi(fancy.Value())
		if err != nil {
			return 0
		}
		return num
	}
	return 0
}

// DescribeFancyNumberBox handles fancy number boxes and returns appropriate descriptions
func DescribeFancyNumberBox(box FancyNumberBox) string {
	number := ExtractFancyNumber(box)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(number))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"

	}
}
