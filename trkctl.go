package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"gopkg.in/yaml.v2"
)

var configuration Conf

type SummaryToday struct {
	Timestamps        []Timestamp   `json:"timestamps"`
	DifferenceFloat   float32       `json:"differenceFloat"`
	TotalAbsoluteTime time.Duration `json:"totalAbsoluteTime"`
}

// TimeStamp struct
type Timestamp struct {
	Date      time.Time `json:"date"`
	IsCheckin bool      `json:"isCheckin"`
}

type Conf struct {
	URL  string `yaml:"url"`
	Port string `yaml:"port"`
	Path string
}

func main() {
	// Reading the configuration file from user
	configuration.getConf()

	if len(os.Args[1:]) != 1 {
		errorStringHandler()
	}

	var argsWithProg string = os.Args[1]

	switch arg := argsWithProg; arg {
	case "checkin":
		timestamp := stampHandler(true)

		fmt.Printf("[Coffee!]☕ Good Morning Martin, you started working at %s !", timestamp.Date)

	case "checkout":
		timestamp := stampHandler(false)
		fmt.Printf("\n[Party]🎉 Closing Time - Go Home Martin, you stopped working at %s !", timestamp.Date)
		//TODO REST call to the server

		//TODO Print summary for todays working hours

	case "list":
		listAllHandler()

	case "help":
		helpStringHandler()

	case "status":
		fmt.Println("STATUS WIP")

	case "info":
		infoStringHandler()

	case "today":
		todaySummaryHandler()

	//Error message on wrong argument
	default:
		errorStringHandler()
	}
	//fmt.Printf(argsWithProg[0])
}

func stampHandler(isCheckin bool) Timestamp {
	json_data, err := json.Marshal(Timestamp{Date: time.Now(), IsCheckin: isCheckin})

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(configuration.URL+":"+configuration.Port+"/stamp", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	decoder := json.NewDecoder(resp.Body)
	var t Timestamp
	fmt.Print(t)
	err = decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	return t
}

func todaySummaryHandler() {
	var timestamps SummaryToday = getSummaryToday()
	loc, _ := time.LoadLocation("Europe/Berlin")
	fmt.Println("\n --- SUMMARY FOR TODAY --- \n\n")
	fmt.Println("NUM	DATE		TIME		TYPE")
	for i, stamp := range timestamps.Timestamps {
		fmt.Printf("%d	%s	%s	%b \n", i+1, stamp.Date.Format("2006-1-2"), stamp.Date.In(loc).Format("15:04:05"), stamp.IsCheckin)
	}

	fmt.Println(timestamps)

	fmt.Println("\n [====================] 100% \n")
}

func listAllHandler() {

	loc, _ := time.LoadLocation("Europe/Berlin")
	fmt.Println("NUM	DATE		TIME		TYPE")
	for i, stamp := range getAllTimestamps() {
		fmt.Printf("%d	%s	%s	%b \n", i+1, stamp.Date.Format("2006-1-2"), stamp.Date.In(loc).Format("15:04:05"), stamp.IsCheckin)
	}
}

func getSummaryToday() SummaryToday {
	resp, err := http.Get(configuration.URL + ":" + configuration.Port + "/summary/day/" + time.Now().Format("2006-01-02"))
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	decoder := json.NewDecoder(resp.Body)
	var timestamps SummaryToday
	err = decoder.Decode(&timestamps)
	if err != nil {
		panic(err)
	}
	return timestamps

}

func getAllTimestamps() []Timestamp {
	resp, err := http.Get(configuration.URL + ":" + configuration.Port + "/list")
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	decoder := json.NewDecoder(resp.Body)
	var timestamps []Timestamp
	err = decoder.Decode(&timestamps)
	if err != nil {
		panic(err)
	}
	return timestamps

}

func connectionTestHandler() string {
	resp, err := http.Get(configuration.URL + ":" + configuration.Port + "/ping")
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	decoder := json.NewDecoder(resp.Body)
	var pong string
	err = decoder.Decode(&pong)
	if err != nil {
		panic(err)
	}
	return pong
}

// UX and Config Features
func errorStringHandler() {
	err := fmt.Errorf("Wrong or missing arguments! Please call `trkctl help` to get help!")
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func (c *Conf) getConf() *Conf {
	user, _ := user.Current()
	homeDirectory := user.HomeDir

	paths := [2]string{homeDirectory + "/.trkconf.yaml", ".trkconf.yaml"}
	var yamlFile []byte
	for _, path := range paths {
		var err error = nil
		yamlFile = nil
		yamlFile, err = ioutil.ReadFile(path)
		if err != nil {
			continue
		}
		err = yaml.Unmarshal(yamlFile, c)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		c.Path = path
		break
	}

	if yamlFile == nil {
		err := fmt.Errorf("Missing Config File. Please place configs to ~/.trkconf.yaml or in current dir .trkconf.yaml")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	return c
}

// Generating help string
func helpStringHandler() {
	multiline := `- Using trkctl ⌚ -
 
trkctl is a CLI tool to communicate with the trkbox server
__

trkctl [command]
	
	checkin 	checks you in
	checkout	checks you out
	status		calculates current present time @ work		
	help		shows this help message
	info		connection test to the trkbox server
	
`

	fmt.Print(multiline)
}

// Generating help string
func infoStringHandler() {
	fmt.Printf("\nUsing Config from Path: %s", configuration.Path)
	fmt.Println(" \n \n--> Starting Connection Tests")
	fmt.Println("...sending ping")
	fmt.Printf("response: %s", connectionTestHandler())
}
