package main

import (
	"fmt"
	"github.com/yslim/pjproject"
)

type MyAccount struct {
	pjsua2.Account
	calls map[pjsua2.Call]int
}

func NewMyAccount() *MyAccount {
	return &MyAccount{pjsua2.NewAccount(), make(map[pjsua2.Call]int)}
}

func (acc *MyAccount) removeCall(call pjsua2.Call) {
	delete(acc.calls, call)
}

func (acc *MyAccount) OnRegState(prm pjsua2.OnRegStateParam) {
	checkThread()

	info := acc.GetInfo()

	var regiState string
	if info.GetRegIsActive() {
		regiState = "REGISTER"
	} else {
		regiState = "UNREGISTER"
	}

	fmt.Printf("[MyAccount] regiState=%v, code=%v\n", regiState, prm.GetCode())
}

func (acc *MyAccount) OnIncomingCall(prm pjsua2.OnIncomingCallParam) {
	call := NewMyCall(acc, prm.GetCallId())
	callInfo := call.GetInfo()
	from := callInfo.GetRemoteUri()
	to := callInfo.GetLocalUri()

	fmt.Printf("[ SipAccount ] IncomingCall %v\n...remoteUri = %v, localUri = %v",
		prm.GetRdata().GetInfo(), from, to)
	fmt.Printf("...WholeMsg = {}", prm.GetRdata().GetWholeMsg())

	acc.calls[call] = 1

	callOpParam := pjsua2.NewCallOpParam()
	callOpParam.SetStatusCode(pjsua2.PJSIP_SC_OK)
	callOpParam.GetOpt().SetAudioCount(1)
	callOpParam.GetOpt().SetVideoCount(0)
	call.Answer(callOpParam)
}
