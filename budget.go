package gads

import (
	"encoding/xml"
	//  "fmt"
)

// A budgetService holds the connection information for the
// budget service.
type BudgetService struct {
	Auth
}

// NewBudgetService creates a new budgetService
func NewBudgetService(auth *Auth) *BudgetService {
	return &BudgetService{Auth: *auth}
}

// A Budget represents an allotment of money to be spent over a fixed
// period of time.
type Budget struct {
	Id         int64  `xml:"budgetId,omitempty"`           // A unique identifier
	Name       string `xml:"name"`                         // A descriptive name
	Period     string `xml:"period"`                       // The period to spend the budget
	Amount     int64  `xml:"amount>microAmount"`           // The amount in cents
	Delivery   string `xml:"deliveryMethod"`               // The rate at which the budget spent. valid options are STANDARD or ACCELERATED.
	References int64  `xml:"referenceCount,omitempty"`     // The number of campaigns using the budget
	Shared     bool   `xml:"isExplicitlyShared,omitempty"` // If this budget was created to be shared across campaigns
	Status     string `xml:"status,omitempty"`             // The status of the budget. can be ENABLED, REMOVED, UNKNOWN
}

// A BudgetOperations maps operations to the budgets they will be performed
// on.  Budgets operations can be 'ADD', 'REMOVE' or 'SET'
type BudgetOperations map[string][]Budget

func budgetError() (err error) {
	return err
}

// Get returns budgets matching a given selector and the total count of matching budgets.
func (s *BudgetService) Get(selector Selector) (budgets []Budget, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "selector"}
	respBody, err := s.Auth.request(
		budgetServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return budgets, totalCount, err
	}
	getResp := struct {
		Size    int64    `xml:"rval>totalNumEntries"`
		Budgets []Budget `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return budgets, totalCount, err
	}
	return getResp.Budgets, getResp.Size, err
}

// Mutate takes a budgetOperations and creates, modifies or destroys the associated budgets.
func (s *BudgetService) Mutate(budgetOperations BudgetOperations) (budgets []Budget, err error) {
	type budgetOperation struct {
		Action string `xml:"operator"`
		Budget Budget `xml:"operand"`
	}
	operations := []budgetOperation{}
	for action, budgets := range budgetOperations {
		for _, budget := range budgets {
			operations = append(operations,
				budgetOperation{
					Action: action,
					Budget: budget,
				},
			)
		}
	}
	respBody, err := s.Auth.request(
		budgetServiceUrl,
		"mutate",
		struct {
			XMLName xml.Name
			Ops     []budgetOperation `xml:"operations"`
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "mutate",
			},
			Ops: operations,
		},
	)
	if err != nil {
		return budgets, err
	}
	mutateResp := struct {
		Budgets []Budget `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return budgets, err
	}
	return mutateResp.Budgets, err
}
