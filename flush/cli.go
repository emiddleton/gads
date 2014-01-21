package main

import (
	"github.com/emiddleton/gads"
	"log"
)

func main() {
	config, err := gads.NewCredentials()
	if err != nil {
		log.Fatal(err)
	}

	// Remove all Campaigns
	cs := gads.NewCampaignService(config.Auth)
	foundCampaigns, err := cs.Get(
		gads.Selector{
			Fields: []string{
				"Id",
				"BudgetId",
				"Name",
				"Status",
				"ServingStatus",
				"StartDate",
				"EndDate",
				"AdServingOptimizationStatus",
				"Settings",
			},
			Predicates: []gads.Predicate{
				{"Status", "EQUALS", []string{"PAUSED"}},
			},
			Ordering: []gads.OrderBy{
				{"Id", "ASCENDING"},
			},
			Paging: &gads.Paging{
				Offset: 0,
				Limit:  100,
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	for idx, _ := range foundCampaigns {
		foundCampaigns[idx].Status = "DELETED"
	}
	_, err = cs.Mutate(gads.CampaignOperations{"SET": foundCampaigns})
	if err != nil {
		log.Fatal(err)
	}

	// Remove all budgets
	bs := gads.NewBudgetService(config.Auth)
	foundBudgets, err := bs.Get(gads.Selector{Fields: []string{"BudgetId", "BudgetName", "Period", "Amount", "DeliveryMethod"}})
	if err != nil {
		log.Fatal(err)
	}
	_, err = bs.Mutate(gads.BudgetOperations{"REMOVE": foundBudgets})
	if err != nil {
		log.Printf("%#v\n", err)
		log.Fatal(err)
	}
}
