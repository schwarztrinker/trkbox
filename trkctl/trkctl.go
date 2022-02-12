package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"sort"

	"gopkg.in/yaml.v2"
)

type Timestamp struct {
	Time string `json:"time"`
	Date string `json:"date"`
	Type string `json:"type"`
}

type Conf struct {
	URL  string `yaml:"url"`
	Port string `yaml:"port"`
	Path string
}

func main() {

	if len(os.Args[1:]) != 1 {
		errorStringHandler()
	}

	// Reading the configuration file from user
	var c Conf
	c.getConf()
	//fmt.Println(c)

	var argsWithProg string = os.Args[1]

	switch arg := argsWithProg; arg {
	case "checkin":
		timestamp := checkInHandler(c)

		fmt.Printf("[Coffee!]☕ Good Morning Martin, you started working at %s %s!", timestamp.Time, timestamp.Date)

	case "checkout":
		fmt.Println("\n[Party]🎉 Closing Time - Go Home Martin, you stopped working at TODO(TIMESTAMP) !")
		//TODO REST call to the server

		//TODO Print summary for todays working hours

	case "list":
		listAllHandler(c)

	case "help":
		helpStringHandler()

	case "status":
		fmt.Println("STATUS WIP")

	case "info":
		infoStringHandler(c)

	//Error message on wrong argument
	default:
		errorStringHandler()
	}
	//fmt.Printf(argsWithProg[0])
}

func listAllHandler(c Conf) {
	resp, err := http.Get(c.URL + ":" + c.Port + "/list")
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

	fmt.Println("NUM	DATE		TIME		TYPE")
	sort.Slice(timestamps[:], func(i, j int) bool {
		return timestamps[i].Date > timestamps[j].Date
	})
	for i, stamp := range timestamps {
		fmt.Printf("%d	%s	%s	%s \n", i+1, stamp.Date, stamp.Time, stamp.Type)
	}
}

func connectionTestHandler(c Conf) string {
	resp, err := http.Get(c.URL + ":" + c.Port + "/ping")
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

func checkInHandler(c Conf) Timestamp {
	resp, err := http.Get(c.URL + ":" + c.Port + "/checkin")
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	decoder := json.NewDecoder(resp.Body)
	var t Timestamp
	err = decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	return t
}

func checkOutHandler() {

}

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
func infoStringHandler(c Conf) {
	fmt.Printf("\nUsing Config from Path: %s", c.Path)
	fmt.Println(" \n \n--> Starting Connection Tests")
	fmt.Println("...sending ping")
	fmt.Printf("response: %s", connectionTestHandler(c))
}
