package language

import "fmt"

func LoadText(fileName string) string {

	fmt.Printf("loading file '%s'.\n", fileName)

	return "this is a sample sentence"
}
