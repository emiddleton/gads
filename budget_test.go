package gads

import (
	"testing"
)

func testBudgetService(t *testing.T) (service *BudgetService) {
	return &BudgetService{Auth: testAuthSetup(t)}
}

func testBudget(t *testing.T) (Budget, func()) {
	s := testBudgetService(t)
	budgets, err := s.Mutate(
		BudgetOperations{
			"ADD": {
				Budget{
					Name:     "testbudget " + rand_str(10),
					Period:   "DAILY",
					Amount:   50000000,
					Delivery: "STANDARD",
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	cleanupBudget := func() {
		_, err = s.Mutate(BudgetOperations{"REMOVE": budgets})
		if err != nil {
			t.Error(err)
		}
	}
	return budgets[0], cleanupBudget
}

func TestBudget(t *testing.T) {
	s := testBudgetService(t)
	budgets, err := s.Mutate(
		BudgetOperations{
			"ADD": {
				Budget{
					Name:     "testbudget " + rand_str(10),
					Period:   "DAILY",
					Amount:   50000000,
					Delivery: "STANDARD",
				},
				Budget{
					Name:     "test budget " + rand_str(10),
					Period:   "DAILY",
					Amount:   50000000,
					Delivery: "STANDARD",
				},
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	defer func(bs []Budget) {
		_, err = s.Mutate(BudgetOperations{"REMOVE": bs})
		if err != nil {
			t.Error(err)
		}
	}(budgets)

	foundBudgets, _, err := s.Get(
		Selector{
			Fields: []string{
				"BudgetId",
				"BudgetName",
				"Period",
				"Amount",
				"DeliveryMethod",
				"BudgetReferenceCount",
				"IsBudgetExplicitlyShared",
				"BudgetStatus",
			},
			Predicates: []Predicate{
				{"Amount", "LESS_THAN_EQUALS", []string{"500000000"}},
				{"BudgetStatus", "EQUALS", []string{"ENABLED"}},
			},
			Ordering: []OrderBy{
				{"BudgetId", "ASCENDING"},
				{"Amount", "ASCENDING"},
			},
			Paging: &Paging{
				Offset: 0,
				Limit:  100,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("found %d budgets\n", len(foundBudgets))
	for _, b := range budgets {
		func(budget Budget) {
			for _, foundBudget := range foundBudgets {
				if foundBudget.Id == budget.Id {
					return
				}
			}
			t.Errorf("budget %d not found in \n%#v\n", budget.Id, foundBudgets)
		}(b)
	}
}
