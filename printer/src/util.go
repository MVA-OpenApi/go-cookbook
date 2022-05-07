package printer

import (
	"errors"
	"fmt"
	"os"
)

func CheckIfFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist!")
		return false
	} else {
		fmt.Println(err)
		return false
	}
}

type Names struct {
	Names []string `json:"names"`
}

type Cities struct {
	Cities []string `json:"cities"`
}