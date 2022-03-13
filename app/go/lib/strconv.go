package lib

import (
	"strconv"
)

func Atoui(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}