package main

import (
   "fmt"
   "pjproject"
   "strings"
)

type SipLogWriter struct {
   name string
}

func (l *SipLogWriter) Write(entry pjsua2.LogEntry) {
   msg := entry.GetMsg()
   strings.Replace(msg, "\r", "", -1)

   if msg[len(msg)-1] == '\n' {
      msg = msg[37 : len(msg)-1]
   }

   fmt.Printf("[ SIP ] %v\n", msg)
}
