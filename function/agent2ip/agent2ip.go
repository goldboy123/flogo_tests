package agent2ip

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"math"
	"strconv"
	"strings"
)


type fnAgent2Ip struct {

}

func (f fnAgent2Ip) Name() string {
	return "agent_to_ip"
}

func (f fnAgent2Ip) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (f fnAgent2Ip) Eval(params ...interface{}) (interface{}, error) {

	if (len(params) > 2){
		return nil,fmt.Errorf("args too much")
	}
	a, err1 := coerce.ToFloat64(params[0])
	if err1 != nil{
		return false,err1
	}

	var ip string
	for i :=0;i < 4 ;i++  {
		tmp := strconv.FormatInt(int64(a / math.Pow(256, float64(i)))%256,10)
		ip += tmp + "."
	}
	ip = strings.TrimRight(ip,".")

	return ip,nil
}

func init(){
	function.Register(&fnAgent2Ip{})
}