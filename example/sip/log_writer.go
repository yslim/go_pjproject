package sip

import (
    "fmt"
    "strings"

    pjsua2 "github.com/yslim/go_pjproject/v2"
)

type LogWriter struct {
    name string
}

func (l *LogWriter) Write(entry pjsua2.LogEntry) {
    msg := entry.GetMsg()
    strings.Replace(msg, "\r", "", -1)

    if msg[len(msg)-1] == '\n' {
        msg = msg[37 : len(msg)-1]
    }

    fmt.Printf("[ SIP ] %s\n", msg)
}
