package http

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func Request(requestUrl, method string, params map[string]string) ([]byte, error) {
	var res []byte
	var err error
	switch method {
	case "POST":
		res, err = Post(requestUrl, params)
	case "Delete":
		res, err = Delete(requestUrl, params)
	case "GET":
		res, err = Get(requestUrl)
	}
	return res, err
}

func Get(requestUrl string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	var response *http.Response
	var err error
	for i := 0; i < 2; i++ {
		response, err = c.Get(requestUrl)
		if err != nil {
			continue
		}
		break
	}
	if response != nil {
		defer response.Body.Close()
		body, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, err
}

func Post(requestUrl string, params map[string]string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	reqContentType := "application/x-www-form-urlencoded"
	reqm := make(url.Values)
	for k, v := range params {
		reqm.Add(k, v)
	}
	encode := reqm.Encode()
	var response *http.Response
	var err error
	for i := 0; i < 2; i++ {
		response, err = c.Post(requestUrl, reqContentType, strings.NewReader(encode))
		if err != nil {
			continue
		}
		break
	}
	if response != nil {
		defer response.Body.Close()
		body, err4 := ioutil.ReadAll(response.Body)
		if err4 != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, err
}

func Delete(requestUrl string, params map[string]string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	uv := make(url.Values)
	for k, v := range params {
		uv.Add(k, v)
	}

	var response *http.Response
	req, err := http.NewRequest("DELETE", requestUrl, strings.NewReader(uv.Encode()))
	for i := 0; i < 2; i++ {
		response, err = c.Do(req)
		if err != nil {
			continue
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		body, err4 := ioutil.ReadAll(response.Body)
		if err4 != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, err
}
