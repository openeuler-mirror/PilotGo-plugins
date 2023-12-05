package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/atune-plugin/plugin"
)

func Upload(cookie string, filePath string, filename string) error {
	// 以二进制方式上传文件
	file := filepath.Join(filePath, filename)
	bodyBuf, contentType, err := GetUploadBody(file)
	if err != nil {
		return err
	}

	upload_addr := "http://" + plugin.GlobalClient.Server() + "/api/v1/upload?filename=" + filename
	request, err := http.NewRequest("POST", upload_addr, bodyBuf)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Cookie", cookie)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
	defer client.CloseIdleConnections()

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取返回结果
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("没获取到：%s", err.Error())
		return err
	}
	var respResult Response
	err = json.Unmarshal(responseBody, &respResult)
	if err != nil {
		logger.Error("解析出错:%s", err.Error())
		return err
	}
	if resp.StatusCode == http.StatusOK && respResult.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New(respResult.Message)
}

// 以二进制格式上传文件
func GetUploadBody(filename string) (*bytes.Reader, string, error) {
	bodyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return bytes.NewReader(bodyBytes), "", err
	}
	return bytes.NewReader(bodyBytes), "multipart/form-data", nil
}

type Response struct {
	StatusCode int         `json:"code"`
	Data       interface{} `json:"data"`
	Message    string      `json:"msg"`
}
