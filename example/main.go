package main

import (
    "fmt"
    "os"
    "os/signal"

    "example/sip"
)

func main() {

    sipUser := SipUser{}

    sip.RegisterEventHandler(&sipUser)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)

    <-c
}

type SipUser struct {
    callId string
    state  string
}

func (su *SipUser) OnSipReady() {
    fmt.Printf("OnSipReady\n")

    su.state = "OnSipReady"

    sip.RegisterAccount("test1", "test1")
}

func (su *SipUser) OnRegState(uid string, isActive bool, code int) {
    fmt.Printf("OnRegState, userId = %s, isActive = %v, code = %v\n", uid, isActive, code)

    if isActive {
        su.state = "OnRegState"
        sip.MakeCall("test1", "test2")
    } else {
        su.state = "OnSipReady"
    }
}

func (su *SipUser) OnIncomingCall(callId string, from string, to string) {
    su.callId = callId
    // sip.Answer(su.callId, sip.OK)
}
