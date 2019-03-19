package robotname

import (
	"time"
	"math/rand"
)

var names = map[string]bool{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = generateUniqueName()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateUniqueName() string {
	for {
		name := generateName()
		if !names[name] {
			names[name] = true
			return name
		}
	}
}

func generateName() string {
	return string([]byte{
		randomLetter(), randomLetter(),
		randomDigit(), randomDigit(), randomDigit(),
	})
}

func randomLetter() byte {
	return byte(rand.Intn(26) + 'A')
}

func randomDigit() byte {
	return byte(rand.Intn(10) + '0')
}
