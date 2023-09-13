package eventlog

import (
	"fmt"
	"strconv"
	"time"
)

var analysisEventsLogFmt = "%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%d\n"

type AnalysisLogger struct {
	RunID         string
	RunSerial     string
	CheckSequence string
	Repository    string
	Shortcode     string
	CommitSHA     string
	IsFullRun     bool
	IsIDERun      bool
}

func (a *AnalysisLogger) Log(runType, stage string) {
	fmt.Printf(analysisEventsLogFmt,
		runType,
		a.RunID,
		a.RunSerial,
		a.CheckSequence,
		a.Shortcode,
		a.Repository,
		a.CommitSHA,
		strconv.FormatBool(a.IsFullRun),
		strconv.FormatBool(a.IsIDERun),
		stage,
		time.Now().Unix(),
	)
}
