package gocrypto

import "testing"

func TestComparePasswords(t *testing.T) {
	pwd := "how2j.online"

	hashStr, err := HashAndSaltPassword(pwd)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(hashStr)

	ok := VerifyPassword(pwd, hashStr)
	if !ok {
		t.Error("passwords mismatch")
		return
	}
	t.Log("passwords match")
}

