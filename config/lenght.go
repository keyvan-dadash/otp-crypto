package config

import "math"

type Lenght int

const (
	Lenght_6 Lenght = iota + 6
	Lenght_7
	Lenght_8
	Lenght_9
	Lenght_10
)

func (l Lenght) Truncate(number int) int {
	return number % int(math.Pow10(int(l)))
}
