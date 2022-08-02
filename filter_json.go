package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // io.EOF
		}

		if json.Valid([]byte(line)) {
			fmt.Print(line)
		} else {
			_, err = fmt.Fprint(os.Stderr, line)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
