package main

import (
   "fmt"
   "os"
   "os/signal"

   pjsua2 "github/yslim/go_pjproject"
   "github/yslim/go_pjproject/example/sip"
)

func main() {

   sipUser := SipUser{}
   sipService := sip.NewSipService(&sipUser)
   sipUser.sipService = sipService

   sipService.RegisterAccount("test1", "test1")

   c := make(chan os.Signal, 1)
   signal.Notify(c, os.Interrupt)

   <- c
}

type SipUser struct {
   sipService *sip.SipService
   callId     string
}

func (su *SipUser) OnRegState(userId string, isActive bool, code pjsua2.Pjsip_status_code) {
   fmt.Printf("[ OnRegState ] userId=%v, isActive=%v, code=%v\n", userId, isActive, code)
   if isActive {
      su.callId = su.sipService.MakeCall("test1", "test1")
   }
}
func (su *SipUser) OnIncomingCall(callIdString string, from string, to string) interface{} {
   su.callId = callIdString
   return "user"
}
