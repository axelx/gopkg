package gopkg

import (
	"fmt"

	"github.com/axelx/gopkg/internal/sum"
)

func Hello() {
	fmt.Println("Hello, World!")
	fmt.Println("Sum: 1 + 2 = ", sum.Sum(1, 2))
}

func Show() {
	fmt.Println("hello Show func: ", sum.Sum(1, 2))
}
