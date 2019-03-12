package gads

import "encoding/xml"

type OfflineConversionService struct {
	Auth
}

type OfflineConversionFeed struct {
	GoogleClickId             string  `xml:"googleClickId"`
	ConversionName            string  `xml:"conversionName"`
	ConversionTime            string  `xml:"conversionTime"`
	ConversionCurrencyCode    string  `xml:"conversionCurrencyCode"`
	ExternalAttributionModel  string  `xml:"externalAttributionModel"`
	ConversionValue           float32 `xml:"conversionValue"`
	ExternalAttributionCredit float32 `xml:"externalAttributionCredit"`
}

type OfflineConversionOperations map[string][]OfflineConversionFeed

func NewOfflineConversionService(auth *Auth) *OfflineConversionService {
	return &OfflineConversionService{Auth: *auth}
}

func (o *OfflineConversionService) Mutate(conversionOperations OfflineConversionOperations) (conversion []OfflineConversionFeed, error error) {
	type conversionOperation struct {
		Action     string                `xml:"operator"`
		Conversion OfflineConversionFeed `xml:"operand"`
	}
	var ops []conversionOperation

	for action, conversion := range conversionOperations {
		for _, con := range conversion {
			ops = append(ops,
				conversionOperation{
					Action:     action,
					Conversion: con,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []conversionOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: ops}
	respBody, err := o.Auth.request(campaignServiceUrl, "mutate", mutation)
	if err != nil {
		return conversion, err
	}
	mutateResp := struct {
		Conversions []OfflineConversionFeed `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return conversion, err
	}
	return mutateResp.Conversions, nil
}
