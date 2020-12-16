package add

import (
	"flogo/core/data/coerce"
	"flogo/core/data/expression/function"
	"fmt"
	"github.com/project-flogo/core/data"
)


type fnAddTest struct {

}

func (f fnAddTest) Name() string {
	return "add test"
}

func (f fnAddTest) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (f fnAddTest) Eval(params ...interface{}) (interface{}, error) {
	if (len(params) > 2){
		return nil,fmt.Errorf("args too much")
	}
	a, err1 := coerce.ToInt(params[0])
	if nil != err1{
		return nil,fmt.Errorf(err1.Error())
	}
	b, err2 := coerce.ToInt(params[1])
	if nil != err2{
		return nil,fmt.Errorf(err2.Error())
	}
	return a+b,nil
}

func init(){
	function.Register(&fnAddTest{})
}