package addition

import (
	"fmt"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}               //creating a pointer of the input ref
	err = ctx.GetInputObject(input) //GetInputObject gets all the activity input as the specified object.
	if err != nil {
		return true, err
	}
	fmt.Println(input.Num1)

	out := input.Num1 + input.Num2

	output := &Output{Output: out}
	err = ctx.SetOutputObject(output) //SetOutputObject sets the activity output as the specified object.

	if err != nil {
		return true, err
	}

	return true, nil
}
