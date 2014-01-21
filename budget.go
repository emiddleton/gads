package gads

import (
	"encoding/xml"
	//  "fmt"
)

var (
	BUDGET_SERVICE_URL = "https://adwords.google.com/api/adwords/cm/v201309/BudgetService"
)

type budgetService struct {
	Auth
}

func NewBudgetService(auth Auth) *budgetService {
	return &budgetService{Auth: auth}
}

type Budget struct {
	Id       int64  `xml:"budgetId,omitempty"`
	Name     string `xml:"name"`
	Period   string `xml:"period"`
	Amount   int64  `xml:"amount>microAmount"`
	Delivery string `xml:"deliveryMethod"`
	Shared   bool   `xml:"isExplicitlyShared,omitempty"`
	Status   string `xml:"status,omitempty"`
}

type BudgetOperations map[string][]Budget

func budgetError() (err error) {
	return err
}

func (s *budgetService) Get(selector Selector) (budgets []Budget, err error) {
	selector.XMLName = xml.Name{"", "selector"}
	respBody, err := s.Auth.Request(
		BUDGET_SERVICE_URL,
		"get",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 get"`
			Sel     Selector
		}{Sel: selector},
	)
	if err != nil {
		return budgets, err
	}
	getResp := struct {
		Size    int64    `xml:"rval>totalNumEntries"`
		Budgets []Budget `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return budgets, err
	}
	return getResp.Budgets, err
}

func (s *budgetService) Mutate(budgetOperations BudgetOperations) (budgets []Budget, err error) {
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
	respBody, err := s.Auth.Request(
		BUDGET_SERVICE_URL,
		"mutate",
		struct {
			XMLName xml.Name          `xml:"https://adwords.google.com/api/adwords/cm/v201309 mutate"`
			Ops     []budgetOperation `xml:"operations"`
		}{Ops: operations},
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
