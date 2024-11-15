package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// 1. Create a sentinel error to represent an invalid ID. In main, use errors.Is to
// check for the sentinel error, and print a message when it is found.
var ErrInvalidID = errors.New("invalid ID")

// 2. Define a custom error type to represent an empty field error. This error should include
// the name of the empty Employee  field. In main, use errors.As to check for this error. Print
// out a message that includes the field name

type EmptyFieldError struct {
	FieldName string
}

func (e EmptyFieldError) Error() string {
	return e.FieldName
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %+v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d %v", count, emp)
		if err != nil {
			switch err := err.(type) {
			case interface{ Unwrap() []error }:
				allErrors := err.Unwrap()
				var messages []string
				for _, e := range allErrors {
					messages = append(messages, ProcessError(e, emp))
				}
				message = message + "allErrors: " + strings.Join(messages, ",")
			default:
				message = message + "error: " + ProcessError(err, emp)
			}
		}
		fmt.Println(message)
	}
}

const data = `
		{
			"id": "ABCD-123",
			"first_name": "Bob",
			"last_name": "Bobson",
			"title": "Senior Manager"
		}
		{
			"id": "XYZ-123",
			"first_name": "Mary",
			"last_name": "Maryson",
			"title": "Vice President"
		}
		{
			"id": "HLXO-829",
			"first_name": "Pierce",
			"last_name": "",
			"title": "Intern"
		}
		{
			"id": "BOTX-263",
			"first_name": "",
			"last_name": "Garciason",
			"title": "Manager"
		}
		{
			"id": "MOXW-821",
			"first_name": "Franklin",
			"last_name": "Wantanbe",
			"title": ""
		}
		{
			"id": "",
			"first_name": "Shelly",
			"last_name": "Shellson",
			"title": "CEO"
		}
		{
			"id": "",
			"first_name": "",
			"last_name": "Shellson",
			"title": "CEO"
		}
	`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var validID = regexp.MustCompile(`\w{4}-\d{3}`)

func ValidateEmployee(e Employee) error {
	var allErrors []error
	if len(e.ID) == 0 {
		// return errors.New("missing ID")
		// return EmptyFieldError{"ID"}
		allErrors = append(allErrors, EmptyFieldError{"ID"})
	}
	if !validID.MatchString(e.ID) {
		// return errors.New("invalid  ID")
		// return ErrInvalidID
		allErrors = append(allErrors, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		// return errors.New("missing FirstName")
		// return EmptyFieldError{"FirstName"}
		allErrors = append(allErrors, EmptyFieldError{"FirstName"})
	}
	if len(e.LastName) == 0 {
		// return errors.New("missing LastName")
		// return EmptyFieldError{"LastName"}
		allErrors = append(allErrors, EmptyFieldError{"LastName"})
	}
	if len(e.Title) == 0 {
		// return errors.New("missing Title")
		// return EmptyFieldError{"Title"}
		allErrors = append(allErrors, EmptyFieldError{"Title"})
	}
	switch len(allErrors) {
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}

func ProcessError(err error, e Employee) string {
	var fieldErr EmptyFieldError
	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("error: invalid ID %s\n", e.ID)
	} else if errors.As(err, &fieldErr) {
		return fmt.Sprintf("empty field %s \n", fieldErr.FieldName)
	} else {
		return fmt.Sprintf("%v", err)
	}
}
