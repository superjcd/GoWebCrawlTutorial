package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	unicode2 "golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
)


func Fetch(url string) ([]byte, error){
	resp, err := http.Get(url) //
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("Error status code with response code of %d \n",
			resp.StatusCode)
	}

	// 判别内容的编码
	e := deteminEncoding(resp.Body)

	// 对Reader进行transoform
	utf8reader := transform.NewReader(resp.Body,
		e.NewEncoder())

	//
	return ioutil.ReadAll(utf8reader) // resp.Body 不是字符串， 是一个Reader

}


func deteminEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024) // 读取前1024个字节
	if err != nil{
		return unicode2.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}