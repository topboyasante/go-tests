package iterations

import (
	"fmt"
	"strings"
)

func Repeat(word string, count int) string {
	var repeated = fmt.Sprint(strings.Repeat(word, count))
	return repeated
}
