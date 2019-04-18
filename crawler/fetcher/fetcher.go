package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


var rateLimiter = time.Tick(100*time.Millisecond)

func Fetch(url string) ([]byte, error)  {

	// "http://www.zhenai.com/zhenghun"
	<-rateLimiter
	resp,err := http.Get(url)
	if err!=nil{
		return nil,err
	}

	defer  resp.Body.Close()

	// 可能需要转码
	//bodyReader :=bufio.NewReader(resp.Body)
	//e := determineEncoding(bodyReader)
	//utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	//all,err := ioutil.ReadAll(utf8Reader)

	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)

}

// 转码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes ,err :=bufio.NewReader(r).Peek(1024)
	if err!=nil{
		//panic(err)
		log.Panicf("Fetcher code %s",err)
		return unicode.UTF8
	}

	e,_,_ :=charset.DetermineEncoding(bytes,"")

	return e
}

