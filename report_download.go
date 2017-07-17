package gads

import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ReportDownloadService struct {
	Auth
}

type reportDefinitionXml struct {
	*ReportDefinition
	XMLName xml.Name
}

func NewReportDownloadService(auth *Auth) *ReportDownloadService {
	return &ReportDownloadService{Auth: *auth}
}

func (s *ReportDownloadService) Get(reportDefinition ReportDefinition) (res interface{}, err error) {
	reportDefinition.Selector.XMLName = xml.Name{baseUrl, "selector"}
	repDef := reportDefinitionXml{
		ReportDefinition: &reportDefinition,
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "reportDefinition",
		},
	}
	body, err := xml.MarshalIndent(repDef, "  ", "  ")
	if err != nil {
		return res, err
	}
	form := url.Values{}
	form.Add("__rdxml", string(body))
	resp, err := s.makeRequest(form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	/*respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return string(respBody), err*/
	return parseReport(resp.Body)
}

func (s *ReportDownloadService) AWQL(awql string, fmt string) (res interface{}, err error) {
	form := url.Values{}
	form.Add("__rdquery", awql)
	form.Add("__fmt", fmt)
	resp, err := s.makeRequest(form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	/*respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return string(respBody), err*/

	// if we didn't get a 200 bubble up the error
	if resp.StatusCode != 200 {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}
		return res, errors.New(string(response))
	}

	return parseReport(resp.Body)
}

// Make our http request using the given form (re-usable for either XML or AWQL)
func (s *ReportDownloadService) makeRequest(form url.Values) (res *http.Response, err error) {
	req, err := http.NewRequest("POST", reportDownloadServiceUrl.Url, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return res, err
	}
	req.Header.Add("developerToken", s.Auth.DeveloperToken)
	req.Header.Add("clientCustomerId", s.Auth.CustomerId)
	req.Header.Add("skipReportHeader", "true")
	req.Header.Add("skipReportSummary", "true")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return s.Client.Do(req)
}

func parseReport(report io.Reader) (collection []map[string]string, err error) {
	reader := csv.NewReader(report)
	header, err := reader.Read()
	if err != nil {
		return collection, err
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return collection, err
		}
		row := make(map[string]string)
		for i := 0; i < len(record); i++ {
			column := header[i]
			row[column] = record[i]
		}
		collection = append(collection, row)
	}
	return collection, err
}
