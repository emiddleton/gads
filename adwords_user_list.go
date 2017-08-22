package gads

import (
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type AdwordsUserListService struct {
	Auth
}

func NewAdwordsUserListService(auth *Auth) *AdwordsUserListService {
	return &AdwordsUserListService{Auth: *auth}
}

type UserListLogicalRule struct {
	Operator     string     `xml:"operator"` // "ALL", "ANY", "NONE", "UNKNOWN"
	RuleOperands []UserList `xml:"ruleOperands"`
}

type UserListConversionType struct {
	Id       *int64
	Name     string
	Category *string `xml:"category"` // "BOOMERANG_EVENT", "OTHER"
}

// Rule structs
type DateRuleItem struct {
	Key   string `xml:"key>name"`
	Op    string `xml:"op"` // "UNKNOWN", "EQUALS", "NOT_EQUAL", "BEFORE", "AFTER"
	Value string `xml:"string"`
}

type NumberRuleItem struct {
	Key   string `xml:"key>name"`
	Op    string `xml:"op"` // "UNKNOWN", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "EQUALS", "NOT_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL"
	Value int64  `xml:"value"`
}

type StringRuleItem struct {
	Key   string `xml:"key>name"`
	Op    string `xml:"op"` // "UNKNOWN", "CONTAINS", "EQUALS", "STARTS_WITH", "ENDS_WITH", "NOT_EQUAL", "NOT_CONTAIN", "NOT_START_WITH", "NOT_END_WITH"
	Value string `xml:"string"`
}

type RuleItemGroup struct {
	Items []interface{} `xml:"items"`
}

type Rule struct {
	Groups []RuleItemGroup `xml:"groups"`
}

type UserListOperations struct {
	XMLName    xml.Name
	Operations []Operation `xml:"operations"`
}

type MutateMembersOperations struct {
	XMLName    xml.Name
	Operations []Operation `xml:"operations"`
}

// Member holds user list member identifiers.
// https://developers.google.com/adwords/api/docs/reference/v201708/AdwordsUserListService.Member
type Member struct {
	Email       string       `xml:"hashedEmail"`
	MobileID    string       `xml:"mobileId"`
	AddressInfo *AddressInfo `xml:"addressInfo,omitempty"`
}

// AddressInfo is an address identifier of a user list member.
// Accessible for whitelisted customers only.
type AddressInfo struct {
	FirstName   string `xml:"hashedFirstName"`
	LastName    string `xml:"hashedLastName"`
	CountryCode string `xml:"countryCode"`
	ZipCode     string `xml:"zipCode"`
}

type MutateMembersOperand struct {
	UserListId int64    `xml:"userListId"`
	RemoveAll  *bool    `xml:"removeAll"`
	Members    []Member `xml:"membersList"`
}

type UserList struct {
	Id                    int64   `xml:"id,omitempty"`
	Readonly              *bool   `xml:"isReadOnly"`
	Name                  string  `xml:"name"`
	Description           string  `xml:"description"`
	Status                string  `xml:"status"` // membership status "OPEN", "CLOSED"
	IntegrationCode       string  `xml:"integrationCode"`
	AccessReason          string  `xml:"accessReason"`          // account access resson "OWNER", "SHARED", "LICENSED", "SUBSCRIBED"
	AccountUserListStatus string  `xml:"accountUserListStatus"` // if share is still active "ACTIVE", "INACTIVE"
	MembershipLifeSpan    int64   `xml:"membershipLifeSpan"`    // number of days cookie stays on list
	Size                  *int64  `xml:"size,omitempty"`
	SizeRange             *string `xml:"sizeRange,omitempty"`          // size range "LESS_THEN_FIVE_HUNDRED","LESS_THAN_ONE_THOUSAND", "ONE_THOUSAND_TO_TEN_THOUSAND","TEN_THOUSAND_TO_FIFTY_THOUSAND","FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND","ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND","THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND","FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION","ONE_MILLION_TO_TWO_MILLION","TWO_MILLION_TO_THREE_MILLION","THREE_MILLION_TO_FIVE_MILLION","FIVE_MILLION_TO_TEN_MILLION","TEN_MILLION_TO_TWENTY_MILLION","TWENTY_MILLION_TO_THIRTY_MILLION","THIRTY_MILLION_TO_FIFTY_MILLION","OVER_FIFTY_MILLION"
	SizeForSearch         *int64  `xml:"sizeForSearch,omitempty"`      // estimated number of google.com users in this group
	SizeRangeForSearch    *string `xml:"sizeRangeForSearch,omitempty"` // same values as size range but for search
	ListType              *string `xml:"sizeType,omitempty"`           // one of "UNKNOWN", "REMARKETING", "LOGICAL", "EXTERNAL_REMARKETING", "RULE_BASED", "SIMILAR"

	// LogicalUserList
	LogicalRules *[]UserListLogicalRule `xml:"rules,omitempty"`

	// BasicUserList
	ConversionTypes *[]UserListConversionType `xml:"conversionTypes"`

	// RuleUserList
	Rule      *Rule   `xml:"rule,omitempty"`
	StartDate *string `xml:"startDate,omitempty"`
	EndDate   *string `xml:"endDate,omitempty"`

	// SimilarUserList
	SeedUserListId          *int64  `xml:"seedUserListId,omitempty"`
	SeedUserListName        *string `xml:"seedUserListName,omitempty"`
	SeedUserListDescription *string `xml:"seedUserListDescription,omitempty"`
	SeedUserListStatus      *string `xml:"seedUserListStatus,omitempty"`
	SeedListSize            *int64  `xml:"seedListSize,omitempty"`

	// CrmBasedUserList
	OptOutLink *string `xml:"optOutLink,omitempty"`

	// Keep track of xsi:type for mutate and mutateMembers
	Xsi_type string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr,omitempty"`
}

// DoubleClick platform user list mapped to AdWords
//func NewExternalRemarketingUserList() (adwordsUserList UserList) {
//}

//
func NewLogicalUserList(name, description, status, integrationCode string, membershipLifeSpan int64, logicalRules []UserListLogicalRule) (adwordsUserList UserList) {
	return UserList{
		Name:               name,
		Description:        description,
		Status:             status,
		IntegrationCode:    integrationCode,
		MembershipLifeSpan: membershipLifeSpan,
		LogicalRules:       &logicalRules,
		Xsi_type:           "LogicalUserList",
	}
}

func NewBasicUserList(name, description, status, integrationCode string, membershipLifeSpan int64, conversionTypes []UserListConversionType) (adwordsUserList UserList) {
	return UserList{
		Name:               name,
		Description:        description,
		Status:             status,
		IntegrationCode:    integrationCode,
		MembershipLifeSpan: membershipLifeSpan,
		ConversionTypes:    &conversionTypes,
		Xsi_type:           "BasicUserList",
	}
}

func NewDateSpecificRuleUserList(name, description, status, integrationCode string, membershipLifeSpan int64, rule Rule, startDate, endDate string) (adwordsUserList UserList) {
	return UserList{
		Name:               name,
		Description:        description,
		Status:             status,
		IntegrationCode:    integrationCode,
		MembershipLifeSpan: membershipLifeSpan,
		Rule:               &rule,
		StartDate:          &startDate,
		EndDate:            &endDate,
		Xsi_type:           "DateSpecificRuleUserList",
	}
}

func NewExpressionRuleUserList(name, description, status, integrationCode string, membershipLifeSpan int64, rule Rule) (adwordsUserList UserList) {
	return UserList{
		Name:               name,
		Description:        description,
		Status:             status,
		IntegrationCode:    integrationCode,
		MembershipLifeSpan: membershipLifeSpan,
		Rule:               &rule,
		Xsi_type:           "ExpressionRuleUserList",
	}
}

func NewSimilarUserList(name, description, status, integrationCode string, membershipLifeSpan int64) (adwordsUserList UserList) {
	return UserList{
		Name:               name,
		Description:        description,
		Status:             status,
		IntegrationCode:    integrationCode,
		MembershipLifeSpan: membershipLifeSpan,
		Xsi_type:           "SimilarUserList",
	}
}

func NewCrmBasedUserList(name, description string, membershipLifeSpan int64, optOutLink string) (adwordsUserList UserList) {
	return UserList{
		Name:               name,
		Description:        description,
		MembershipLifeSpan: membershipLifeSpan,
		OptOutLink:         &optOutLink,
		Xsi_type:           "CrmBasedUserList",
	}
}

func NewMutateMembersOperand() *MutateMembersOperand {
	return new(MutateMembersOperand)
}

// Get returns an array of adwords user lists and the total number of adwords user lists matching
// the selector.
//
// Example
//
//   ads, totalCount, err := adwordsUserListService.Get(
//     gads.Selector{
//       Fields: []string{
//         "Id",
//         "Name",
//         "Status",
//         "Labels",
//       },
//       Predicates: []gads.Predicate{
//         {"Id", "EQUALS", []string{adGroupId}},
//       },
//     },
//   )
//
// Selectable fields are
//   "Id", "IsReadOnly", "Name", "Description", "Status", "IntegrationCode", "AccessReason",
//   "AccountUserListStatus", "MembershipLifeSpan", "Size", "SizeRange", "SizeForSearch",
//   "SizeRangeForSearch", "ListType"
//
//   BasicUserList
//     "ConversionType"
//
//   LogicalUserList
//     "Rules"
//
//   SimilarUserList
//     "SeedUserListId", "SeedUserListName", "SeedUserListDescription", "SeedUserListStatus",
//     "SeedListSize"
//
// filterable fields are
//   "Id", "Name", "Status", "IntegrationCode", "AccessReason", "AccountUserListStatus",
//   "MembershipLifeSpan", "Size", "SizeForSearch", "ListType"
//
//   SimilarUserList
//     "SeedUserListId", "SeedListSize"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201708/AdwordsUserListService#get
//
func (s AdwordsUserListService) Get(selector Selector) (userLists []UserList, err error) {
	selector.XMLName = xml.Name{Space: baseUrl, Local: "serviceSelector"}
	respBody, err := s.Auth.request(
		adwordsUserListServiceUrl,
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
		nil,
	)
	if err != nil {
		return userLists, err
	}
	getResp := struct {
		Size      int64      `xml:"rval>totalNumEntries"`
		UserLists []UserList `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return userLists, err
	}
	return getResp.UserLists, err
}

// Mutate adds/sets a collection of user lists. Returns a list of User Lists
//
// Example
// 	auls := gads.NewAdwordsUserListService(&config.Auth)
//
//	crmList := gads.NewCrmBasedUserList("Test List", "Just a list to test with", 0, "http://mytest.com/optout")
//
//	ops := gads.UserListOperations{
//		Operations: []gads.Operation{
//			gads.Operation{
//				Operator: "ADD",
//				Operand: crmList,
//			},
//		},
//	}
//
//	resp, err := auls.Mutate(ops)
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201708/AdwordsUserListService#mutate
//
func (s *AdwordsUserListService) Mutate(userListOperations UserListOperations) (adwordsUserLists []UserList, err error) {

	userListOperations.XMLName = xml.Name{
		Space: baseRemarketingUrl,
		Local: "mutate",
	}

	respBody, err := s.Auth.request(adwordsUserListServiceUrl, "mutate", userListOperations, nil)
	if err != nil {
		return adwordsUserLists, err
	}
	mutateResp := struct {
		AdwordsUserLists []UserList `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adwordsUserLists, err
	}

	return mutateResp.AdwordsUserLists, err
}

// Mutate adds/removes members/emails to specified user list.
// Note: Only 1 operation per userListId
//
// Example
// 	mmo := gads.NewMutateMembersOperand()
// 	mmo.UserListId = resp[0].Id

// 	var members []string
// 	members = append(members,"brian@test.com")
// 	members = append(members,"test@test.com")

// 	mmo.Members = members

// 	mutateMembersOperations := gads.MutateMembersOperations{
// 		Operations: []gads.Operation{
// 			gads.Operation{
// 				Operator: "ADD",
// 				Operand: mmo,
// 			},
// 		},
// 	}
//
// 	auls.MutateMembers(mutateMembersOperations)
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201708/AdwordsUserListService#mutateMembers
//
func (s *AdwordsUserListService) MutateMembers(mutateMembersOperations MutateMembersOperations) (adwordsUserLists []UserList, err error) {
	mutateMembersOperations.XMLName = xml.Name{
		Space: baseRemarketingUrl,
		Local: "mutateMembers",
	}

	respBody, err := s.Auth.request(adwordsUserListServiceUrl, "mutateMembers", mutateMembersOperations, nil)
	if err != nil {
		return adwordsUserLists, err
	}
	mutateResp := struct {
		AdwordsUserLists []UserList `xml:"rval>userLists"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adwordsUserLists, err
	}

	return mutateResp.AdwordsUserLists, err
}

// MarshalXML is custom XML marshalling logc for the MutateMembersOperand object
func (mmo MutateMembersOperand) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	mmo.encodeAndNormalize()
	e.EncodeToken(start)
	e.EncodeElement(&mmo.UserListId, xml.StartElement{Name: xml.Name{Space: baseRemarketingUrl, Local: "userListId"}})
	e.EncodeElement(&mmo.Members, xml.StartElement{Name: xml.Name{Space: baseRemarketingUrl, Local: "members"}})
	e.EncodeToken(start.End())
	return nil
}

func (mmo *MutateMembersOperand) encodeAndNormalize() {
	for _, member := range mmo.Members {
		h256 := sha256.New()
		io.WriteString(h256, strings.ToLower(strings.TrimSpace(member.Email)))
		member.Email = fmt.Sprintf("%x", h256.Sum(nil))

		// https://developers.google.com/adwords/api/docs/reference/v201708/AdwordsUserListService.AddressInfo
		if member.AddressInfo != nil {
			addr := member.AddressInfo

			// First Name
			h256 = sha256.New()
			io.WriteString(h256, normalize(addr.FirstName))
			addr.FirstName = fmt.Sprintf("%x", h256.Sum(nil))

			// Last Name
			h256 = sha256.New()
			io.WriteString(h256, normalize(addr.LastName))
			addr.LastName = fmt.Sprintf("%x", h256.Sum(nil))
		}
	}
}

func normalize(in string) string {
	in = strings.TrimSpace(in)
	in = strings.Replace(in, " ", "", -1)
	return strings.ToLower(in)
}
