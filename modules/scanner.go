package modules

import (
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
	"time"
)

func (S *Scanner) init() {
	S.loadConfig()
	flag.StringVar(&S.ip, "t", "", "ip you want to scan.")
	flag.StringVar(&S.ports, "p", "", "ports you want to scan.(optional with the scan program default)")
	flag.StringVar(&S.scanModule, "m", "k", "scan modules.f for fscan,k for kscan,a for all.")
	flag.StringVar(&S.execParam, "e", "", "additional scan param.just useful in single scan module.")
	flag.StringVar(&S.filePath, "f", "", "multi scan source.")
	flag.BoolVar(&S.check, "c", false, "get serverless ip's info.")
	flag.Parse()
	if S.check {
		S.scanModule = "c"
	} else if S.ip == "" || !S.checkIpAddress() {
		log.Fatal("pls give me a valid ip.")
	}
}
func (S *Scanner) checkIpAddress() bool {
	address := net.ParseIP(S.ip)
	if address != nil {
		return true
	} else {
		return false
	}
}
func (S *Scanner) loadConfig() {
	dataBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(dataBytes, &S.c)
	if err != nil {
		log.Fatal(err)
	}
}
func (S *Scanner) start() {
	ScanUrl := fmt.Sprintf("%s?ip=%s&ports=%s&modules=%s&exeParam=%s", S.c["server"], S.ip, S.ports, S.scanModule, S.execParam)
	client := resty.New()
	headers := map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.42"}
	client.SetHeaders(headers)
	client.SetTimeout(900 * time.Second)
	resp, err := client.R().Get(ScanUrl)
	if err != nil {
		log.Fatal(err)
	}
	S.Results = string(resp.Body())
}
func (S *Scanner) Run() {
	S.init()
	S.start()
}
