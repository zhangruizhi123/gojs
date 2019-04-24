package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/robertkrimen/otto"
)

//定义全局的vm对象
var vm *otto.Otto = otto.New()

func gofunc(call otto.FunctionCall) otto.Value {
	val := call.Argument(0)
	obj := val.Object()
	name, _ := obj.Get("name")
	fmt.Println("gofunc:", obj, val, name)
	result, _ := vm.ToValue(2)
	return result
}

/*
读取文件
*/
func readFile(call otto.FunctionCall) otto.Value {
	name, err := call.Argument(0).ToString()
	if err == nil {
		data, err2 := ioutil.ReadFile(name)
		if err2 == nil {
			result, _ := vm.ToValue(string(data))
			return result
		} else {
			result, _ := vm.ToValue(err2)
			return result
		}
	} else {
		result, _ := vm.ToValue(err)
		return result
	}

}

//写入文件
func writeFile(call otto.FunctionCall) otto.Value {

	name, err1 := call.Argument(0).ToString()
	if err1 == nil {
		text, err2 := call.Argument(1).ToString()
		if err2 == nil {
			e := ioutil.WriteFile(name, []byte(text), 666)
			result, _ := vm.ToValue(e)
			return result
		} else {
			result, _ := vm.ToValue(err2)
			return result
		}
	} else {
		result, _ := vm.ToValue(err1)
		return result
	}

}

//列出文件
func listFile(call otto.FunctionCall) otto.Value {
	result := make(map[string]interface{}, 0)
	ph, err := call.Argument(0).ToString()
	if err == nil {
		fileList := make([]map[string]interface{}, 0)
		fileInfo, err := ioutil.ReadDir(ph)
		if err == nil {
			for _, info := range fileInfo {
				inf := make(map[string]interface{}, 0)
				inf["isdir"] = info.IsDir()
				inf["time"] = info.ModTime().Format("2006-01-02 15:04:05")
				inf["name"] = info.Name()
				inf["size"] = info.Size()
				fileList = append(fileList, inf)
			}
			result["success"] = 0
			result["messgae"] = "成功"
			result["data"] = fileList
			returns, _ := vm.ToValue(result)
			return returns
		} else {
			result["success"] = 201
			result["messgae"] = err.Error()
			returns, _ := vm.ToValue(result)
			return returns
		}

	} else {
		result["success"] = 101
		result["messgae"] = err.Error()
		msg, _ := vm.ToValue(result)
		return msg
	}
	result["success"] = 0
	result["messgae"] = "成功"
	msg, _ := vm.ToValue(result)
	return msg

}
func delFile(call otto.FunctionCall) otto.Value {
	result := make(map[string]interface{}, 0)
	ph, err := call.Argument(0).ToString()
	if err == nil {
		err := os.Remove(ph)
		if err == nil {
			result["success"] = 0
			result["messgae"] = "删除成功"
			msg, _ := vm.ToValue(result)
			return msg
		} else {
			result["success"] = 201
			result["messgae"] = err.Error()
			msg, _ := vm.ToValue(result)
			return msg
		}
	} else {
		result["success"] = 101
		result["messgae"] = err.Error()
		msg, _ := vm.ToValue(result)
		return msg
	}
	result["success"] = 0
	result["messgae"] = "成功"
	msg, _ := vm.ToValue(result)
	return msg
}

func main() {
	data, err := ioutil.ReadFile("main.js")
	if err == nil {
		vm.Set("arg", os.Args)
		vm.Set("gofunc", gofunc)
		vm.Set("readFile", readFile)
		vm.Set("writeFile", writeFile)
		vm.Set("listFile", listFile)
		vm.Set("delFile", delFile)

		//运行js
		_, err2 := vm.Run(string(data))
		if err2 != nil {
			fmt.Println(err2)
		}
		//调用js方法
		//val, _ := vm.Call("call", nil, 22)
		//fmt.Println(val)
	} else {
		fmt.Println("not find main.js")
	}

}
