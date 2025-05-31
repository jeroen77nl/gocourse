package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Struct, interfaces, enums
	// CalcualteArea, UserInfo, NewHTTPRequest

	// UPPERCASE
	// constants

	// snake_case
	// variables
	// user_id, first_name, http_request

	// mixedCase
	// variables
	// javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
