package seb

import (
	"fmt"
	"time"
)

func EpochString(input string) string {
	nano := time.Now().UnixNano()
	return fmt.Sprintf("String %s with epoch %d", input, nano)
}
