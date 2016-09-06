package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
	"log/syslog"
	"os"
)
// "github.com/hashicorp/serf/serf"

var (
	debug        = flag.Bool("debug", false, "Enable debugging")
	printVersion = flag.Bool("version", false, "Print the version number")
	logToSyslog  = flag.Bool("syslog", false, "Log to syslog")
)

func main() {
	flag.Parse()
	if *printVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}
	if *debug {
		log.SetLevel(log.DebugLevel)
	}
	if *logToSyslog {
		hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
		if err == nil {
			log.AddHook(hook)
		}
	}
	os.Exit(0)
}
