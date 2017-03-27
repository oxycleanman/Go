package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a value: ")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	fmt.Println("You entered: " + text)
}
