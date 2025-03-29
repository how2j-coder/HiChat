package ecode

import "com/chat/service/pkg/errcode"

var (
	uploadNo       = 5
	uploadName     = "upload"
	uploadBaseCode = errcode.HCode(uploadNo)

	ErrCreateUpload = errcode.NewError(uploadBaseCode+1, "failed to create "+uploadName)
)

