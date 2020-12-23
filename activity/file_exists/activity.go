package file_exists

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-sleep")

const (
	url     = "url"
)

// SleepActivity is an Activity that can stop flow execution for given time duration.
// inputs : {interval, intervalType}
// outputs: none
type SleepActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &SleepActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *SleepActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *SleepActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	url := context.GetInput(url)
	logger.Info(url)
	context.SetOutput("status",true)
	return true, nil
}