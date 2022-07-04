package main

import (
	"fmt"

	uuid "github.com/kamuiroeru/tgs6s3"
)

func main() {
	uuid4 := uuid.GenerateUuid4()
	fmt.Println(uuid4)
	fmt.Println(uuid4.String())
}
