package gads

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	version               = "v201708"
	rootUrl               = "https://adwords.google.com/api/adwords/cm/"
	baseUrl               = "https://adwords.google.com/api/adwords/cm/" + version
	rootMcmUrl            = "https://adwords.google.com/api/adwords/mcm/"
	baseMcmUrl            = "https://adwords.google.com/api/adwords/mcm/" + version
	rootRemarketingUrl    = "https://adwords.google.com/api/adwords/rm/"
	baseRemarketingUrl    = "https://adwords.google.com/api/adwords/rm/" + version
	rootReportDownloadUrl = "https://adwords.google.com/api/adwords/reportdownload/"
	baseReportDownloadUrl = "https://adwords.google.com/api/adwords/reportdownload/" + version
	rootTrafficUrl        = "https://adwords.google.com/api/adwords/o/"
	baseTrafficUrl        = "https://adwords.google.com/api/adwords/o/" + version
)

type ServiceUrl struct {
	Url  string
	Name string
}

// exceptions
var (
	ERROR_NOT_YET_IMPLEMENTED = fmt.Errorf("Not yet implemented")
)

var (
	configJson = flag.String("config_json", "./config.json", "API credentials")

	// service urls
	adGroupAdServiceUrl               = ServiceUrl{baseUrl, "AdGroupAdService"}
	adGroupBidModifierServiceUrl      = ServiceUrl{baseUrl, "AdGroupBidModifierService"}
	adGroupCriterionServiceUrl        = ServiceUrl{baseUrl, "AdGroupCriterionService"}
	adGroupExtensionSettingServiceUrl = ServiceUrl{baseUrl, "AdGroupExtensionSettingService"}
	adGroupFeedServiceUrl             = ServiceUrl{baseUrl, "AdGroupFeedService"}
	adGroupServiceUrl                 = ServiceUrl{baseUrl, "AdGroupService"}
	adParamServiceUrl                 = ServiceUrl{baseUrl, "AdParamService"}
	adwordsUserListServiceUrl         = ServiceUrl{baseRemarketingUrl, "AdwordsUserListService"}
	batchJobServiceUrl                = ServiceUrl{baseUrl, "BatchJobService"}
	biddingStrategyServiceUrl         = ServiceUrl{baseUrl, "BiddingStrategyService"}
	budgetOrderServiceUrl             = ServiceUrl{baseUrl, "BudgetOrderService"}
	budgetServiceUrl                  = ServiceUrl{baseUrl, "BudgetService"}
	campaignExtensionSettingUrl       = ServiceUrl{baseUrl, "CampaignExtensionSettingService"}
	campaignCriterionServiceUrl       = ServiceUrl{baseUrl, "CampaignCriterionService"}
	campaignFeedServiceUrl            = ServiceUrl{baseUrl, "CampaignFeedService"}
	campaignServiceUrl                = ServiceUrl{baseUrl, "CampaignService"}
	campaignSharedSetServiceUrl       = ServiceUrl{baseUrl, "CampaignSharedSetService"}
	constantDataServiceUrl            = ServiceUrl{baseUrl, "ConstantDataService"}
	conversionTrackerServiceUrl       = ServiceUrl{baseUrl, "ConversionTrackerService"}
	customerFeedServiceUrl            = ServiceUrl{baseUrl, "CustomerFeedService"}
	customerServiceUrl                = ServiceUrl{baseMcmUrl, "CustomerService"}
	customerSyncServiceUrl            = ServiceUrl{baseUrl, "CustomerSyncService"}
	dataServiceUrl                    = ServiceUrl{baseUrl, "DataService"}
	experimentServiceUrl              = ServiceUrl{baseUrl, "ExperimentService"}
	feedItemServiceUrl                = ServiceUrl{baseUrl, "FeedItemService"}
	feedMappingServiceUrl             = ServiceUrl{baseUrl, "FeedMappingService"}
	feedServiceUrl                    = ServiceUrl{baseUrl, "FeedService"}
	geoLocationServiceUrl             = ServiceUrl{baseUrl, "GeoLocationService"}
	labelServiceUrl                   = ServiceUrl{baseUrl, "LabelService"}
	locationCriterionServiceUrl       = ServiceUrl{baseUrl, "LocationCriterionService"}
	managedCustomerServiceUrl         = ServiceUrl{baseMcmUrl, "ManagedCustomerService"}
	mediaServiceUrl                   = ServiceUrl{baseUrl, "MediaService"}
	mutateJobServiceUrl               = ServiceUrl{baseUrl, "MutateJobService"}
	offlineConversionFeedServiceUrl   = ServiceUrl{baseUrl, "OfflineConversionFeedService"}
	reportDefinitionServiceUrl        = ServiceUrl{baseUrl, "ReportDefinitionService"}
	reportDownloadServiceUrl          = ServiceUrl{baseReportDownloadUrl, ""}
	sharedCriterionServiceUrl         = ServiceUrl{baseUrl, "SharedCriterionService"}
	sharedSetServiceUrl               = ServiceUrl{baseUrl, "SharedSetService"}
	targetingIdeaServiceUrl           = ServiceUrl{baseTrafficUrl, "TargetingIdeaService"}
	trafficEstimatorServiceUrl        = ServiceUrl{baseTrafficUrl, "TrafficEstimatorService"}
)

func (s ServiceUrl) String() string {
	if s.Name != "" {
		return s.Url + "/" + s.Name
	}
	return s.Url
}

type Auth struct {
	CustomerId     string
	DeveloperToken string
	UserAgent      string
	PartialFailure bool
	Testing        *testing.T   `json:"-"`
	Client         *http.Client `json:"-"`
}

//
// Selector structs
//
type DateRange struct {
	Min string `xml:"min"`
	Max string `xml:"max"`
}

type Predicate struct {
	Field    string   `xml:"field"`
	Operator string   `xml:"operator"`
	Values   []string `xml:"values"`
}

type OrderBy struct {
	Field     string `xml:"field"`
	SortOrder string `xml:"sortOrder"`
}

type Paging struct {
	Offset int64 `xml:"https://adwords.google.com/api/adwords/cm/v201708 startIndex"`
	Limit  int64 `xml:"https://adwords.google.com/api/adwords/cm/v201708 numberResults"`
}

type Selector struct {
	XMLName    xml.Name
	Fields     []string    `xml:"fields"`
	Predicates []Predicate `xml:"predicates"`
	DateRange  *DateRange  `xml:"dateRange,omitempty"`
	Ordering   []OrderBy   `xml:"ordering"`
	Paging     *Paging     `xml:"paging,omitempty"`
}

type AWQLQuery struct {
	XMLName xml.Name
	Query   string `xml:"query"`
}

type Options struct {
	CustomerID string
}

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.DayOfWeek
// Days of the week.
// MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY
type DayOfWeek string

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.MinuteOfHour
// Minutes in an hour. Currently only 0, 15, 30, and 45 are supported
// ZERO, FIFTEEN, THIRTY, FORTY_FIVE
type MinuteOfHour string

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.GeoRestriction
// A restriction used to determine if the request context's geo should be matched.
// UNKNOWN, LOCATION_OF_PRESENCE
type GeoRestriction string

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.PolicyData
// Approval and policy information attached to an entity.
type PolicyData struct {
	DisapprovalReasons []DisapprovalReason `xml:"https://adwords.google.com/api/adwords/cm/v201708 disapprovalReasons,omitempty"`
	PolicyDataType     string              `xml:"https://adwords.google.com/api/adwords/cm/v201708 PolicyData.Type,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.DisapprovalReason
// Container for information about why an AdWords entity was disapproved.
type DisapprovalReason struct {
	ShortName string `xml:"https://adwords.google.com/api/adwords/cm/v201708 shortName,omitempty"`
}

// error parsers
func selectorError() (err error) {
	return err
}

func (a *Auth) request(serviceUrl ServiceUrl, action string, body interface{}, opts *Options) (respBody []byte, err error) {
	type devToken struct {
		XMLName xml.Name
	}
	type soapReqHeader struct {
		XMLName          xml.Name
		UserAgent        string `xml:"userAgent"`
		DeveloperToken   string `xml:"developerToken"`
		ClientCustomerId string `xml:"clientCustomerId,omitempty"`
		PartialFailure   bool   `xml:"partialFailure,omitempty"`
	}

	type soapReqBody struct {
		Body interface{}
	}

	type soapReqEnvelope struct {
		XMLName xml.Name
		Header  soapReqHeader `xml:"Header>RequestHeader"`
		Body    soapReqBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}

	header := soapReqHeader{
		XMLName:        xml.Name{serviceUrl.Url, "RequestHeader"},
		UserAgent:      a.UserAgent,
		DeveloperToken: a.DeveloperToken,
	}

	if opts != nil && opts.CustomerID != "" {
		header.ClientCustomerId = opts.CustomerID
	} else if a.CustomerId != "" {
		header.ClientCustomerId = a.CustomerId
	}

	reqBody, err := xml.MarshalIndent(
		soapReqEnvelope{
			XMLName: xml.Name{"http://schemas.xmlsoap.org/soap/envelope/", "Envelope"},
			Header:  header,
			Body:    soapReqBody{body},
		},
		"  ", "  ")
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest("POST", serviceUrl.String(), bytes.NewReader(reqBody))
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("Accept", "multipart/*")
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	contentLength := fmt.Sprintf("%d", len(reqBody))
	req.Header.Add("Content-length", contentLength)
	req.Header.Add("SOAPAction", action)
	if a.Testing != nil {
		a.Testing.Logf("request ->\n%s\n%#v\n%s\n", req.URL.String(), req.Header, string(reqBody))
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if a.Testing != nil {
		a.Testing.Logf("respBody ->\n%s\n%s\n", string(respBody), resp.Status)
	}

	type soapRespHeader struct {
		RequestId    string `xml:"requestId"`
		ServiceName  string `xml:"serviceName"`
		MethodName   string `xml:"methodName"`
		Operations   int64  `xml:"operations"`
		ResponseTime int64  `xml:"responseTime"`
	}

	type soapRespBody struct {
		Response []byte `xml:",innerxml"`
	}

	soapResp := struct {
		XMLName xml.Name       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Header  soapRespHeader `xml:"Header>RequestHeader"`
		Body    soapRespBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &soapResp)
	if err != nil {
		return respBody, err
	}
	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 405 || resp.StatusCode == 500 {
		fault := Fault{}
		err = xml.Unmarshal(soapResp.Body.Response, &fault)
		if err != nil {
			return respBody, err
		}
		return soapResp.Body.Response, &fault.Errors
	}
	return soapResp.Body.Response, err
}
