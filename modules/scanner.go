package modules

import (
	"bufio"
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
	flag.StringVar(&S.outPutPath, "o", "", "output file path.")
	flag.BoolVar(&S.check, "c", false, "get serverless ip's info.")
	flag.Parse()
	if S.check {
		S.scanModule = "c"
		return
	} else if S.outPutPath != "" {
		if S.isFileExists() {
			log.Fatalln("output file already exists.")
		}
		os.Create(S.outPutPath)
	}
	if S.filePath != "" {
		S.loadTargetsFile()
		return
	} else if S.ip == "" {
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
func (S *Scanner) scan() {
	Scanarams := map[string]string{"ip": S.ip, "ports": S.ports, "execParam": S.execParam}
	client := resty.New()
	headers := map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.42"}
	client.SetHeaders(headers)
	client.SetBaseURL(S.c["server"])
	client.SetTimeout(900 * time.Second)
	client.SetQueryParams(Scanarams)
	var resp *resty.Response
	var err error
	switch S.scanModule {
	case "k":
		resp, err = client.R().Get("/k")
		break
	case "f":
		resp, err = client.R().Get("/f")
		break
	case "c":
		resp, err = client.R().Get("/ip")
		break
	}
	if err != nil {
		log.Fatal(err)
	}
	S.Results = string(resp.Body())
}
func (S *Scanner) Run() {
	defer o.Close()
	S.init()
	if len(S.targets) != 0 {
		for _, v := range S.targets {
			S.ip = v
			if !S.checkIpAddress() {
				log.SetPrefix("[-] ")
				log.Println("Target " + v + "Format Error.")
				continue
			}
			log.SetPrefix("[*] ")
			log.Println("Scanning Target " + v)
			S.scan()
			S.outPutToFile()
			fmt.Println(S.Results)
		}
	} else {
		S.scan()
		S.outPutToFile()
		fmt.Println(S.Results)
	}
}
func (S *Scanner) loadTargetsFile() {
	file, err := os.Open(S.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ip := scanner.Text()
		S.targets = append(S.targets, ip)
	}
}
func (S *Scanner) outPutToFile() {
	if S.outPutPath != "" {
		o, err := os.Open(S.outPutPath)
		if err != nil {
			log.Fatal(err)
		}
		o.WriteString(S.Results + "\n")
	}
}
func (S *Scanner) isFileExists() bool {
	_, err := os.Lstat(S.outPutPath)
	return !os.IsNotExist(err)
}
