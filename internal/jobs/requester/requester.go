package requester

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Requester struct {
	Exp    string
	URL    url.URL
	Method string
	Body   any
}

func (r *Requester) Expression() string {
	return r.Exp
}

func (r *Requester) Execute() func() {
	return func() {
		c := http.DefaultClient
		bs, err := json.Marshal(r.Body)
		if err != nil {
			log.Println("error sending request:", err)
			return
		}

		req, err := http.NewRequest(r.Method, r.URL.String(), bytes.NewReader(bs))
		if err != nil {
			log.Println("error assembling request:", err)
			return
		}

		_, err = c.Do(req)
		if err != nil {
			log.Println("error sending request:", err)
			return
		}
	}
}
