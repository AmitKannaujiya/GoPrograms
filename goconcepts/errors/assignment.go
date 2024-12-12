package errors

import "fmt"

type MyError struct {
	code int
	msg  string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("Error occured : %s, with code : %d", me.msg, me.code)
}

func validateZipCode(zipcode int) *MyError {
	if zipcode < 0 {
		return &MyError{code: 101, msg: "Zipcode is negative"}
	}
	if zipcode < 1000 || zipcode >9999 {
		return &MyError{code: 102, msg: "Zipcode len should be in 4 digits"}
	}
	return nil
}