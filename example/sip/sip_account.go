package sip

import (
   "fmt"

   pjsua2 "github/yslim/go_pjproject"
)

type Account struct {
   userId     string
   sipService *SipService
}

func NewAccount(userId string, sipService *SipService) *Account {
   ac := &Account{userId, sipService}
   return ac
}

func (sa *Account) OnRegState(prm pjsua2.OnRegStateParam) {
   sa.sipService.checkThread()

   account := sa.sipService.activeAccounts[sa.userId]
   if account == nil {
      fmt.Printf("[ SipAccount ] OnRegState, account not found(userId=%v)\n", sa.userId)
      return
   }

   info := account.GetInfo()

   var regiState string
   if info.GetRegIsActive() {
      regiState = "REGISTER"
   } else {
      regiState = "UNREGISTER"
   }
   fmt.Printf("[ SipAccount ] %v : code = %v\n", regiState, prm.GetCode())

   sa.sipService.onRegState(info.GetUri(), info.GetRegIsActive(), prm.GetCode())
}

func (sa *Account) OnIncomingCall(prm pjsua2.OnIncomingCallParam) {
   account := sa.sipService.getAccount(sa.userId)
   sipCall := NewCall(sa.sipService)
   call := pjsua2.NewDirectorCall(sipCall, account, prm.GetCallId())
   sipCall.call = call
   callInfo := call.GetInfo()

   fmt.Printf("[ SipAccount ] IncomingCall %v\n"+
        "...remoteUri = %v, localUri = %v\n",
      prm.GetRdata().GetInfo(), callInfo.GetRemoteUri(), callInfo.GetLocalUri())

   sa.sipService.addCall(callInfo.GetCallIdString(), call)

   user := sa.sipService.onIncomingCall(
      call.GetInfo().GetCallIdString(),
      callInfo.GetRemoteUri(), callInfo.GetLocalUri())

   callOpParam := pjsua2.NewCallOpParam()
   if user == nil {
      fmt.Printf("[ SipAccount ] No Available User, Answering PJSIP_SC_BUSY_HERE\n")
      callOpParam.SetStatusCode(pjsua2.PJSIP_SC_BUSY_HERE)
      call.Answer(callOpParam)
   } else {
      callOpParam.SetStatusCode(pjsua2.PJSIP_SC_OK)
      callOpParam.GetOpt().SetAudioCount(1)
      callOpParam.GetOpt().SetVideoCount(0)
      call.Answer(callOpParam)
   }
}
