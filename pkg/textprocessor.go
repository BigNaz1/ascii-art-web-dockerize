package pkg

import (
	"fmt"
	"strings"
)

func TextProcessor(textInput, font string) (string, error) {
	fmt.Print("The program made it to TextProcessor!")
	err := CheckString(textInput)
	if err != nil {
		fmt.Print("\n")
		return "", err
	}
	input := strings.Split(textInput, "\\n")
	lines := strings.Split(GetBannerFile("banners/"+font+".txt"), "\n")
	result := StoreAsciiArt(input, lines)
	// fmt.Println(result)
	return result, err
}
