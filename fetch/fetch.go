package fetch

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Fetch(url string) ([]byte,error){


	client := &http.Client{}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	req,err :=http.NewRequest("GET",newUrl,nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	cookie1 := "sid=5693eb86-6d3c-41fa-8588-385fcd3589a1; bdVid=8329519008861681203; ec=CDYHPVAn-1625651639237-969ff4efbe1a5-912636910; FSSBBIl1UgzbN7NO=5DNbmKUc4XoJ4O7YsVZglKGpvakwtMkSVVkcWakbb0aqUUw3_iKFGFpaAJMU5pGhxMXmGw9Ee5n8tk9kpuWXtnq; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1625651644,1625712353,1625738058; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1625738195; _efmdata=NvmRRf6XXQgt151EMZa5ZGcCAZNVOjuifWJoXa7hBCfNIBjV1qYsVfWoUn9pcU4bUqwaNFAuBjqcaM40SGmz2uTpjB/soeAtP8INqrlMFQo=; _exid=o5gF2mecuIA8Qo0vE0XGQTWrAuUgJRoiG42mi0CW38YKTMknzAojp/WOVpo8dMaCuZ05AVnCf7CsAVStj+xzfg==; FSSBBIl1UgzbN7NP=53iMhTCkbneWqqqm_f3SxWaK77LKOVr1EaW00kDuT0e4ospXmpW6Dpn0ZbBOcqlqzcdt1eyuEFiaSaFo1DP7XcnPt_PfBK1t8edgioirfENMqJYewR7E__6ldblePESmnjT7MCz9CLqfsHJvmpyxkgfnlgEHegHZPPz8G1eyldObVp.Dwn9rMRjhzaiBLl3B_cOTpDv7l_fkBZzqdR6ZRh7kMsqY3faZ.1IiO4GpGmW3vLEzxfYSXKRnAP4Y0E4qfAmRIDKDGFIEtRlukEp1tBW"
	req.Header.Add("cookie", cookie1)

	resp,err := client.Do(req)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode !=http.StatusOK{
		return nil,fmt.Errorf("wrong status code:%d",resp.StatusCode)
	}
	e:=determineEncoding(resp.Body)
	utf8Reader :=transform.NewReader(resp.Body,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes,err := bufio.NewReader(r).Peek(1024)

	if err!=nil{
		log.Printf("fetcher err %v",err)
	}
	e,_,_ :=charset.DetermineEncoding(bytes,"")
	return e
}