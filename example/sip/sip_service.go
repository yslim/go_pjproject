package sip

import (
    "fmt"
    "strings"
    "sync"
    "time"

    pjsua2 "github.com/yslim/go_pjproject"
)

var (
    endpoint pjsua2.Endpoint
    accounts map[string]*Account
    calls    map[string]*Call
    handlers = make(map[IUserService]bool)
    mutex    sync.Mutex

    account *Account
)

func init() {
    initEndPoint()
}

func initEndPoint() {
    endpoint = pjsua2.NewEndpoint()
    accounts = make(map[string]*Account)
    calls = make(map[string]*Call)

    // Create endpoint
    endpoint.LibCreate()

    // Initialize endpoint
    epConfig := pjsua2.NewEpConfig()
    epConfig.GetUaConfig().SetUserAgent(config.AgentName)
    epConfig.GetUaConfig().SetMaxCalls(config.MaxCall)

    epConfig.GetLogConfig().SetLevel(config.PjLogLevel)
    if config.PjLogLevel > 0 {
        epConfig.GetLogConfig().SetWriter(pjsua2.NewDirectorLogWriter(new(LogWriter)))
    }

    endpoint.LibInit(epConfig)
    endpoint.AudDevManager().SetNullDev()

    transportConfig := pjsua2.NewTransportConfig()
    transportConfig.SetPort(config.LocalPort)

    var transport pjsua2.Pjsip_transport_type_e
    if strings.EqualFold(config.Transport, "UDP") {
        transport = pjsua2.PJSIP_TRANSPORT_UDP
    } else if strings.EqualFold(config.Transport, "TLS") {
        transport = pjsua2.PJSIP_TRANSPORT_TLS
    } else if strings.EqualFold(config.Transport, "TCP") {
        transport = pjsua2.PJSIP_TRANSPORT_TCP
    } else {
        fmt.Printf("unknown config.Transport = %s\n", config.Transport)
    }

    endpoint.TransportCreate(transport, transportConfig)
    endpoint.LibStart()

    fmt.Printf("[ sip.Service ] Available codecs:\n")
    for i := 0; i < int(endpoint.CodecEnum2().Size()); i++ {
        c := endpoint.CodecEnum2().Get(i)
        fmt.Printf("\t - %s (priority: %d)\n", c.GetCodecId(), c.GetPriority())
        if !strings.HasPrefix(c.GetCodecId(), "PCM") {
            endpoint.CodecSetPriority(c.GetCodecId(), 0)
        }
    }

    time.AfterFunc(1*time.Second, func() {
        checkThread()
        emitEvent("OnSipReady")
    })

    fmt.Printf("[ sip.Service ] PJSUA2 STARTED ***\n")
}

func createLocalAccount(uid string, password string) *Account {
    // create test account
    accountConfig := pjsua2.NewAccountConfig()
    accountConfig.SetIdUri("sip:test1@pjsip.org")
    accountConfig.GetRegConfig().SetRegistrarUri("sip:sip.pjsip.org")
    cred := pjsua2.NewAuthCredInfo("digest", "*", "test1", 0, "test1")
    accountConfig.GetSipConfig().GetAuthCreds().Add(cred)

    account := NewAccount()

    pjAccount := pjsua2.NewDirectorAccount(account)
    pjAccount.Create(accountConfig)

    account.Account = pjAccount

    fmt.Printf("createLocalAccount: Account Created, URI=%s\n", account.GetInfo().GetUri())

    return account
}

func getRemoteURI(remoteUser string) string {
    remoteUri := strings.Builder{}

    remoteUri.WriteString("sip:")
    remoteUri.WriteString(remoteUser)
    remoteUri.WriteString("@")
    remoteUri.WriteString(config.ProxyAddress)
    remoteUri.WriteString(":")
    remoteUri.WriteString(fmt.Sprintf("%d", config.ProxyPort))
    if !strings.EqualFold(config.Transport, "UDP") {
        if strings.EqualFold(config.Transport, "TCP") {
            remoteUri.WriteString(";transport=tcp")
        } else {
            remoteUri.WriteString(";transport=tls")
        }
    }
    return remoteUri.String()
}

func emitEvent(event string, params ...interface{}) {
    // emit event at new thread
    go func() {
        for h := range handlers {
            switch event {
            case "OnSipReady":
                h.OnSipReady()
            case "OnRegState":
                h.OnRegState(params[0].(string), params[1].(bool), params[2].(int))
            case "OnIncomingCall":
                h.OnIncomingCall(params[0].(string), params[1].(string), params[2].(string))
            default:
                fmt.Printf("emitEvent, unknown event = %s\n", event)
            }
        }
    }()
}

func onRegState(uri string, active bool, code pjsua2.Pjsip_status_code) {
    emitEvent("OnRegState", uri, active, int(code))
}

func onIncomingCall(callId string, from string, to string) {
    emitEvent("OnIncomingCall", callId, from, to)
}

func checkThread() {
    mutex.Lock()
    defer mutex.Unlock()

    if !endpoint.LibIsThreadRegistered() {
        endpoint.LibRegisterThread(config.AgentName)
    }
}

// public functions

func RegisterEventHandler(sipUser IUserService) {
    handlers[sipUser] = true
}

func RegisterAccount(uid string, password string) {
    checkThread()

    fmt.Printf("RegisterAccount, uid=%v\n", uid)

    account = createLocalAccount(uid, password)
}

func MakeCall(fromUid string, toUid string) string {
    checkThread()

    remoteUri := getRemoteURI(toUid)

    // make outgoing call
    call := NewCall()
    call.Call = pjsua2.NewDirectorCall(call, account)

    callOpParam := pjsua2.NewCallOpParam(true)
    callOpParam.GetOpt().SetAudioCount(1)

    fmt.Printf("MakeCall, From=%s, To=%s\n", account.GetInfo().GetUri(), remoteUri)

    call.MakeCall(remoteUri, callOpParam)

    ci := call.GetInfo()
    callId := ci.GetCallIdString()

    return callId
}
