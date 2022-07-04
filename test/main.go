package main

import (
	"fmt"

	"github.com/kamuiroeru/tgs6s3"
)

func main() {
	uuid4 := tgs6s3.GenerateUuid4()
	fmt.Println(uuid4)
	fmt.Println(uuid4.String())
}
