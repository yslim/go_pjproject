package sip

import (
    "fmt"

    pjsua2 "github.com/yslim/go_pjproject"
)

type Account struct {
    pjsua2.Account
}

func NewAccount() *Account {
    return &Account{}
}

func (ac *Account) OnRegState(prm pjsua2.OnRegStateParam) {
    checkThread()

    info := ac.GetInfo()

    var regiState string

    if info.GetRegIsActive() {
        regiState = "REGISTER"
    } else {
        regiState = "UNREGISTER"
    }

    fmt.Printf("OnRegState, regiState=%v, code=%v\n", regiState, prm.GetCode())

    onRegState(info.GetUri(), info.GetRegIsActive(), prm.GetCode())
}

func (ac *Account) OnIncomingCall(prm pjsua2.OnIncomingCallParam) {
    call := NewCall()
    call.Call = pjsua2.NewDirectorCall(call, ac, prm.GetCallId())

    callInfo := call.GetInfo()
    from := callInfo.GetRemoteUri()
    to := callInfo.GetLocalUri()
    callId := callInfo.GetCallIdString()

    fmt.Printf("OnIncomingCall, from=%v, to=%v, callId=%s\n", from, to, callId)

    onIncomingCall(callId, from, to)
}

func (ac *Account) OnInstantMessage(prm pjsua2.OnInstantMessageParam) {
}

func (ac *Account) OnInstantMessageStatus(prm pjsua2.OnInstantMessageStatusParam) {
}

func (ac *Account) OnRegStarted(prm pjsua2.OnRegStartedParam) {
}

func (ac *Account) OnIncomingSubscribe(prm pjsua2.OnIncomingSubscribeParam) {
}

func (ac *Account) OnTypingIndication(prm pjsua2.OnTypingIndicationParam) {
}

func (ac *Account) OnMwiInfo(prm pjsua2.OnMwiInfoParam) {
}
