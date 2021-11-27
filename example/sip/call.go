package sip

import (
    "fmt"
    "sync"

    pjsua2 "github.com/yslim/go_pjproject/v2"
)

type Call struct {
    pjsua2.Call
    state pjsua2.Pjsip_inv_state
    mutex sync.Mutex
}

func NewCall() *Call {
    return &Call{}
}

func (ca *Call) OnCallState(prm pjsua2.OnCallStateParam) {
    ca.mutex.Lock()
    defer ca.mutex.Unlock()

    ci := ca.GetInfo()

    if ca.state == ci.GetState() {
        return
    }
    ca.state = ci.GetState()

    callId := ci.GetCallIdString()

    fmt.Printf("OnCallState=%v, RemoteUri=%v, callId=%s\n", ci.GetStateText(), ci.GetRemoteUri(), callId)
}

func (ca *Call) OnCallSdpCreated(prm pjsua2.OnCallSdpCreatedParam) {
    ci := ca.GetInfo()

    fmt.Printf("OnCallSdpCreated, SendingSdp=%v\n", prm.GetSdp().GetWholeSdp())

    sipCallId := ci.GetCallIdString()
    sipRole := ci.GetRole()

    fmt.Printf("OnCallSdpCreated, sipCallId=%v, role=%v, state=%v\n",
        sipCallId, sipRole, ci.GetState())

}

func (ca *Call) OnInstantMessage(prm pjsua2.OnInstantMessageParam) {
    fmt.Printf("OnInstantMessage, From: %s, To: %s, Message: %s\n",
        prm.GetFromUri(), prm.GetToUri(), prm.GetMsgBody())
}

func (ca *Call) OnInstantMessageStatus(prm pjsua2.OnInstantMessageStatusParam) {
    fmt.Printf("OnInstantMessageStatus\n")
}

func (ca *Call) OnStreamCreated(prm pjsua2.OnStreamCreatedParam) {
    fmt.Printf("OnStreamCreated\n")
}

func (ca *Call) OnStreamDestroyed(prm pjsua2.OnStreamDestroyedParam) {
    fmt.Printf("OnStreamDestroyed\n")
}

func (ca *Call) OnDtmfDigit(prm pjsua2.OnDtmfDigitParam) {
    fmt.Printf("OnDtmfDigit\n")
}

func (ca *Call) OnCallTransferRequest(prm pjsua2.OnCallTransferRequestParam) {
    fmt.Printf("OnCallTransferRequest\n")
}

func (ca *Call) OnCallTransferStatus(prm pjsua2.OnCallTransferStatusParam) {
    fmt.Printf("OnCallTransferStatus\n")
}

func (ca *Call) OnCallReplaceRequest(prm pjsua2.OnCallReplaceRequestParam) {
    fmt.Printf("OnCallReplaceRequest\n")
}

func (ca *Call) OnCallReplaced(prm pjsua2.OnCallReplacedParam) {
    fmt.Printf("OnCallReplaced\n")
}

func (ca *Call) OnCallMediaTransportState(prm pjsua2.OnCallMediaTransportStateParam) {
    fmt.Printf("OnCallMediaTransportState\n")
}

func (ca *Call) OnCreateMediaTransport(prm pjsua2.OnCreateMediaTransportParam) {
    fmt.Printf("OnCreateMediaTransport\n")
}

func (ca *Call) OnCreateMediaTransportSrtp(prm pjsua2.OnCreateMediaTransportSrtpParam) {
    fmt.Printf("OnCreateMediaTransportSrtp\n")
}

func (ca *Call) OnCallMediaEvent(prm pjsua2.OnCallMediaEventParam) {
    fmt.Printf("OnCallMediaEvent\n")
}

func (ca *Call) OnCallTsxState(prm pjsua2.OnCallTsxStateParam) {
    fmt.Printf("OnCallTsxState\n")
}

func (ca *Call) OnCallMediaState(prm pjsua2.OnCallMediaStateParam) {
    fmt.Printf("OnCallMediaState\n")
}

func (ca *Call) OnCallRxReinvite(prm pjsua2.OnCallRxReinviteParam) {
    fmt.Printf("OnCallRxReinvite\n")
}

func (ca *Call) OnCallRxOffer(prm pjsua2.OnCallRxOfferParam) {
    fmt.Printf("OnCallRxOffer\n")
}

func (ca *Call) OnCallTxOffer(prm pjsua2.OnCallTxOfferParam) {
    fmt.Printf("OnCallTxOffer\n")
}

func (ca *Call) OnTypingIndication(prm pjsua2.OnTypingIndicationParam) {
    fmt.Printf("OnTypingIndication\n")
}

func (ca *Call) OnCallRedirected(prm pjsua2.OnCallRedirectedParam) {
    fmt.Printf("OnCallRedirected\n")
}
