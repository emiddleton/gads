package gads

import (
	"encoding/xml"
)

type ManagedCustomerService struct {
	Auth
}

func NewManagedCustomerService(auth *Auth) *ManagedCustomerService {
	return &ManagedCustomerService{Auth: *auth}
}

type AccountLabel struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}

type Account struct {
	Name                  string         `xml:"name,omitempty"`
	CanManageClients      bool           `xml:"canManageClients,omitempty"`
	ExcludeHiddenAccounts bool           `xml:"excludeHiddenAccounts,omitempty"`
	CustomerId            string         `xml:"customerId,omitempty"`
	DateTimeZone          string         `xml:"dateTimeZone,omitempty"`
	CurrencyCode          string         `xml:"currencyCode,omitempty"`
	AccountLabels         []AccountLabel `xml:"accountLabels,omitempty"`
}

func (m *ManagedCustomerService) Get(selector Selector) (accounts []Account, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := m.Auth.request(
		managedCustomerServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: mcmUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return accounts, totalCount, err
	}
	getResp := struct {
		Size     int64     `xml:"rval>totalNumEntries"`
		Accounts []Account `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return accounts, totalCount, err
	}
	return getResp.Accounts, getResp.Size, err
}
