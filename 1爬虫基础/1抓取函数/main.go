package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


//Points:
//   1 使用http.Get， 返回Response
//   2 要对Response 进行close
//   3 结合http的状态码来进行一些错误处理


func main(){
	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun") //
	if err != nil{
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Error, Status code is :%s", resp.StatusCode)
	}
	all, err := ioutil.ReadAll(resp.Body) // resp.Body 不是字符串， 是一个Reader
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", all)
}


