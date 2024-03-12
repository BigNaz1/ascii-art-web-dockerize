package pkg

import (
	"errors"
	"fmt"
)

func CheckString(s string) error {
	fmt.Print("The program made it to checkstring!")
	if s == "" {
		fmt.Println("String is empty.")
	}
	for _, r := range s {
		if r < 32 || r > 126 {
			return errors.New("Invalid Character:" + string(r))
		}
	}
	return nil
}
