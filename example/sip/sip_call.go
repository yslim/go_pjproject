package sip

import (
   "fmt"

   pjsua2 "github.com/yslim/go_pjproject"
)

type Call struct {
   sipService *SipService
   call       pjsua2.Call
}

func NewCall(sipService *SipService) *Call {
   return &Call{sipService, nil}
}

func (sc *Call) OnCallState(prm pjsua2.OnCallStateParam) {
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
