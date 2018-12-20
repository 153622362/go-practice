package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


var rateLimiter = time.Tick(100 * time.Millisecond)
//读取页面源码
func Fetch(url string) ([]byte, error) {
	<-rateLimiter //限制速率

	//resp, err := http.Get( url) //seed 种子->request engine
	//
	//if err != nil {
	//	return nil, err
	//}
	//defer  resp.Body.Close()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll( resp.Body )
}