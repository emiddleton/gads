package gads

type ReportDefinitionService struct {
	Auth
}

func NewReportDefinitionService(auth *Auth) *ReportDefinitionService {
	return &ReportDefinitionService{Auth: *auth}
}
