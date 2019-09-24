package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pydr/utils"
)

type Tel struct {
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type params struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    *Tel     `json:"tel"`
	Time   int64    `json:"time"`
	TplId  int      `json:"tpl_id"`
}

type response struct {
	Result   int    `json:"result"`
	Errormsg string `json:"errormsg"`
	Ext      string `json:"ext"`
	Fee      int    `json:"fee"`
	Sid      string `json:"sid"`
}

func (c *Client) Send(smsTpl int, mobile string, content ...string) (int, error) {
	nonce := genNonce(10000)
	timeStamp := time.Now().Unix()
	url := APIBASE + "sdkappid=" + c.appid + "&random=" + nonce
	sig := makeSign(c.secretKey, nonce, strconv.FormatInt(timeStamp, 10)[:10], mobile)
	p := params{
		Ext:    "",
		Extend: "",
		Params: content,
		Sig:    sig,
		Sign:   "",
		Tel: &Tel{
			Mobile:     mobile,
			Nationcode: "86",
		},
		Time:  timeStamp,
		TplId: smsTpl,
	}

	data, err := json.Marshal(p)
	if err != nil {
		return -1, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return -1, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := utils.Request(c.httpClient, req, 3)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var ret response
	if err = json.Unmarshal(body, &ret); err != nil {
		return -1, err
	}
	if ret.Result != 0 {
		return ret.Result, errors.New(ret.Errormsg)
	}

	return 0, nil
}
