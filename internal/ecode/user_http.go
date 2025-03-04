package ecode

import "com/chat/service/pkg/errcode"

var (
	userNO       = 1
	userName     = "user"
	userBaseCode = errcode.HCode(userNO)

	ErrCreateUser     = errcode.NewError(userBaseCode+1, "failed to create "+userName)
	ErrDeleteByIDUser = errcode.NewError(userBaseCode+2, "failed to delete "+userName)
	ErrUpdateByIDUser = errcode.NewError(userBaseCode+3, "failed to update "+userName)
	ErrGetByIDUser    = errcode.NewError(userBaseCode+4, "failed to get "+userName+" details")
	ErrListUser       = errcode.NewError(userBaseCode+5, "failed to list of "+userName)

	// error codes are globally unique, adding 1 to the previous error code
)
