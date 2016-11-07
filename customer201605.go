package gads

import "encoding/xml"

type CustomerService201605 struct {
	Auth
}

func NewCustomerService201605(auth *Auth) *CustomerService201605 {
	return &CustomerService201605{Auth: *auth}
}

func (s *CustomerService201605) Get() (customer *Customer, err error) {
	respBody, err := s.Auth.request(
		customerService201605Url,
		"getCustomers",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: mcm201605Url,
				Local: "get",
			},
		},
		nil,
	)
	if err != nil {
		return customer, err
	}

	getResp := struct {
		Customer *Customer `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return customer, err
	}
	return getResp.Customer, err
}
