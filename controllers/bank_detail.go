package controllers

import (
	"log"
	"time"

	dbMod "touristapp.com/db"
	Models "touristapp.com/models"
)

//AddBankDetail inserts a new bank info fo a supplier to the database
func AddBankDetail(newBankDetail Models.NewBankDetail) (int64, error) {
	queryNewbankDetail := `
		INSERT INTO bank_details(beneficiary_name, beneficiary_account_number, beneficiary_address,
			beneficiary_bank, beneficiary_branch_name, beneficiary_bank_address, beneficiary_ifsc_code,
			beneficiary_swift_code, beneficiary_ibn_code, beneficiary_bic_code, is_verified,
			created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $12)
		RETURNING id
	`
	var bankDetailID int64

	if err := dbMod.DB.QueryRow(
		queryNewbankDetail,
		newBankDetail.BeneficiaryName,
		dbMod.NewNullString(newBankDetail.BeneficiaryAccountNumber),
		dbMod.NewNullString(newBankDetail.BeneficiaryAddress),
		dbMod.NewNullString(newBankDetail.BeneficiaryBank),
		dbMod.NewNullString(newBankDetail.BeneficiaryBranchName),
		dbMod.NewNullString(newBankDetail.BeneficiaryBankAddress),
		dbMod.NewNullString(newBankDetail.BeneficiaryIFSCCode),
		dbMod.NewNullString(newBankDetail.BeneficiarySwiftCode),
		dbMod.NewNullString(newBankDetail.BeneficiaryIBNCode),
		dbMod.NewNullString(newBankDetail.BeneficiaryBICCode),
		newBankDetail.IsVerified,
		time.Now(),
	).Scan(&bankDetailID); err != nil {
		dbMod.Rollback()
		log.Printf("Error while inserting or scanning the bankdetail: %s", err)
		return 0, err
	}
	return bankDetailID, nil
}
