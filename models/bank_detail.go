package models

//NewBankDetail is the model for inserting a new row
//of bank details for a supplier
type NewBankDetail struct {
	BeneficiaryName          string `json:"beneficiary_name" form:"beneficiary_name"`
	BeneficiaryAccountNumber string `json:"beneficiary_account_number" form:"beneficiary_account_number"`
	BeneficiaryAddress       string `json:"beneficiary_address" form:"beneficiary_address"`
	BeneficiaryBank          string `json:"beneficiary_bank" form:"beneficiary_bank"`
	BeneficiaryBranchName    string `json:"beneficiary_branch_name" form:"beneficiary_branch_name"`
	BeneficiaryBankAddress   string `json:"beneficiary_bank_address" form:"beneficiary_bank_address"`
	BeneficiaryIFSCCode      string `json:"beneficiary_ifsc_code" form:"beneficiary_ifsc_code"`
	BeneficiarySwiftCode     string `json:"beneficiary_swift_code" form:"beneficiary_swift_code"`
	BeneficiaryIBNCode       string `json:"beneficiary_ibn_code" form:"beneficiary_ibn_code"`
	BeneficiaryBICCode       string `json:"beneficiary_bic_code" form:"beneficiary_bic_code"`
	IsVerified               bool   `json:"is_verified" form:"is_verified"`
}
