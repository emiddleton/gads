package gads

import "encoding/xml"

type CustomerService struct {
	Auth
}

type Customer struct {
	ID                  string `xml:"customer_id,omitempty"`
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

func (s *CustomerService) GetCustomers() (customers []Customer, totalCount int64, err error) {
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
	)
	if err != nil {
		return customers, totalCount, err
	}
	getResp := struct {
		Size      int64      `xml:"rval>totalNumEntries"`
		Customers []Customer `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return customers, totalCount, err
	}
	return getResp.Customers, getResp.Size, err
}
