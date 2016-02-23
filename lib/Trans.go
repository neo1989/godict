package trans

import (
	"fmt"
	"github.com/nareix/curl"
	"time"
)

var url string = "http://www.iciba.com/"

func Trans(word string) string {

	req := curl.Get(url + word)

	req.DialTimeout(time.Second * 10) // TCP Connection Timeout
	req.Timeout(time.Second * 30)     // Download Timeout

	req.Progress(func(p curl.ProgressStatus) {}, time.Second)

	res, err := req.Do()

	if err == nil {
		return res.Body
	} else {
		return ""
	}

}
