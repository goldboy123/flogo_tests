package file_exists

import (
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&MyActivity{})
}


// Activity is an Activity that is used to log a message to the console
// inputs : url
// outputs: status
type MyActivity struct {
	metadata *activity.Metadata
}


func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval
func (a *MyActivity) Eval(ctx activity.Context) (done bool, err error) {

	//url := ctx.GetInput("url").(string)
	//ctx.Logger().Debugf(url)
	////output := &Output{Status:true}
	//err = ctx.SetOutput("status",true)
	//if err != nil{
	//	return false , err
	//}
	return true, nil
}
