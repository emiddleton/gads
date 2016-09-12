package gads

import "encoding/xml"

type CustomerService struct {
	Auth
}

type Customer struct {
	ID                  string `xml:"customerId,omitempty"`
	CurrencyCode        string `xml:"currencyCode,omitempty"`
	DateTimeZone        string `xml:"dateTimeZone,omitempty"`
	DescriptiveName     string `xml:"descriptiveName,omitempty"`
	CompanyName         string `xml:"companyName,omitempty"`
	CanManageClients    bool   `xml:"canManageClients,omitempty"`
	IsTestAccount       bool   `xml:"testAccount,omitempty"`
	AutoTaggingEnabled  bool   `xml:"autoTaggingEnabled,omitempty"`
	TrackingURLTemplate string `xml:"trackingUrlTemplate,omitempty"`
}

func NewCustomerService(auth *Auth) *CustomerService {
	return &CustomerService{Auth: *auth}
}

func (s *CustomerService) GetCustomers() (customers []Customer, err error) {
	respBody, err := s.Auth.request(
		customerServiceUrl,
		"getCustomers",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
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
	return getResp.Customers, err
}
