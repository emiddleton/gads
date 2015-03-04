package gads

import (
//	"encoding/xml"
//	"fmt"
)

type AdGroupBidModifierService struct {
	Auth
}

func NewAdGroupBidModifierService(auth *Auth) *AdGroupBidModifierService {
	return &AdGroupBidModifierService{Auth: *auth}
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
