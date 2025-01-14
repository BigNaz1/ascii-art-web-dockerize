package pkg

import (
	"log"
	"os"
)

/*
GetBannerFile will read the file specified in the parameter passed and
return its string value
*/

func GetBannerFile(s string) string {
	file, err := os.ReadFile(s)
	if err != nil {
		log.Fatal("Error: invalid file or directory for banner file: " + s)
	}

	return string(file)
}
