package gads

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type BatchJobService struct {
	Auth
}

type BatchJobPage struct {
	TotalNumEntries int        `xml:"rval>totalNumEntries"`
	BatchJobs       []BatchJob `xml:"rval>entries"`
}

type BatchJobOperations struct {
	BatchJobOperations []BatchJobOperation
}

type Operation struct {
	Operator string      `xml:"https://adwords.google.com/api/adwords/cm/v201609 operator"`
	Operand  interface{} `xml:"operand"`
	Xsi_type string      `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr,omitempty"`
}

type BatchJobOperation struct {
	Operator string   `xml:"operator"`
	Operand  BatchJob `xml:"operand"`
}

type BatchJob struct {
	Id               int64                    `xml:"id,omitempty" json:",string"`
	Status           string                   `xml:"status,omitempty"`
	ProgressStats    *ProgressStats           `xml:"progressStats,omitempty"`
	UploadUrl        *TemporaryUrl            `xml:"uploadUrl,omitempty"`
	DownloadUrl      *TemporaryUrl            `xml:"downloadUrl,omitempty"`
	ProcessingErrors *BatchJobProcessingError `xml:"processingErrors,omitempty"`
}

type TemporaryUrl struct {
	Url        string `xml:"url"`
	Expiration string `xml:"expiration,"`
}

type BatchJobProcessingError struct {
	FieldPath   string `xml:"fieldPath"`
	Trigger     string `xml:"trigger"`
	ErrorString string `xml:"errorString"`
	Reason      string `xml:"reason"`
}

type ProgressStats struct {
	NumOperationsExecuted    int64 `xml:"numOperationsExecuted" json:",string"`
	NumOperationsSucceeded   int64 `xml:"numOperationsSucceeded" json:",string"`
	EstimatedPercentExecuted int   `xml:"estimatedPercentExecuted"`
	NumResultsWritten        int64 `xml:"numResultsWritten" json:",string"`
}

type MutateResults struct {
	Result    MutateResult   `xml:"result"`
	ErrorList []MutateErrors `xml:"errorList"`
	Index     int            `xml:"index"`
}

type MutateErrors struct {
	Errors EntityError `xml:"errors"`
}

type MutateResult interface{}

func NewBatchJobService(auth *Auth) *BatchJobService {
	return &BatchJobService{Auth: *auth}
}

// Get queries the status of existing BatchJobs
//
//	Example
//
//	batchJobs, err := batchJobService.Get(
// 		gads.Selector{
//			Fields: []string{
//				"Id",
//				"Status",
//				"DownloadUrl",
//				"ProcessingErrors",
//				"ProgressStats",
//			},
//			Predicates: []gads.Predicate{
//				{"Id", "EQUALS", []string{strconv.FormatInt(jobId, 10)}},
//			},
//		},
//	)
//
// 	https://developers.google.com/adwords/api/docs/reference/v201609/BatchJobService#get
func (s *BatchJobService) Get(selector Selector) (batchJobPage BatchJobPage, err error) {

	selector.XMLName = xml.Name{baseUrl, "selector"}
	respBody, err := s.Auth.request(
		batchJobServiceUrl,
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
		nil,
	)
	if err != nil {
		return batchJobPage, err
	}

	err = xml.Unmarshal([]byte(respBody), &batchJobPage)

	return batchJobPage, err
}

// Mutate allows you to create or update a BatchJob
//
//	Example
//
//	resp, err := batchJobService.Mutate(
// 		gads.BatchJobOperations{
//			BatchJobOperations: []gads.BatchJobOperation{
//				gads.BatchJobOperation{
//					Operator: "ADD",
//					Operand: gads.BatchJob{},
//				},
//			},
//		},
//	)
//
// 	https://developers.google.com/adwords/api/docs/reference/v201609/BatchJobService#mutate
func (s *BatchJobService) Mutate(batchJobOperations BatchJobOperations) (batchJobs []BatchJob, err error) {

	mutation := struct {
		XMLName xml.Name
		Ops     []BatchJobOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: batchJobOperations.BatchJobOperations}
	respBody, err := s.Auth.request(batchJobServiceUrl, "mutate", mutation, nil)
	if err != nil {
		return batchJobs, err
	}
	mutateResp := struct {
		BatchJobs []BatchJob `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return batchJobs, err
	}

	return mutateResp.BatchJobs, err
}

func (s *BatchJobService) Query() {

}

func (mr *MutateResults) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			tag := start.Name.Local
			switch tag {
			case "index":
				if err := dec.DecodeElement(&mr.Index, &start); err != nil {
					return err
				}
			case "errorList":
				if err := dec.DecodeElement(&mr.ErrorList, &start); err != nil {
					return err
				}
			case "AdGroup":
				ag := AdGroup{}
				err := dec.DecodeElement(&ag, &start)
				if err != nil {
					return err
				}
				mr.Result = ag
			case "AdGroupAd":
				aga := AdGroupAds{}
				err := dec.DecodeElement(&aga, &start)
				if err != nil {
					return err
				}
				mr.Result = aga
			case "AdGroupAdLabel":
				agal := AdGroupAdLabel{}
				err := dec.DecodeElement(&agal, &start)
				if err != nil {
					return err
				}
				mr.Result = agal
			case "AdGroupCriterion":
				agc := AdGroupCriterions{}
				err := dec.DecodeElement(&agc, &start)
				if err != nil {
					return err
				}
				mr.Result = agc
			case "AdGroupCriterionLabel":
				agcl := AdGroupCriterionLabel{}
				err := dec.DecodeElement(&agcl, &start)
				if err != nil {
					return err
				}
				mr.Result = agcl
			case "AdGroupExtensionSetting":
				cl := AdGroupExtensionSetting{}
				err := dec.DecodeElement(&cl, &start)
				if err != nil {
					return err
				}
				mr.Result = cl
			case "AdGroupLabel":
				agl := AdGroupLabel{}
				err := dec.DecodeElement(&agl, &start)
				if err != nil {
					return err
				}
				mr.Result = agl
			case "Budget":
				b := Budget{}
				err := dec.DecodeElement(&b, &start)
				if err != nil {
					return err
				}
				mr.Result = b
			case "Campaign":
				c := Campaign{}
				err := dec.DecodeElement(&c, &start)
				if err != nil {
					return err
				}
				mr.Result = c
			case "CampaignCriterion":
				cc := CampaignCriterions{}
				err := dec.DecodeElement(&cc, &start)
				if err != nil {
					return err
				}
				mr.Result = cc
			case "CampaignExtensionSetting":
				cl := CampaignExtensionSetting{}
				err := dec.DecodeElement(&cl, &start)
				if err != nil {
					return err
				}
				mr.Result = cl
			case "CampaignLabel":
				cl := CampaignLabel{}
				err := dec.DecodeElement(&cl, &start)
				if err != nil {
					return err
				}
				mr.Result = cl
			case "result":
				break
			default:
				return fmt.Errorf("unknown MutateResults field %s", tag)
			}
		}
	}
	return err
}

// getXsiType validates the schema instance type and returns it since Bulk Mutate requires it to be set
func getXsiType(objectName string) (string, bool) {
	switch {
	case strings.Contains(objectName, "AdGroupExtensionSettingOperation"):
		return "AdGroupExtensionSettingOperation", true
	case strings.Contains(objectName, "AdGroupAdLabelOperation"):
		return "AdGroupAdLabelOperation", true
	case strings.Contains(objectName, "AdGroupAdOperation"):
		return "AdGroupAdOperation", true
	case strings.Contains(objectName, "AdGroupBidModifierOperation"):
		return "AdGroupBidModifierOperation", true
	case strings.Contains(objectName, "AdGroupCriterionLabelOperation"):
		return "AdGroupCriterionLabelOperation", true
	case strings.Contains(objectName, "AdGroupCriterionOperation"):
		return "AdGroupCriterionOperation", true
	case strings.Contains(objectName, "AdGroupLabelOperation"):
		return "AdGroupLabelOperation", true
	case strings.Contains(objectName, "AdGroupOperation"):
		return "AdGroupOperation", true
	case strings.Contains(objectName, "BudgetOperation"):
		return "BudgetOperation", true
	case strings.Contains(objectName, "CampaignAdExtensionOperation"):
		return "CampaignAdExtensionOperation", true
	case strings.Contains(objectName, "CampaignExtensionSettingOperation"):
		return "CampaignExtensionSettingOperation", true
	case strings.Contains(objectName, "CampaignCriterionOperation"):
		return "CampaignCriterionOperation", true
	case strings.Contains(objectName, "CampaignLabelOperation"):
		return "CampaignLabelOperation", true
	case strings.Contains(objectName, "CampaignOperation"):
		return "CampaignOperation", true
	case strings.Contains(objectName, "FeedItemOperation"):
		return "FeedItemOperation", true
	default:
		return "", false
	}
}
