package ecode

import "com/chat/service/pkg/errcode"

var (
	roleNo = 3
	roleName = "role"
	roleBaseCode = errcode.HCode(roleNo)

	ErrCreateRole = errcode.NewError(roleBaseCode + 1, "failed to create " + roleName)
)