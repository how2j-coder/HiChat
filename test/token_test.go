package test

import (
	"HiChat/common"
	"fmt"
	"testing"
)

func TestCrateToken(t *testing.T) {

	type Us struct {
		Name string `json:"name"`
	}

	s := Us{Name: "12312"}
	fmt.Println(common.GenerateTaken(s, "how2j"))
}

func TestParseToken(t *testing.T) {
	s := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbmZvX3QiOnsibmFtZSI6IjEyMzEyIn19.TePAG_qRo4uK0-UL4Nnub-Joj4-AnSGNgaAuweIECyG224WqmA6HvuQJiU4YjsTCI6oGRakIUn9vhNm6d0uXhDcWSLY7D7hNfppy4iGIbWQAmNe8GwGwAW9LqOI3QEEPZ9rQy0X7fXLUT66s01Ix2pK5Zwm2QKBqLD-a0nvTMP4UyDHlhBkrEuInQbp8npBhY1qDZiHBfVt2BkGXv8Vxs91dzr_3M3FBtXfuBDBUm7NwLQfIA7IceGoqInvrHrHBvTBo1fDZ_dNzDYCu6Rk37IuTxk_FehogWkr1Lj6lZZevuP8frHRjsdGnA69TEmjnb-Uq4dRhr__1C2I4CMlByw"
	fmt.Println(common.DecryptTaken(s))
}
