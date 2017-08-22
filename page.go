package gads

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.Page
// Contains the results from a get call.
type Page struct {
	TotalNumEntries int    `xml:"https://adwords.google.com/api/adwords/cm/v201708 totalNumEntries,omitempty"`
	PageType        string `xml:"https://adwords.google.com/api/adwords/cm/v201708 Page.Type,omitempty"`
}
