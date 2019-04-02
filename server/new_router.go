package server

import (
	"reflect"
	"testing"
)

func getMethodParams(handler *Op) {
	handlerType := reflect.TypeOf(handler)

	// 遍历方法
	for i := 0; i < handlerType.NumMethod(); i++ {
		method := handlerType.Method(i)
		methodName := method.Name // method name

		methodObj := reflect.ValueOf(handler).MethodByName(methodName) // method obj

		methodType := method.Type

		args := make([]reflect.Value, methodType.NumIn()-1) // method params array

		// 遍历参数
		for j := 1; j < methodType.NumIn(); j++ {
			//argType := methodType.In(j).Kind().String()

		}

		methodObj.Call(args)
	}

}

type Invoker struct {
	op     *Op
	url    string
	method *Method
}

type Method struct {
	methodType  reflect.Type
	methodValue reflect.Value

	argsTypes  []reflect.Type
	argsValues []reflect.Value

	resultTypes  []reflect.Type
	resultValues []reflect.Value

	name string
	url  string
}

func (v *Invoker) invoke(args []interface{}) []reflect.Value {
	in := make([]reflect.Value, v.method.methodType.NumIn() - 1)

	for i := 0; i< len(in); i++ {
		in[i] = reflect.ValueOf(args[i])
	}

	result := v.method.methodValue.Call(in)
	return result
}

func Build(url string, handler *Op, methodName string) *Invoker {
	invoker := &Invoker{}
	invoker.url = url
	invoker.op = handler
	invoker.method = buildMethod(handler, methodName)
	return invoker
}

func buildMethod(handler *Op, methodName string) *Method {
	_method := &Method{}

	handlerType := reflect.TypeOf(handler)

	method, ok := handlerType.MethodByName(methodName)
	if ! ok {

	}

	methodObj := reflect.ValueOf(handler).MethodByName(methodName) // method obj

	methodType := method.Type

	args := make([]reflect.Value, methodType.NumIn()-1) // method params array
	argsTypes := make([]reflect.Type, methodType.NumIn()-1) // method params array

	// 遍历参数
	//for j := 1; j < methodType.NumIn(); j++ {
		//argType := methodType.In(j).Kind().String()
		//append(argsTypes, argType)

	//}

	_method.argsTypes = argsTypes
	_method.argsValues = args
	_method.methodValue = methodObj
	_method.methodType = methodType


	methodObj.Call(args)
	return _method
}

func TestNewRoute(t *testing.T) {
	//op := &Op{}
	//invoker := Build("/echo", op, "echo")
	//
	//key := "key"
	//value := "value"

	//invoker.invoke({key, value})
}

func register(fun func(name string)) {
	fun("")
}