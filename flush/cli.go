package main

import (
	"crypto/rand"
	"github.com/emiddleton/gads"
	"golang.org/x/oauth2"
	"log"
)

func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func main() {
	config, err := gads.NewCredentials(oauth2.NoContext)
	if err != nil {
		log.Fatal(err)
	}

	ags := gads.NewAdGroupService(config.Auth)
	foundAdGroups, err := ags.Get(
		gads.Selector{
			Fields: []string{
				"Id",
				"CampaignId",
				"CampaignName",
				"Name",
				"Status",
				"Settings",
				"ContentBidCriterionTypeGroup",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	for idx, _ := range foundAdGroups {
		foundAdGroups[idx].Status = "REMOVED"
		foundAdGroups[idx].Name = rand_str(20)
	}

	_, err = ags.Mutate(gads.AdGroupOperations{"SET": foundAdGroups})

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
			Ordering: []gads.OrderBy{
				{"Id", "ASCENDING"},
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	for idx, _ := range foundCampaigns {
		foundCampaigns[idx].Status = "REMOVED"
		foundCampaigns[idx].Name = rand_str(20)
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
