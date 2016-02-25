package trans

import (
	//"fmt"
	"github.com/nareix/curl"
	"regexp"
	"time"
)

var url string = "http://www.iciba.com/"

func getPage(word string) string {

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

func clean(page string) string {
	re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	result := re.ReplaceAllString(page, "")
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	result = re.ReplaceAllString(result, "")
	re, _ = regexp.Compile("\\r")
	result = re.ReplaceAllString(result, "")
	re, _ = regexp.Compile("\\n")
	result = re.ReplaceAllString(result, "")
	re, _ = regexp.Compile("\\t")
	result = re.ReplaceAllString(result, "")
	re, _ = regexp.Compile(" {2,}")
	result = re.ReplaceAllString(result, "")

	return result

}

func Trans(word string) string {
	page := clean(getPage(word))
	re, _ := regexp.Compile(`<ul class='base-list switch_part' >.*</ul><br/>`)
	result := re.FindStringSubmatch(page)

	return result[0]
}
