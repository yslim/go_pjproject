package main

import (
   "fmt"
   "pjproject"
)

type MyCall struct {
   pjsua2.Call
   account *MyAccount
}

func NewMyCall(acc *MyAccount, callId int) *MyCall {
   return &MyCall{pjsua2.NewCall(acc, callId), acc}
}

func (c *MyCall) OnCallState(prm pjsua2.OnCallStateParam) {
   ci := c.GetInfo()

   fmt.Printf("[ SipCall ] onCallState %v, aor = %v", ci.GetStateText(), ci.GetRemoteUri())

   if ci.GetState() == pjsua2.PJSIP_INV_STATE_DISCONNECTED {
      fmt.Printf("[ SipCall ] Call Closed, CallId=%v, AOR=%v, reason=%v, lastStatusCode=%v",
         ci.GetCallIdString(), ci.GetRemoteUri(),
         ci.GetLastReason(), ci.GetLastStatusCode())

      c.account.removeCall(c)

      // Delete the call
      pjsua2.DeleteCall(c.Call)
   }
}
