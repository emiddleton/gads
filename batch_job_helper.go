package gads

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

type BatchJobHelper struct {
	Auth
}

func NewBatchJobHelper(auth *Auth) *BatchJobHelper {
	return &BatchJobHelper{Auth: *auth}
}

//	UploadBatchJobOperations uploads batch operations to an BatchJob.UploadUrl from BatchJobService.Mutate
//
//	Example
//
//	ago := gads.AdGroupOperations{
//			"ADD": {
//				gads.AdGroup{
// 					Name:       "test ad group " + rand_str(10),
// 					Status:     "PAUSED",
// 					CampaignId: campaignId,
// 				},
// 				gads.AdGroup{
// 					Name:       "test ad group " + rand_str(10),
// 					Status:     "PAUSED",
// 					CampaignId: campaignId,
// 				},
// 			},
// 		}
//
//	var operations []interface{}
//	operations = append(operations, ago)
//	err = batchJobHelper.UploadBatchJobOperations(operations, UploadUrl)
//
//
//	https://developers.google.com/adwords/api/docs/guides/batch-jobs?hl=en#upload_operations_to_the_upload_url
func (s *BatchJobHelper) UploadBatchJobOperations(jobOperations []interface{}, url TemporaryUrl) (err error) {

	var operations []Operation
	for _, operation := range jobOperations {
		if operationType, valid := getXsiType(reflect.ValueOf(operation).Type().String()); valid {
			switch reflect.TypeOf(operation).Kind() {
			case reflect.Map:
				ops := reflect.ValueOf(operation)

				keys := ops.MapKeys()

				for _, action := range keys {
					jobs := ops.MapIndex(action)

					for i := 0; i < jobs.Len(); i++ {

						operations = append(operations,
							Operation{
								Operator: action.String(),
								Operand:  jobs.Index(i).Interface(),
								Xsi_type: operationType,
							},
						)
					}
				}
			}
		}
	}

	if len(operations) > 0 {
		mutation := struct {
			XMLName xml.Name
			Ops     []Operation `xml:"operations"`
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "mutate",
			},
			Ops: operations,
		}

		client := &http.Client{}

		// Need to get the upload url
		req, err := http.NewRequest("POST", url.Url, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/xml")
		req.Header.Set("Content-Length", "0")
		req.Header.Set("x-goog-resumable", "start")

		response, err := client.Do(req)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		// If we got a valid upload url it will be 201
		if response.StatusCode != http.StatusCreated {
			respBody, err := ioutil.ReadAll(response.Body)

			if err != nil {
				return err
			}
			return errors.New(fmt.Sprintf("Invalid response received. %v received. Body: %v", response.StatusCode, string(respBody)))
		}

		location := response.Header.Get("Location")

		reqBody, err := xml.MarshalIndent(mutation, "  ", "  ")
		bodyLength := len(reqBody)

		if err != nil {
			return err
		}

		req, err = http.NewRequest("PUT", location, bytes.NewReader(reqBody))

		if err != nil {
			return err
		}

		// Set headers for incremental upload
		req.Header.Set("Content-Type", "application/xml")
		req.Header.Set("Content-Length", string(bodyLength))
		req.Header.Set("Content-Range", fmt.Sprintf("bytes 0-%v/%v", bodyLength-1, bodyLength))

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		// Added some logging/"poor man's" debugging to inspect outbound SOAP requests
		if level := os.Getenv("DEBUG"); level != "" {
			fmt.Printf("response ->\n%s\n", string(respBody))
		}

		// resp seems to only return 200's and there is no error handling, but if we happen to get invalid status lets try to do something with it
		if resp.StatusCode != http.StatusOK {
			return errors.New("Non-200 response returned Body: " + string(respBody))
		}
	}

	return err
}

//	DownloadBatchJob download batch operations from an BatchJob.DownloadUrl from BatchJobService.Get
//
//	Example
//
//	mutateResult, err := batchJobHelper.DownloadBatchJob(*batchJobs.BatchJobs[0].DownloadUrl)
//
//	https://developers.google.com/adwords/api/docs/guides/batch-jobs?hl=en#download_the_batch_job_results_and_check_for_errors
func (s *BatchJobHelper) DownloadBatchJob(url TemporaryUrl) (mutateResults []MutateResults, err error) {
	resp, err := http.Get(url.Url)

	if err != nil {
		return mutateResults, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	// Added some logging/"poor man's" debugging to inspect outbound SOAP requests
	if level := os.Getenv("DEBUG"); level != "" {
		fmt.Printf("response ->\n%s\n", string(respBody))
	}

	soapResp := struct {
		MutateResults []MutateResults `xml:"rval"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &soapResp)
	if err != nil {
		return mutateResults, err
	}

	return soapResp.MutateResults, err
}
