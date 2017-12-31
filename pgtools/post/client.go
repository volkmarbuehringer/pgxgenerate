package post

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	httpstat "github.com/tcnksm/go-httpstat"
)

func PostReq(ctx1 context.Context, url string, anfrage io.Reader) (httpstat.Result, io.ReadCloser, error) {

	var result httpstat.Result

	if req, err := http.NewRequest("POST", url, anfrage); err != nil {
		return result, nil, err
	} else {
		// Create a httpstat powered context
		req.Close = true
		req.Header.Add("Content-Type", "application/xml; charset=utf-8")
		req.Header.Set("Connection", "close")

		ctx := httpstat.WithHTTPStat(ctx1, &result)
		req = req.WithContext(ctx)

		if resp, err := client.Do(req); err != nil {
			//			daten, _ = weaxml.Tester()
			return result, nil, errors.Wrapf(err, "postclient")

		} else {
			return result, resp.Body, err

		}

	}

}

var client http.Client

func init() {

	conn, err := strconv.Atoi(os.Getenv("PU_CONNECT_TIMEOUT"))
	if err != nil {
		conn = 3000
	}

	httpt, err := strconv.Atoi(os.Getenv("PU_HTTP_TIMEOUT"))
	if err != nil {
		httpt = 3000
	}

	tr := http.Transport{
		DisableKeepAlives: true,
		Proxy:             http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   time.Duration(conn) * time.Millisecond,
			KeepAlive: 0,
		}).Dial,
		TLSHandshakeTimeout:   time.Duration(conn) * time.Second,
		ResponseHeaderTimeout: time.Duration(conn) * time.Second,
	}

	client = http.Client{Transport: &tr,
		Timeout: time.Duration(httpt) * time.Millisecond,
	}

}
