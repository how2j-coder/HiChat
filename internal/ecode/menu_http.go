package ecode

import "com/chat/service/pkg/errcode"

var (
	menuNo       = 2
	menuName     = "menu"
	menuBaseCode = errcode.HCode(menuNo)

	ErrCreateMenu = errcode.NewError(menuBaseCode+1, "failed to create "+menuName)
)
