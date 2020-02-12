package gads

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DownloadFormatCSV         = "CSV"
	DownloadFormatCSVFOREXCEL = "CSVFOREXCEL"
	DownloadFormatTSV         = "TSV"
	DownloadFormatXML         = "XML"
	DownloadFormatGZIPPED_CSV = "GZIPPED_CSV"
	DownloadFormatGZIPPED_XML = "GZIPPED_XML"

	DateRangeTypeToday            = "TODAY"
	DateRangeTypeYesterday        = "YESTERDAY"
	DateRangeTypeLast7days        = "LAST_7_DAYS"
	DateRangeTypeLastWeek         = "LAST_WEEK"
	DateRangeTypeLastBusinessWeek = "LAST_BUSINESS_WEEK"
	DateRangeTypeThisMonth        = "THIS_MONTH"
	DateRangeTypeLastMonth        = "LAST_MONTH"
	DateRangeTypeAllTime          = "ALL_TIME"
	DateRangeTypeCustomDate       = "CUSTOM_DATE"
	DateRangeTypeLast14Days       = "LAST_14_DAYS"
	DateRangeType30Days           = "LAST_30_DAYS"
	DateRangeTypeThisWeekSunToday = "THIS_WEEK_SUN_TODAY"
	DateRangeTypeThisWeekMonToday = "THIS_WEEK_MON_TODAY"
	DateRangeTypeLastWeekSunSat   = "LAST_WEEK_SUN_SAT"
)

type ReportDownloaderService struct {
	Auth
}

type ReportType struct {
	ReportName     string
	ReportType     string
	DateRangeType  string
	DownloadFormat string
}

func NewReportDownloaderService(auth *Auth) *ReportDownloaderService {
	return &ReportDownloaderService{Auth: *auth}
}

func (r *ReportDownloaderService) GetReport(s Selector, rType ReportType) (respBody []byte, err error) {
	service := reportDownloadServiceUrl

	return r.Auth.report(service, s, rType)
}

func (a *Auth) report(serviceUrl ServiceUrl, s Selector, rType ReportType) (respBody []byte, err error) {
	report := struct {
		XMLName        xml.Name
		Sel            Selector `xml:"selector"`
		ReportName     string   `xml:"reportName,omitempty"`
		ReportType     string   `xml:"reportType,omitempty"`
		DateRangeType  string   `xml:"dateRangeType,omitempty"`
		DownloadFormat string   `xml:"downloadFormat,omitempty"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "reportDefinition",
		},
		Sel:            s,
		ReportName:     rType.ReportName,
		ReportType:     rType.ReportType,
		DateRangeType:  rType.DateRangeType,
		DownloadFormat: rType.DownloadFormat,
	}

	reqBody, err := xml.MarshalIndent(report, "  ", "  ")
	if err != nil {
		return []byte{}, err
	}
	query := `__fmt: CSV
__rdquery: SELECT CampaignId, AdGroupId, Impressions, Clicks, Cost 

FROM ADGROUP_PERFORMANCE_REPORT 

WHERE AdGroupStatus IN [ENABLED, PAUSED] DURING LAST_7_DAYS`

	req, err := http.NewRequest("POST", serviceUrl.String(), bytes.NewReader([]byte(query)))
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("Accept", "multipart/*")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=------------------------12d01fae60c7b559")
	contentLength := fmt.Sprintf("%d", len(reqBody))
	req.Header.Add("Content-Length", contentLength)
	resp, err := a.Client.Do(req)

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf(string(respBody))
		return []byte{}, err
	}
	return []byte(respBody), err
}
