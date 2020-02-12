package gads

import (
	"encoding/xml"
)

type AdService struct {
	Auth
}

func NewAdService(auth *Auth) *AdService {
	return &AdService{Auth: *auth}
}

type AssetLink struct {
	Asset Asset `xml:"asset"`
}

// Media represents an audio, image or video file.
type Ad struct {
	Id                  int64    `xml:"id,omitempty"`
	Url                 string   `xml:"url,omitempty"`
	DisplayUrl          string   `xml:"displayUrl,omitempty"`
	FinalUrls           []string `xml:"finalUrls,omitempty"`
	FinalMobileUrls     []string `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls        []string `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate []string `xml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix      []string `xml:"finalUrlSuffix,omitempty"`
	Type                string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	AdType              string   `xml:"type"` // "DEPRECATED_AD", "IMAGE_AD", "PRODUCT_AD", "TEMPLATE_AD", "TEXT_AD", "THIRD_PARTY_REDIRECT_AD", "DYNAMIC_SEARCH_AD", "CALL_ONLY_AD",
	// "EXPANDED_TEXT_AD", "RESPONSIVE_DISPLAY_AD", "SHOWCASE_AD", "GOAL_OPTIMIZED_SHOPPING_AD", "EXPANDED_DYNAMIC_SEARCH_AD", "GMAIL_AD",
	// "RESPONSIVE_SEARCH_AD", "MULTI_ASSET_RESPONSIVE_DISPLAY_AD", "UNIVERSAL_APP_AD", "UNKNOWN"

	// MultiAssetResponsiveDisplayAd
	LogoImage []AssetLink `xml:"logoImages,omitempty"` // list asset link
}

func (s *AdService) Get(selector Selector) (ads []Ad, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		adServiceUrl,
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
		return ads, totalCount, err
	}
	getResp := struct {
		Size int64 `xml:"rval>totalNumEntries"`
		Ads  []Ad  `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return ads, totalCount, err
	}
	return getResp.Ads, getResp.Size, err
}
