package gads

import "encoding/xml"

type ManagedCustomerService struct {
	Auth
	CustomerID string
}

type ManagedCustomer struct {
	ID                    int64          `xml:"customerId,omitempty"`
	Name                  string         `xml:"name"`
	CanManageClients      bool           `xml:"canManageClients,omitempty"`
	CurrencyCode          string         `xml:"currencyCode"`
	DateTimeZone          string         `xml:"dateTimeZone"`
	IsTestAccount         bool           `xml:"testAccount,omitempty"`
	AccountLabels         []AccountLabel `xml:"accountLabels,omitempty"`
	ExcludeHiddenAccounts bool           `xml:"excludeHiddenAccounts,omitempty"`
}

type ManagedCustomerLink struct {
	ManagerCustomerId      int64  `xml:"managerCustomerId"`
	ClientCustomerId       int64  `xml:"clientCustomerId"`
	LinkStatus             string `xml:"linkStatus"`
	PendingDescriptiveName string `xml:"pendingDescriptiveName"`
	IsHidden               bool   `xml:isHidden"`
}

type ManagedCustomerOperations map[string][]ManagedCustomer

type ManagedCustomerPage struct {
	Size                 int64                 `xml:"rval>totalNumEntries"`
	ManagedCustomers     []ManagedCustomer     `xml:"rval>entries"`
	ManagedCustomerLinks []ManagedCustomerLink `xml:"rval>links"`
}

type AccountLabel struct {
	Id   int64  `xml:"id"`
	Name string `xml:"name"`
}

func NewManagedCustomerService(auth *Auth, customerID string) *ManagedCustomerService {
	return &ManagedCustomerService{Auth: *auth, CustomerID: customerID}
}

func (s *ManagedCustomerService) Get(selector Selector) (managedCustomerPage ManagedCustomerPage, totalCount int64, err error) {
	selector.XMLName = xml.Name{baseMcmUrl, "serviceSelector"}
	respBody, err := s.Auth.request(
		managedCustomerServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseMcmUrl,
				Local: "get",
			},
			Sel: selector,
		},
		&Options{CustomerID: s.CustomerID},
	)
	if err != nil {
		return managedCustomerPage, totalCount, err
	}
	getResp := ManagedCustomerPage{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return managedCustomerPage, totalCount, err
	}
	return getResp, totalCount, nil
}

func (s *ManagedCustomerService) Mutate(managedCustomerOperations ManagedCustomerOperations) (managedCustomers []ManagedCustomer, err error) {
	type managedCustomerOperation struct {
		Action          string          `xml:"https://adwords.google.com/api/adwords/cm/v201708 operator"`
		ManagedCustomer ManagedCustomer `xml:"operand"`
	}

	operations := []managedCustomerOperation{}
	for action, managedCustomers := range managedCustomerOperations {
		for _, managedCustomer := range managedCustomers {
			operations = append(operations,
				managedCustomerOperation{
					Action:          action,
					ManagedCustomer: managedCustomer,
				},
			)
		}
	}

	mutation := struct {
		XMLName xml.Name
		Ops     []managedCustomerOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseMcmUrl,
			Local: "mutate",
		},
		Ops: operations,
	}

	respBody, err := s.Auth.request(managedCustomerServiceUrl, "mutate", mutation, nil)
	if err != nil {
		return managedCustomers, err
	}

	mutateResp := struct {
		ManagedCustomers []ManagedCustomer `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return managedCustomers, err
	}

	return mutateResp.ManagedCustomers, err
}
