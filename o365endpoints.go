package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type O365EndPoint struct {
	Id int
	ServiceArea string
	ServiceAreaDisplayName string
	URLs []string
	IPs []string
	TCPPorts string	
	ExpressRoute bool
	Category string
	Required bool
}

func getO365NWJSONTest() ([]byte, error) {
	input_file := "input_t1.json"
	data, err := os.ReadFile(input_file)
	if err != nil {
		return nil, fmt.Errorf("Failed to get json from input file")
	}
	return data, nil
}

func getO365NWJSON() ([]byte, error) {
	url := "https://endpoints.office.com/endpoints/worldwide?clientrequestid=b10c5ed1-bad1-445f-b386-b919946339a7"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get JSON from %s: %v", url, err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch body: %v", err)
	}
	return data, nil
}
func getO365Endpoints() ([]O365EndPoint, error) {
	jsonData, err := getO365NWJSON()
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(jsonData))

	var eps []O365EndPoint
	if err := json.Unmarshal(jsonData, &eps); err != nil {
		return nil, fmt.Errorf("failed to parse JSON %v", err)
	}
	t, _ := json.MarshalIndent(eps, "", "  ")
	fmt.Println(string(t))
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}