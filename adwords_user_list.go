package gads

import (
	"encoding/xml"
	"fmt"
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
}

type UserList struct {
	Id                    int64   `xml:"id"`
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
	}
}

func NewSimilarUserList(name, description, status, integrationCode string, membershipLifeSpan int64) (adwordsUserList UserList) {
	return UserList{
		Name:               name,
		Description:        description,
		Status:             status,
		IntegrationCode:    integrationCode,
		MembershipLifeSpan: membershipLifeSpan,
	}
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
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdwordsUserListService#get
//
func (s AdwordsUserListService) Get(selector Selector) (userLists []UserList, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
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
	)
	if err != nil {
		return userLists, err
	}
	getResp := struct {
		Size      int64      `xml:"rval>totalNumEntries"`
		UserLists []UserList `xml:"rval>entries"`
	}{}
	fmt.Printf("%s\n", respBody)
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return userLists, err
	}
	return getResp.UserLists, err
}

// Mutate is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdwordsUserListService#mutate
//
func (s *AdwordsUserListService) Mutate(adwordsUserListOperations UserListOperations) (adwordsUserLists []UserList, err error) {
	return adwordsUserLists, ERROR_NOT_YET_IMPLEMENTED
}
