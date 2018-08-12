package greynoise

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// NewRequest prepares http.Request to call the Greynoise API
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/json; charset=UTF-8")
	req.Header.Add("Accept", "text/json")
	req.Header.Add("Key", c.Key)
	req.Header.Add("User-Agent", "gogreynoise")

	return req, nil
}

// Do will execute the http.Request and decode the result
func (wd *Client) Do(req *http.Request, v interface{}) error {
	if dump, err := httputil.DumpRequestOut(req, true); err == nil {
		os.Stdout.Write(dump)
	}

	resp, err := wd.Client.Do(req)
	if err != nil {
		return err
	}

	r := resp.Body
	defer r.Close()

	if true {
	} else if dump, err := httputil.DumpResponse(resp, true); err == nil {
		os.Stdout.Write(dump)

		resp.Body = ioutil.NopCloser(bytes.NewBuffer(dump))
	}

	r2 := r

	if resp.StatusCode != http.StatusOK {
		result := ErrorResult{}
		json.NewDecoder(r2).Decode(&result)
		return fmt.Errorf(result.Error)
	}

	err = json.NewDecoder(r2).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
