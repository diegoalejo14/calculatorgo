package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*func main() {
	operation := readInput()
	operator := readInput()
	c := calc{}
	fmt.Println(c.calculate(operation, operator))
}*/

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func (calc) calculate(values string, operator string) int {
	valuesArr := strings.Split(values, operator)
	a := parse(valuesArr[0])
	b := parse(valuesArr[1])
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		fmt.Println("Operation isn't defined")
		return -1
	}
}
func parse(value string) int {
	operator1, error1 := strconv.Atoi(value)
	if error1 != nil {
		fmt.Printf("Problem in convert process of " + value)
	}
	return operator1
}

type calc struct {
}
