// Package downloadfile implements a file download for Flogo
package filedownload2

// Imports
import (
	"time"

	"github.com/nareix/curl"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// Constants
const (
	ivURL         = "url"
	ovResult      = "result"
	ivtimes       = "ivtimes"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the action
	url := context.GetInput(ivURL).(string)
	ivtimes := context.GetInput(ivtimes).(int)
	// Create new request
	req := curl.Get(url)

	// Set timeouts
	// DialTimeout is the TCP Connection Timeout
	// Timeout is the Download Timeout
	req.DialTimeout(time.Second * 10)
	req.Timeout(time.Second * 30)

	// Specify a progress monitor, otherwise it doesn't work
	req.Progress(func(p curl.ProgressStatus) {}, time.Second)

	// Execute the request and return the result
	for k:= 0;k<ivtimes ; k++ {
		res, err := req.Do()
		if err != nil {
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		if res.StatusCode == 200 {
			context.SetOutput(ovResult, "OK")
			return true, nil
		}
		time.Sleep(10 * time.Second)
	}


	context.SetOutput(ovResult, "ERROR")
	return true, nil
}


