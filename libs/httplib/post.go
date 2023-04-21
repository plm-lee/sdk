package httplib

import (
	"bytes"
	"github.com/plm-lee/sdk/libs/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// PostForm post请求，Content-Type: form－data
func PostForm(url string, params map[string]interface{}) (body string, err error) {
	bf := new(bytes.Buffer)
	w := multipart.NewWriter(bf)
	for k, v := range params {
		w.WriteField(k, utils.ToString(v))
	}
	w.Close()

	req, _ := http.NewRequest("POST", url, bf)
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	body = string(b)
	return
}
