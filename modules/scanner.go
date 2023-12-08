package modules

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

var (
	O *os.File
	W *bufio.Writer
)

func (S *Scanner) init() {
	S.loadConfig()
	flag.IntVar(&S.scanModule, "t", 1, "use tool 1 is fscan,2 is kscan,3 is dirscan,deault is 1")
	flag.StringVar(&S.execParam, "p", "", "Scanner param for your selected mode")
	flag.StringVar(&S.outPutPath, "o", "", "output file path.")
	flag.BoolVar(&S.check, "c", false, "get serverless ip's info.")
	flag.Parse()
	if !S.check {
		if S.outPutPath != "" {
			if S.isFileExists() {
				log.SetPrefix("[-] ")
				log.Fatalln("output file already exists.")
			}
			O, _ = os.OpenFile(S.outPutPath, os.O_WRONLY|os.O_CREATE, 0766)
			W = bufio.NewWriter(O)
		}
		if S.execParam == "" {
			log.SetPrefix("[-] ")
			log.Println("Give Scanner param for your selected mode")
		}
		if !S.checkIsValidMode() {
			log.SetPrefix("[-] ")
			log.Fatalln("Invalid Scan Mode.(One Of 1,2,3)")
		}
	}

}
func (S *Scanner) loadConfig() {
	dataBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.SetPrefix("[-] ")
		log.Fatal(err)
	}
	err = yaml.Unmarshal(dataBytes, &S.c)
	if err != nil {
		log.SetPrefix("[-] ")
		log.Fatal(err)
	}
	if S.c["server"] == "" {
		log.SetPrefix("[-] ")
		log.Fatal("pls set server address at config.yaml")
	}
}
func (S *Scanner) scan() {
	client := resty.New()
	headers := map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.42"}
	client.SetHeaders(headers)
	client.SetBaseURL(S.c["server"])
	client.SetTimeout(900 * time.Second)
	if S.check {
		resp, err := client.R().Get("/ip")
		if err != nil {
			log.SetPrefix("[-] ")
			log.Fatalln("get ip info failed")
		}
		S.results = string(resp.Body())
		return
	}
	client.SetHeader("Content-Type", "application/json")
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	data := params{
		Type: S.scanModule,
		Args: S.execParam,
	}
	log.SetPrefix("[!] ")
	log.Println("Scanning.........................................")
	resp, err := client.R().SetBody(data).Post("/scan")
	if err != nil {
		log.SetPrefix("[-] ")
		log.Fatalln("scan task failed")
	}
	S.results = string(resp.Body())
}
func (S *Scanner) Run() {
	S.init()
	S.scan()
	fmt.Println(S.results)
	if !S.check {
		S.outPutToFile()
		W.Flush()
		O.Close()
	}
}
func (S *Scanner) outPutToFile() {
	if S.outPutPath != "" {
		_, err := W.WriteString(S.results + "\n")
		if err != nil {
			log.SetPrefix("[-] ")
			log.Fatal(err)
		}
	}
}
func (S *Scanner) isFileExists() bool {
	_, err := os.Lstat(S.outPutPath)
	return !os.IsNotExist(err)
}
func (S *Scanner) checkIsValidMode() bool {
	ModeValid := []int{1, 2, 3}
	for _, v := range ModeValid {
		if S.scanModule == v {
			return true
		}
	}
	return false
}
