package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"
)

func main1() {
	//conn,err := net.Dial("tcp","www.baidu.com:443")
	conn,err := tls.Dial("tcp","www.baidu.com:443",nil)
	checkErr(err)
	defer conn.Close()

	var body string = `HEAD / HTTP/1.1
Host: www.baidu.com
Connection: close
Pragma: no-cache
Cache-Control: no-cache
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.62 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9
Cookie: BIDUPSID=65A798B44CCE10A5BCD39F684F344275; PSTM=1554367203; BD_UPN=12314353; BDUSS=NQVkp4bVEtR1ZEUUpjMzgwNndHR2VQNFJRUmp3MEZOd0JmbXR-ZGpvLUNzRU5kSVFBQUFBJCQAAAAAAAAAAAEAAABIhaVMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIIjHF2CIxxdcj; BAIDUID=DD27FFEBDEC9F3E043EEE544A5AFC7DA:FG=1; MCITY=-%3A; ispeed_lsm=2; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BD_HOME=1; H_PS_PSSID=1454_21092_29523_29521_29720_29567_29221_26350; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; delPer=0; BD_CK_SAM=1; PSINO=6; COOKIE_SESSION=2838_0_9_1_16_10_0_0_9_6_2_0_0_0_0_0_1567132769_0_1567417615%7C9%2336_93_1565325857%7C9; H_PS_645EC=86b3mugFPu9%2BrRsECFh60jp%2F6EqndVBDMV50emhHiaVJUAemFqnKx11iCvBX8o%2BMlaDk

`
	_, errWrite := conn.Write([]byte(body))
	checkErr(errWrite)

	var buf [1024]byte

	for {
		n, errRead := conn.Read(buf[0:])
		if errRead != nil {
			if errRead == io.EOF {
				break
			}else {
				log.Fatalln(errRead)
			}
		}
		fmt.Println(string(buf[0:n]))
	}


}

func checkErr(err error){
	bufio.NewReader(os.Stdin)
	if err != nil{
		log.Fatalln(err)
	}
}
