package server


import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type MyMath struct{
	a int
}


func (mm *MyMath) Add(num1 float64,num2 float64 ) float64 {
	reply := num1+num2
	return reply
}


func (mm *MyMath) Sub(num1 float64,num2 float64 ) float64 {
	reply := num1-num2
	return reply
}

func (mm *MyMath) Fuck(name string, obj *MyMath) string {
	return name + strconv.Itoa(obj.a)
}


func run() {
	m := new(MyMath)
	typ := reflect.TypeOf(m)
	//遍历方法
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		mname := method.Name//方法名字
		fmt.Println("method:"+mname)
		fun := reflect.ValueOf(m).MethodByName(mname)
		ty := method.Type
		args := make([]reflect.Value, ty.NumIn()-1)
		//遍历参数
		for j:= 1; j< ty.NumIn(); j++ {
			//参数类型
			arg_type := ty.In(j).Kind().String()
			fmt.Println("参数" + strconv.Itoa(j) + ":" + arg_type)
			switch arg_type {
			case "float64":
				args[j-1] = reflect.ValueOf(float64(j))
			default:
				fmt.Println("i down'n knowon type " + arg_type)

			}
		}
		ret := fun.Call(args)
		fmt.Println(ret[0])
	}

}

func TestRun(t *testing.T) {
	run()
}