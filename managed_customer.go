package gads

import "encoding/xml"

type ManagedCustomerService struct {
	Auth
}

type ManagedCustomer struct {
	Name             string `xml:"name"`
	CompanyName      string `xml:"companyName"`
	CustomerID       string `xml:"customerId"`
	CanManageClients bool   `xml:"canManageClients"`
	CurrencyCode     string `xml:"currencyCode"`
	DateTimeZone     string `xml:"dateTimeZone"`
	IsTestAccount    bool   `xml:"testAccount"`
}

func NewManagedCustomerService(auth *Auth) *ManagedCustomerService {
	return &ManagedCustomerService{Auth: *auth}
}

func (s *ManagedCustomerService) GetCustomers(selector Selector, customerID string) (customers []ManagedCustomer, totalCount int64, err error) {
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
		&Options{CustomerID: customerID},
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
