package model

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model
	VoucherName        string `json:"voucher_name"`
	VoucherDescription string `json:"voucher_description"`
	Discount           int    `json:"discount"`
}
