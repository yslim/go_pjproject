package sip

import (
   "fmt"
   pjsua2 "pjproject"
)

type SipCall struct {
   sipService *SipService
   call       pjsua2.Call
}

func NewSipCall(sipService *SipService) *SipCall {
   return &SipCall{sipService, nil}
}

func (sc *SipCall) OnCallState(prm pjsua2.OnCallStateParam) {
   ci := sc.call.GetInfo()

   fmt.Printf("[ SipCall ] onCallState %v, aor = %v", ci.GetStateText(), ci.GetRemoteUri())

   if ci.GetState() == pjsua2.PJSIP_INV_STATE_DISCONNECTED {
      fmt.Printf("[ SipCall ] Call Closed, CallId=%v, AOR=%v, reason=%v, lastStatusCode=%v",
         ci.GetCallIdString(), ci.GetRemoteUri(),
         ci.GetLastReason(), ci.GetLastStatusCode())

      sc.sipService.removeCall(ci.GetCallIdString())

      // Delete the call
      pjsua2.DeleteCall(sc.call)
   }
}
