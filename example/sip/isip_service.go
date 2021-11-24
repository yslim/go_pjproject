package sip

import (
	pjsua2 "github.com/yslim/go_pjproject"
)

type ISipService interface {
	OnRegState(userId string, isActive bool, code pjsua2.Pjsip_status_code)
	OnIncomingCall(callIdString string, from string, to string) interface{}
}
