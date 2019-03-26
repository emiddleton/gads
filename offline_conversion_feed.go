package gads

import "encoding/xml"

type OfflineConversionService struct {
	Auth
}

const (
	uploadConversionAction = "ADD"
)

type OfflineConversionFeed struct {
	GoogleClickId             string  `xml:"googleClickId"`
	ConversionName            string  `xml:"conversionName"`
	ConversionTime            string  `xml:"conversionTime,omitempty"`
	ConversionCurrencyCode    string  `xml:"conversionCurrencyCode,omitempty"`
	ExternalAttributionModel  string  `xml:"externalAttributionModel,omitempty"`
	ConversionValue           float64 `xml:"conversionValue,omitempty"`
	ExternalAttributionCredit float64 `xml:"externalAttributionCredit,omitempty"`
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

	for _, conversion := range conversionOperations {
		for _, con := range conversion {
			ops = append(ops,
				conversionOperation{
					Action:     uploadConversionAction,
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
	respBody, err := o.Auth.request(offlineConversionFeedServiceUrl, "mutate", mutation)
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
