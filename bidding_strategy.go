package gads

import (
//	"encoding/xml"
//	"fmt"
)

type BiddingStrategyService struct {
	Auth
}

func NewBiddingStrategyService(auth *Auth) *BiddingStrategyService {
	return &BiddingStrategyService{Auth: *auth}
}

/*
type AdGroupBidModifier struct {
  CampaignId        int64     `xml:"campaignId"`
  AdGroupId         int64     `xml:"adGroupId"`
  Criterion         Criterion `xml:"criterion"`
  BidModifier       float64   `xml:"bidModifier"`
  BidModifierSource string    `xml:"bidModifierSource"`
}
*/
