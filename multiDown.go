package main

import (
	"net/http"
	"io/ioutil"
	"strconv"
	"sort"
	"bytes"
	"errors"
)

// temp save bytes in channel
type content struct {
	id     int
	buffer []byte
}
type contents []content

func (c contents) Len() int {
	return len(c)
}
func (c contents) Less(i, j int) bool {
	return c[i].id < c[j].id
}
func (c contents) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Simple Multi-thread Download
func Download(url string, connNum int) ([]byte, error) {

	// check thread num
	if connNum < 1 || connNum > 33 {
		return nil, errors.New("1 < connNum < 32")
	}

	if connNum != 1 {

		// Get content-Length
		req, err := http.NewRequest("HEAD", url, nil)
		if err != nil {
			return nil, err
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		all, err := strconv.ParseInt(resp.Header.Get("content-Length"), 10, 64)
		per := all / int64(connNum)
		if err != nil {
			return nil, err
		}

		channel := make(chan content)
		for i := 0; i < connNum; i++ {
			s := strconv.FormatInt(per*int64(i)+1, 10)
			if s == "1" {
				s = "0"
			}
			e := strconv.FormatInt(per*int64(i+1), 10)
			if connNum-i == 1 {
				e = resp.Header.Get("content-Length")
			}
			go func(i int, s string, e string, channel chan content) {
				client2 := http.Client{}
				req2, err := http.NewRequest("GET", url, nil)
				if err != nil {
					panic(err)
					return
				}
				req2.Header.Set("Range", "bytes="+s+"-"+e)
				resp2, err := client2.Do(req2)
				if err != nil {
					panic(err)
					return
				}
				b, err := ioutil.ReadAll(resp2.Body)
				if err != nil {
					panic(err)
					return
				}
				channel <- content{id: i, buffer: b}
			}(i, s, e, channel)
		}
		con := make(contents, connNum)
		for i := 0; i < connNum; i++ {
			con[i] = <-channel
		}
		sort.Sort(con)
		var buf bytes.Buffer
		for i := 0; i < connNum; i++ {
			buf.Write(con[i].buffer)
		}
		return buf.Bytes(), nil
	} else {
		resp2, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
}

func DownloadToFile(filename string, url string, connNum int) error {
	b, err := Download(url, connNum)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0755)
}
