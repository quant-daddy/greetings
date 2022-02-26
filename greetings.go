package greetings

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	return fmt.Sprintf("Hi %v. Welcome!", name), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hellos(names []string) (messages []string, err error) {
	for _, n := range names {
		r, err := Hello(n)
		if err != nil {
			return messages, err
		}
		messages = append(messages, r)
	}
	return messages, err;
}