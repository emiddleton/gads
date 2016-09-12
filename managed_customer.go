package gads

import "encoding/xml"

type ManagedCustomerService struct {
	Auth
	CustomerID string
}

type ManagedCustomer struct {
	ID               string `xml:"customerId"`
	Name             string `xml:"name"`
	CompanyName      string `xml:"companyName"`
	CanManageClients bool   `xml:"canManageClients"`
	CurrencyCode     string `xml:"currencyCode"`
	DateTimeZone     string `xml:"dateTimeZone"`
	IsTestAccount    bool   `xml:"testAccount"`
}

func NewManagedCustomerService(auth *Auth, customerID string) *ManagedCustomerService {
	return &ManagedCustomerService{Auth: *auth, CustomerID: customerID}
}

func (s *ManagedCustomerService) Get(selector Selector) (customers []ManagedCustomer, totalCount int64, err error) {
	selector.XMLName = xml.Name{Space: "", Local: "serviceSelector"}
	respBody, err := s.Auth.request(
		managedCustomerServiceUrl,
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
		&Options{CustomerID: s.CustomerID},
	)
	if err != nil {
		return customers, totalCount, err
	}
	getResp := struct {
		Size             int64             `xml:"rval>totalNumEntries"`
		ManagedCustomers []ManagedCustomer `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return customers, totalCount, err
	}
	return getResp.ManagedCustomers, getResp.Size, err
}
