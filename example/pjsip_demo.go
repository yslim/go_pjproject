package main

import (
	"fmt"
	"github.com/yslim/pjproject"
	"sync"
	"time"
)

var (
	mutex     sync.Mutex
	endpoint  = pjsua2.NewEndpoint()
	logWriter = pjsua2.NewDirectorLogWriter(new(SipLogWriter))
)

func checkThread() {
	mutex.Lock()
	defer mutex.Unlock()

	if !endpoint.LibIsThreadRegistered() {
		endpoint.LibRegisterThread("")
	}
}

func main() {
	// Create endpoint
	epConfig := pjsua2.NewEpConfig()
	epConfig.GetLogConfig().SetLevel(4)
	epConfig.GetLogConfig().SetWriter(logWriter)
	endpoint.LibCreate()

	// Init library
	endpoint.LibInit(epConfig)
	endpoint.AudDevManager().SetNullDev()

	// Transport
	transportConfig := pjsua2.NewTransportConfig()
	transportConfig.SetPort(5060)
	endpoint.TransportCreate(pjsua2.PJSIP_TRANSPORT_UDP, transportConfig)

	// Start library
	endpoint.LibStart()

	fmt.Printf("[ SipService ] Available codecs:\n")
	for i := 0; i < int(endpoint.CodecEnum().Size()); i++ {
		c := endpoint.CodecEnum().Get(i)
		fmt.Printf("\t - %s (priority: %d)\n", c.GetCodecId(), c.GetPriority())
	}

	fmt.Printf("[ SipService ] PJSUA2 STARTED ***\n")

	// Add account
	accountConfig := pjsua2.NewAccountConfig()
	accountConfig.SetIdUri("sip:test1@pjsip.org")
	accountConfig.GetRegConfig().SetRegistrarUri("sip:sip.pjsip.org")
	cred := pjsua2.NewAuthCredInfo("digest", "*", "test1", 0, "test1")
	accountConfig.GetSipConfig().GetAuthCreds().Add(cred)

	sipAccount := NewMyAccount()
	sipAccount.Create(accountConfig)

	time.Sleep(2 * time.Second)
}
