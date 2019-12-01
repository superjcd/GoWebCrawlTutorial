package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

//Points:
//   1 使用text.transform 来对原来对Reader进行编码转化
//   2 使用charset.DetermineEncoding 来获取编码信息
//   3 使用bufio.NewReader 可以生成新对Reader


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

	// 判别内容的编码
    e := deteminEncoding(resp.Body)

	// 对Reader进行transoform
	utf8reader := transform.NewReader(resp.Body,
		e.NewEncoder())

	all, err := ioutil.ReadAll(utf8reader) // resp.Body 不是字符串， 是一个Reader
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func deteminEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024) // 读取前1024个字节
	if err != nil{
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}


