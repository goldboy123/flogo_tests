package file_exists

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{})
}

type Output struct {
	Status bool `md:"file exists"` // the file is exists
}

func (o Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"status" : o.Status,
	}
}

func (o Output) FromMap(values map[string]interface{}) error {
	var err error
	o.Status,err = coerce.ToBool(values["status"])
	if err != nil{
		return err
	}
	return nil
}

type Input struct {
	Url string `md:"file url"` // the url to request
}


func (i Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"url" : i.Url,
	}
}



func (i Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Url, _ = coerce.ToString(values["url"])
	if err != nil {
		return err
	}
	return nil
}



var activityMd = activity.ToMetadata(&Input{})

// Activity is an Activity that is used to log a message to the console
// inputs : none
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	input := &Input{}
	ctx.GetInputObject(input)

	url := input.Url
	ctx.Logger().Debugf(url)
	//output := &Output{Status:true}
	err = ctx.SetOutput("status",true)
	if err != nil{
		return false , err
	}
	return true, nil
}
