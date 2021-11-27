package sip

// IUserService : User callback functions
type IUserService interface {
	OnSipReady()
	OnRegState(uid string, isActive bool, code int)
	OnIncomingCall(callId string, from string, to string)
	// OnCallConnected(callId string)
	// OnCallClosed(callId string, lastReason string)
	// OnInstantMessage(from string, to string, msg string)
	// OnInstantMessageStatus(from string, to string, code int, reason string, msgBody string)
}
