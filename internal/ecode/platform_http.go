package ecode

import "com/chat/service/pkg/errcode"

var (
	platformNo       = 4
	platformName     = "platform"
	platformBaseCode = errcode.HCode(platformNo)

	ErrCreatePlatform = errcode.NewError(platformBaseCode+1, "failed to create "+platformName)
)
