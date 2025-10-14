package exec

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type ScriptsRun struct {
	JobId         string
	ScriptType    string
	ScriptContent string
	Params        string
	TimeOut       int
}

func ExecScript(ip string, sr ScriptsRun) {
	addr := fmt.Sprintf("http://%s:8277/plugin/automation/script/exec", ip)
	dataMarshed, err := json.Marshal(sr)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	request, err := http.NewRequest("POST", addr, bytes.NewBuffer(dataMarshed))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer response.Body.Close()
}
