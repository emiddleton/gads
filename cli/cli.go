package main

import (
	"encoding/json"
	"fmt"
	"github.com/emiddleton/gads"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	config, err := gads.NewCredentials(oauth2.NoContext)
	if err != nil {
		log.Fatal(err)
	}
	bs := gads.NewBudgetService(config.Auth)

	var pageSize int64 = 500
	var offset int64 = 0
	paging := gads.Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	fmt.Printf("\nBudgets\n")
	for {
		foundBudgets, totalCount, err := bs.Get(gads.Selector{
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
			Paging: &paging,
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, budget := range foundBudgets {
			budgetJson, _ := json.MarshalIndent(budget, "", "  ")
			fmt.Printf("  %s\n", string(budgetJson))
		}
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}

	// show all Campaigns
	cs := gads.NewCampaignService(config.Auth)
	offset = 0
	paging = gads.Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	fmt.Printf("\nCampaigns\n")
	for {
		foundCampaigns, totalCount, err := cs.Get(
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
					"AdvertisingChannelType",
					"AdvertisingChannelSubType",
					"Labels",
					"TrackingUrlTemplate",
					"UrlCustomParameters",
				},
				Predicates: []gads.Predicate{
					{"Status", "EQUALS", []string{"PAUSED"}},
				},
				Ordering: []gads.OrderBy{
					{"Id", "ASCENDING"},
				},
				Paging: &paging,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		for _, campaign := range foundCampaigns {
			campaignJson, _ := json.MarshalIndent(campaign, "", "  ")
			fmt.Printf("%s\n", campaignJson)
		}
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}

	ags := gads.NewAdGroupService(config.Auth)
	offset = 0
	paging = gads.Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	fmt.Printf("\nAdGroups\n")
	for {
		foundAdGroups, totalCount, err := ags.Get(
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
				Paging: &paging,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		for _, adGroup := range foundAdGroups {
			adGroupJson, _ := json.MarshalIndent(adGroup, "", "  ")
			fmt.Printf("%#v\n", adGroupJson)
		}
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}

	agas := gads.NewAdGroupAdService(config.Auth)
	offset = 0
	paging = gads.Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	fmt.Printf("\nAds\n")
	for {
		foundAds, totalCount, err := agas.Get(
			gads.Selector{
				Fields: []string{
					"AdGroupId",
					"Status",
					"AdGroupCreativeApprovalStatus",
					"AdGroupAdDisapprovalReasons",
					"AdGroupAdTrademarkDisapproved",
				},
				Ordering: []gads.OrderBy{
					{"AdGroupId", "ASCENDING"},
					{"Id", "ASCENDING"},
				},
				Paging: &paging,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		for _, ad := range foundAds {
			adJson, _ := json.MarshalIndent(ad, "", "  ")
			fmt.Printf("%s\n", adJson)
		}
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}
}
