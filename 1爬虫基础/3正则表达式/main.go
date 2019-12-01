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
	"regexp"
)

//Points:
// 1 使用regexp库来对文本信息进行正则匹配， 包括结合FindaAll方法对使用等等。其他对是正则表达式本身， 就不多介绍


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

	//
	all, err := ioutil.ReadAll(utf8reader) // resp.Body 不是字符串， 是一个Reader
	if err != nil{
		panic(err)
	}
	printCityList(all)
}



func deteminEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024) // 读取前1024个字节
	if err != nil{
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte){
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]+>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		fmt.Printf("City: %s, Url: %s \n", match[2], match[1] )
		}
	fmt.Printf("We get %d cities", len(matches))
	}



