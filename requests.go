package rbd_iscsi_client

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)
func Request(apiUrl,method string,payload interface{},ContentType interface{},auth BasicAuth,timeout int) ([]byte,int,error){
	var req *http.Request
	if payload != nil {
		//check payload type
		if ContentType.(string) == "application/x-www-form-urlencoded"{
			data:=payload.(url.Values)
			req,_ = http.NewRequest(method,apiUrl,strings.NewReader(data.Encode()))
		}else{
			req,_ = http.NewRequest(method,apiUrl,payload.(io.Reader))
		}
	}else{
		req,_ = http.NewRequest(method,apiUrl,nil)
	}
	if req == nil {
		return nil,-1,nil
	}
	if ! reflect.DeepEqual(auth,BasicAuth{}) {
		req.SetBasicAuth(auth.Username,auth.Password)
	}

	req.Header.Set("User-Agent", UserAgent)
	if ContentType != nil {
		req.Header.Set("Content-Type",ContentType.(string))
	}
	req.Header.Set("Accept",JsonContentType)
	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	resp,err := client.Do(req)
	if err != nil {
		return nil,-1,err
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,-1,err
	}
	return body,resp.StatusCode,nil
}