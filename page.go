package gads

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.Page
// Contains the results from a get call.
type Page struct {
	TotalNumEntries int    `xml:"https://adwords.google.com/api/adwords/cm/v201802 totalNumEntries,omitempty"`
	PageType        string `xml:"https://adwords.google.com/api/adwords/cm/v201802 Page.Type,omitempty"`
}
