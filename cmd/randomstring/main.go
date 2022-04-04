package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	length := flag.Int("length", 10, "Length of the string to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	runes := ascii()
	l := len(runes)
	var s strings.Builder
	for i := 0; i < *length; i++ {
		r := runes[rand.Intn(l)]
		s.WriteRune(r)
	}
	fmt.Println(s.String())
}

func ascii() []rune {
	runes := make([]rune, 0, 107)
	for r := ' '; r <= '~'; r++ {
		runes = append(runes, r)
	}
	return runes
}
