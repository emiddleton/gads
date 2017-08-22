package gads

import "encoding/xml"

type CustomerService struct {
	Auth
}

type Customer struct {
	ID                         int64                      `xml:"customerId,omitempty"`
	CurrencyCode               string                     `xml:"currencyCode,omitempty"`
	DateTimeZone               string                     `xml:"dateTimeZone,omitempty"`
	DescriptiveName            string                     `xml:"descriptiveName,omitempty"`
	CanManageClients           bool                       `xml:"canManageClients,omitempty"`
	IsTestAccount              bool                       `xml:"testAccount,omitempty"`
	AutoTaggingEnabled         bool                       `xml:"autoTaggingEnabled,omitempty"`
	TrackingURLTemplate        string                     `xml:"trackingUrlTemplate,omitempty"`
	ConversionTrackingSettings ConversionTrackingSettings `xml:"conversionTrackingSettings"`
	RemarketingSettings        RemarketingSettings        `xml:"remarketingSettings"`
}

type RemarketingSettings struct {
	Snippet string `xml:"snippet"`
}

func NewCustomerService(auth *Auth) *CustomerService {
	return &CustomerService{Auth: *auth}
}

// GetCustomers Important Notes:
// Starting with v201607, if clientCustomerId is specified in the request header, only details of that customer will be returned.
// To do this for prior versions, use the get() method instead.
func (s *CustomerService) GetCustomers() (customers []Customer, err error) {
	respBody, err := s.Auth.request(
		customerServiceUrl,
		"getCustomers",
		struct {
			XMLName xml.Name
		}{
			XMLName: xml.Name{
				Space: baseMcmUrl,
				Local: "getCustomers",
			},
		},
		nil,
	)
	if err != nil {
		return customers, err
	}

	getResp := struct {
		Customers []Customer `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return customers, err
	}
	return getResp.Customers, nil
}
