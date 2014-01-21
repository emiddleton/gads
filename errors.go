package gads

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type OperationError struct {
	Code      int64  `xml:"OperationError>Code"`
	Details   string `xml:"OperationError>Details"`
	ErrorCode string `xml:"OperationError>ErrorCode"`
	Message   string `xml:"OperationError>Message"`
}

type BudgetError struct {
	Path    string `xml:"fieldPath"`
	String  string `xml:"errorString"`
	Trigger string `xml:"trigger"`
	Type    string `xml:"ApiError.Type"`
	Reason  string `xml:"reason"`
}

type ApiExceptionFault struct {
	Message string        `xml:"message"`
	Type    string        `xml:"ApplicationException.Type"`
	Errors  []BudgetError `xml:"errors"`
}

type ErrorsType struct {
	ApiExceptionFaults []ApiExceptionFault `xml:"ApiExceptionFault"`
}

type Fault struct {
	XMLName     xml.Name   `xml:"Fault"`
	FaultCode   string     `xml:"faultcode"`
	FaultString string     `xml:"faultstring"`
	Errors      ErrorsType `xml:"detail"`
}

func (f *ErrorsType) Error() string {
	errors := []string{}
	for _, e := range f.ApiExceptionFaults {
		errors = append(errors, fmt.Sprintf("%s", e.Message))
	}
	return strings.Join(errors, "\n")
}
