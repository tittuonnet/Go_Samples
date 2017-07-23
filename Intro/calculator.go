package calculator

import "fmt"

func calc(a, b float32, c string) (bool, float32) {
	//returns status and result
	switch c {
	case "+":
		return true, (a + b)
	case "*":
		return true, (a * b)
	case "-":
		return true, (a - b)
	case "/":
		return true, (a / b)
	default:
		return false, 0
	}
}

func Calculator() {
	status, result := calc(5.8, 10, "+")
	if status == true {
		fmt.Printf("The result of operation is %f\n", result)
	} else {
		fmt.Println("Invalid operation")
	}
	status, result = calc(1205.8, 10, "-")
	if status == true {
		fmt.Printf("The result of operation is %f\n", result)
	} else {
		fmt.Println("Invalid operation")
	}
	status, result = calc(122.118, 10.0012, "*")
	if status == true {
		fmt.Printf("The result of operation is %f\n", result)
	} else {
		fmt.Println("Invalid operation")
	}
	status, result = calc(25, 5, "/")
	if status == true {
		fmt.Printf("The result of operation is %f\n", result)
	} else {
		fmt.Println("Invalid operation")
	}
	status, result = calc(127, 5, "/")
	if status == true {
		fmt.Printf("The result of operation is %f\n", result)
	} else {
		fmt.Println("Invalid operation")
	}
	status, result = calc(127, 5, "A")
	if status == true {
		fmt.Printf("The result of operation is %f\n", result)
	} else {
		fmt.Println("Invalid operation")
	}
}
