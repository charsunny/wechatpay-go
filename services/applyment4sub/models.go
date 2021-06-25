package applyment4sub

const (
	SALES_SCENES_STORE        = "SALES_SCENES_STORE"
	SALES_SCENES_MP           = "SALES_SCENES_MP"
	SALES_SCENES_MINI_PROGRAM = "SALES_SCENES_MINI_PROGRAM"
	SALES_SCENES_WEB          = "SALES_SCENES_WEB"
	SALES_SCENES_APP          = "SALES_SCENES_APP"
	SALES_SCENES_WEWORK       = "SALES_SCENES_WEWORK"
)

type ApplymentRequest struct {
	BusinessCode    string           `json:"business_code"`
	ContactInfo     *ContactInfo     `json:"contact_info"`
	SubjectInfo     *SubjectInfo     `json:"subject_info"`
	BusinessInfo    *BusinessInfo    `json:"business_info"`
	SettlementInfo  *SettlementInfo  `json:"settlement_info"`
	BankAccountInfo *BankAccountInfo `json:"bank_account_info"`
	AdditionInfo    *AdditionInfo    `json:"addition_info,omitempty"`
}
type ContactInfo struct {
	ContactName     string `json:"contact_name"`
	ContactIDNumber string `json:"contact_id_number,omitempty"`
	Openid          string `json:"openid,omitempty"`
	MobilePhone     string `json:"mobile_phone"`
	ContactEmail    string `json:"contact_email"`
}
type BusinessLicenseInfo struct {
	LicenseCopy   string `json:"license_copy"`
	LicenseNumber string `json:"license_number"`
	MerchantName  string `json:"merchant_name"`
	LegalPerson   string `json:"legal_person"`
}
type CertificateInfo struct {
	CertCopy       string `json:"cert_copy"`
	CertType       string `json:"cert_type"`
	CertNumber     string `json:"cert_number"`
	MerchantName   string `json:"merchant_name"`
	CompanyAddress string `json:"company_address"`
	LegalPerson    string `json:"legal_person"`
	PeriodBegin    string `json:"period_begin"`
	PeriodEnd      string `json:"period_end"`
}
type OrganizationInfo struct {
	OrganizationCopy string `json:"organization_copy"`
	OrganizationCode string `json:"organization_code"`
	OrgPeriodBegin   string `json:"org_period_begin"`
	OrgPeriodEnd     string `json:"org_period_end"`
}
type IDCardInfo struct {
	IDCardCopy      string `json:"id_card_copy"`
	IDCardNational  string `json:"id_card_national"`
	IDCardName      string `json:"id_card_name"`
	IDCardNumber    string `json:"id_card_number"`
	CardPeriodBegin string `json:"card_period_begin"`
	CardPeriodEnd   string `json:"card_period_end"`
}
type IDDocInfo struct {
	IDDocCopy      string `json:"id_doc_copy"`
	IDDocName      string `json:"id_doc_name"`
	IDDocNumber    string `json:"id_doc_number"`
	DocPeriodBegin string `json:"doc_period_begin"`
	DocPeriodEnd   string `json:"doc_period_end"`
}
type IdentityInfo struct {
	IDDocType  string      `json:"id_doc_type"`
	IDCardInfo *IDCardInfo `json:"id_card_info,omitempty"`
	IDDocInfo  *IDDocInfo  `json:"id_doc_info,omitempty"`
	Owner      string      `json:"owner"`
}
type UboInfo struct {
	IDType         string `json:"id_type"`
	IDCardCopy     string `json:"id_card_copy"`
	IDCardNational string `json:"id_card_national"`
	IDDocCopy      string `json:"id_doc_copy"`
	Name           string `json:"name"`
	IDNumber       string `json:"id_number"`
	IDPeriodBegin  string `json:"id_period_begin"`
	IDPeriodEnd    string `json:"id_period_end"`
}
type SubjectInfo struct {
	SubjectType           string               `json:"subject_type"`
	BusinessLicenseInfo   *BusinessLicenseInfo `json:"business_license_info,omitempty"`
	CertificateInfo       *CertificateInfo     `json:"certificate_info,omitempty"`
	OrganizationInfo      *OrganizationInfo    `json:"organization_info,omitempty"`
	CertificateLetterCopy string               `json:"certificate_letter_copy,omitempty"`
	IdentityInfo          *IdentityInfo        `json:"identity_info,omitempty"`
	UboInfo               *UboInfo             `json:"ubo_info,omitempty"`
}
type BizStoreInfo struct {
	BizStoreName     string   `json:"biz_store_name"`
	BizAddressCode   string   `json:"biz_address_code"`
	BizStoreAddress  string   `json:"biz_store_address"`
	StoreEntrancePic []string `json:"store_entrance_pic"`
	IndoorPic        []string `json:"indoor_pic"`
	BizSubAppid      string   `json:"biz_sub_appid"`
}
type MpInfo struct {
	MpAppid    string   `json:"mp_appid"`
	MpSubAppid string   `json:"mp_sub_appid"`
	MpPics     []string `json:"mp_pics"`
}
type MiniProgramInfo struct {
	MiniProgramAppid    string   `json:"mini_program_appid"`
	MiniProgramSubAppid string   `json:"mini_program_sub_appid"`
	MiniProgramPics     []string `json:"mini_program_pics"`
}
type AppInfo struct {
	AppAppid    string   `json:"app_appid"`
	AppSubAppid string   `json:"app_sub_appid"`
	AppPics     []string `json:"app_pics"`
}
type WebInfo struct {
	Domain           string `json:"domain"`
	WebAuthorisation string `json:"web_authorisation"`
	WebAppid         string `json:"web_appid"`
}
type WeworkInfo struct {
	CorpID     string   `json:"corp_id"`
	SubCorpID  string   `json:"sub_corp_id"`
	WeworkPics []string `json:"wework_pics"`
}
type SalesInfo struct {
	SalesScenesType []string         `json:"sales_scenes_type"`
	BizStoreInfo    *BizStoreInfo    `json:"biz_store_info,omitempty"`
	MpInfo          *MpInfo          `json:"mp_info,omitempty"`
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"`
	AppInfo         *AppInfo         `json:"app_info,omitempty"`
	WebInfo         *WebInfo         `json:"web_info,omitempty"`
	WeworkInfo      *WeworkInfo      `json:"wework_info,omitempty"`
}
type BusinessInfo struct {
	MerchantShortname string     `json:"merchant_shortname"`
	ServicePhone      string     `json:"service_phone"`
	SalesInfo         *SalesInfo `json:"sales_info"`
}
type SettlementInfo struct {
	SettlementID        string   `json:"settlement_id"`
	QualificationType   string   `json:"qualification_type"` // https://pay.weixin.qq.com/wiki/doc/apiv3_partner/terms_definition/chapter1_1_3.shtml#part-7
	Qualifications      []string `json:"qualifications,omitempty"`
	ActivitiesID        string   `json:"activities_id,omitempty"`
	ActivitiesRate      string   `json:"activities_rate,omitempty"`
	ActivitiesAdditions []string `json:"activities_additions,omitempty"`
}
type BankAccountInfo struct {
	BankAccountType string `json:"bank_account_type"`
	AccountName     string `json:"account_name"`
	AccountBank     string `json:"account_bank"`
	BankAddressCode string `json:"bank_address_code"`
	BankBranchID    string `json:"bank_branch_id,omitempty"`
	BankName        string `json:"bank_name,omitempty"`
	AccountNumber   string `json:"account_number"`
}
type AdditionInfo struct {
	LegalPersonCommitment string   `json:"legal_person_commitment,omitempty"`
	LegalPersonVideo      string   `json:"legal_person_video,omitempty"`
	BusinessAdditionPics  []string `json:"business_addition_pics,omitempty"`
	BusinessAdditionMsg   string   `json:"business_addition_msg,omitempty"`
}

type ApplymentResponse struct {
	ApplymentID string `json:"applyment_id"`
}

type ApplyAuditResponse struct {
	BusinessCode      string         `json:"business_code"`
	ApplymentID       int64          `json:"applyment_id"`
	SubMchid          string         `json:"sub_mchid"`
	SignURL           string         `json:"sign_url"`
	ApplymentState    string         `json:"applyment_state"`
	ApplymentStateMsg string         `json:"applyment_state_msg"`
	AuditDetail       []*AuditDetail `json:"audit_detail,omitempty"`
}
type AuditDetail struct {
	Field        string `json:"field"`
	FieldName    string `json:"field_name"`
	RejectReason string `json:"reject_reason"`
}

type BankAccountInfoResponse struct {
	BankAccountInfo
	VerifyResult     string `json:"verify_result"`
	VerifyFailReason string `json:"verify_fail_reason"`
}
