package picker

import (
	"math/rand"
)

func Pick(options []string) string {
	return options[rand.Intn(len(options))]
}
