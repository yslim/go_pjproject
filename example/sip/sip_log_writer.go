package sip

import (
   "fmt"
   "strings"

   "github.com/yslim/go_pjproject"
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

   fmt.Printf("[ SIP ] %v\n", msg)
}
