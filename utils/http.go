package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/OAB/utils/log"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

// Request http request
type Request struct {
	header http.Header
	query  url2.Values
	body   *bytes.Reader
	client http.Client
}

// Response http response
type Response struct {
	Code int
	Data []byte
}

func (r *Response) Parse(data interface{}) error {
	return json.Unmarshal(r.Data, data)
}

// NewHTTPCli new a http client
func NewHTTPCli() Request {
	return Request{
		client: http.Client{Timeout: time.Second * 30},
		query:  url2.Values{},
		body:   bytes.NewReader([]byte("")),
		header: http.Header{},
	}
}

// SetHeader set header
func (r *Request) SetHeader(key, val string) {
	r.header.Set(key, val)
}

// SetQuery set query
func (r *Request) SetQuery(key, val string) {
	r.query.Set(key, val)
}

func (r *Request) GetQuery() string {
	return r.query.Encode()
}

// SetBody set body
func (r *Request) SetBody(data interface{}) {
	ret, _ := json.Marshal(data)
	r.body = bytes.NewReader(ret)
}

// SetTimeout set timeout
func (r *Request) SetTimeout(t time.Duration) {
	r.client.Timeout = t
}

func (r *Request) Get(url string) (*Response, error) {
	return r.Request("GET", url)
}

func (r *Request) Post(url string) (*Response, error) {
	return r.Request("POST", url)
}

func (r *Request) Put(url string) (*Response, error) {
	return r.Request("PUT", url)
}

func (r *Request) Delete(url string) (*Response, error) {
	return r.Request("DELETE", url)
}

func (r *Request) Request(method, url string) (*Response, error) {
	var (
		resp = &Response{}
		err  error
		re   *http.Request
		ret  *http.Response
	)
	defer func() {
		if err != nil {
			log.Errorf("[utils.http.Request] request: %s, err: %s", url, err.Error())
		}
	}()
	if len(r.query) != 0 {
		sed := "&"
		if strings.Index(url, "?") == -1 {
			sed = "?"
		}
		url = fmt.Sprintf("%s%s%s", url, sed, r.query.Encode())
	}
	re, err = http.NewRequest(method, url, r.body)
	if err != nil {
		return resp, err
	}
	//log.Debugf("http url: %s", re.URL.String())
	re.Header = r.header
	ret, err = r.client.Do(re)
	if ret != nil {
		defer ret.Body.Close()
	}
	if err != nil {
		return resp, err
	}
	resp.Code = ret.StatusCode
	resp.Data, err = ioutil.ReadAll(ret.Body)
	return resp, err
}
