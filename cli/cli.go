package main

import (
	"fmt"
	"github.com/emiddleton/gads"
	"log"
)

func main() {
	config, err := gads.NewCredentials()
	if err != nil {
		log.Fatal(err)
	}
	bs := gads.NewBudgetService(config.Auth)

	foundBudgets, err := bs.Get(gads.Selector{Fields: []string{"BudgetId", "BudgetName", "Period", "Amount", "DeliveryMethod"}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nBudgets\n")
	for _, budget := range foundBudgets {
		fmt.Printf("%#v\n", budget)
	}

	// show all Campaigns
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
	fmt.Printf("\nCampaigns\n")
	for _, campaign := range foundCampaigns {
		fmt.Printf("%#v\n", campaign)
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

	fmt.Printf("\nAdGroups\n")
	for _, adGroup := range foundAdGroups {
		fmt.Printf("%#v\n", adGroup)
	}

  agas := gads.NewAdGroupAdService(config.Auth)
  foundAds, err := agas.Get(
    gads.Selector{
      Fields: []string{
        "AdGroupId",
        "Status",
        "AdGroupCreativeApprovalStatus",
        "AdGroupAdDisapprovalReasons",
        "AdGroupAdTrademarkDisapproved",
      },
      Ordering: []gads.OrderBy{
        {"AdGroupId","ASCENDING"},
        {"Id","ASCENDING"},
      },
    },
  )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAds\n")
	for _, ad := range foundAds {
		fmt.Printf("%#v\n", ad)
	}

}
