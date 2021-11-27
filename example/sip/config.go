package sip

type tSipProperties struct {
    AgentName string `properties:"agent-name"`

    LocalPort    uint   `properties:"local-port"`

    ProxyAddress string `properties:"proxy-address"`
    ProxyPort    uint   `properties:"proxy-port"`

    Transport string `properties:"transport"`
    UseSrtp   bool   `properties:"use-srtp"`

    PjLogLevel uint `properties:"pj-log-level"`

    MaxCall        uint `properties:"max-call"`
    MediaPortStart uint `properties:"media-port-start"`
}

var config = tSipProperties{
    AgentName: "pjsip",
    LocalPort: 5060,
    ProxyAddress: "pjsip.org",
    ProxyPort: 5060,
    Transport: "UDP",
    UseSrtp: false,
    PjLogLevel: 4,
    MaxCall: 1,
    MediaPortStart: 4000,
}