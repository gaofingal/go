package main

import "fmt"

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}

		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {

	result, errorMsg := Divide(100, 10)
	if errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	_, errorMsg = Divide(100, 0)
	if errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
