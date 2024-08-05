package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("nombre es vacio ")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	format := make([]string, 3)
	format[0] = "Hola %v, Bienvenido"
	format[1] = "Que bueno verte %v"
	format[2] = "Hey %v, Saludos"

	return format[rand.Intn(len(format))]
}
