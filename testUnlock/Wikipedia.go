package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func WikipediaEditable(c http.Client) result.Result {
	resp, err := url.GET(c, "https://zh.wikipedia.org/w/index.php?title=Wikipedia%3A%E6%B2%99%E7%9B%92&action=edit")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		return result.Result{Status: status.Failed}
	}

	if strings.Contains(bodyString, "Banned") {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
