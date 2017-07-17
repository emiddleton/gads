package gads

import (
	"encoding/xml"
)

type ReportDefinition struct {
	Selector       Selector `xml:"selector"`
	ReportName     string   `xml:"reportName"`
	ReportType     string   `xml:"reportType"`
	DateRangeType  string   `xml:"dateRangeType"`
	DownloadFormat string   `xml:"downloadFormat"`
}

type ReportDefinitionField struct {
	FieldName           string          `xml:"fieldName"`
	DisplayFieldName    string          `xml:"displayFieldName"`
	XmlAttributeName    string          `xml:"xmlAttributeName"`
	FieldType           string          `xml:"fieldType"`
	FieldBehavior       string          `xml:"fieldBehavior"`
	EnumValues          []string        `xml:"enumValues"`
	CanSelect           bool            `xml:"canSelect"`
	CanFilter           bool            `xml:"canFilter"`
	IsEnumType          bool            `xml:"isEnumType"`
	IsBeta              bool            `xml:"isBeta"`
	IsZeroRowCompatible bool            `xml:"isZeroRowCompatible"`
	EnumValuePairs      []EnumValuePair `xml:"enumValuePairs"`
}

type EnumValuePair struct {
	EnumValue        string `xml:"enumValue"`
	EnumDisplayValue string `xml:"enumDisplayvalue"`
}

type ReportDefinitionService struct {
	Auth
}

type ReportDefinitionRequest struct {
	ReportType string `xml:"reportType"`
}

func NewReportDefinitionService(auth *Auth) *ReportDefinitionService {
	return &ReportDefinitionService{Auth: *auth}
}

func (s *ReportDefinitionService) GetReportFields(report string) (fields []ReportDefinitionField, err error) {
	respBody, err := s.Auth.request(
		reportDefinitionServiceUrl,
		"get",
		struct {
			XMLName    xml.Name
			ReportType string `xml:"reportType"`
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "getReportFields",
			},
			ReportType: report,
		},
		nil,
	)
	if err != nil {
		return fields, err
	}
	getResp := struct {
		Fields []ReportDefinitionField `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return fields, err
	}
	return getResp.Fields, err
}
