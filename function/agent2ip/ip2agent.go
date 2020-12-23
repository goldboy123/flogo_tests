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


type fnIp2Agent struct {

}

func (f fnIp2Agent) Name() string {
	return "ip2agent"
}

func (f fnIp2Agent) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (f fnIp2Agent) Eval(params ...interface{}) (interface{}, error) {

	if (len(params) > 2){
		return nil,fmt.Errorf("args too much")
	}
	a, err1 := coerce.ToString(params[0])
	if err1 != nil{
		return false,err1
	}
	d := strings.Split(a,".")
	var  agentid float64

	for i:=0;i<4;i++ {
		tmp, _:=strconv.Atoi(d[i])
		agentid += math.Pow(256, float64(i)) * float64(tmp)
	}

	ret := strconv.FormatInt(int64(agentid),10)
	return ret,nil
}

func init(){
	function.Register(&fnIp2Agent{})
}