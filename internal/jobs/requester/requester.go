package requester

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

type Requester struct {
	Exp    string
	URL    url.URL
	Method string
	Body   *string
}

func (r *Requester) Expression() string {
	return r.Exp
}

func (r *Requester) Execute() func() {
	return func() {
		c := &http.Client{
			Timeout: time.Minute,
		}

		// var bodyReader *bytes.Reader
		// if r.Body != nil {
		// 	bodyReader = bytes.NewReader([]byte(*r.Body))
		// }

		req, err := http.NewRequest(r.Method, r.URL.String(), nil)
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
