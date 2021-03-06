package trans

import (
	"fmt"
    "github.com/parnurzeal/gorequest"
	"regexp"
	"time"
)

var url string = "http://www.iciba.com/"
var translated []string

func start() time.Time {
	return time.Now()
}

func end(s time.Time) {
	fmt.Printf("  **查询用时： %.2fs**\n\n", time.Now().Sub(s).Seconds())
}

func pageClean(page string) string {
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

func getPage(word string) string {

    _, body, errs := gorequest.New().Get(url + word).End()
    if errs == nil {
		return pageClean(body)
    }
    return ""
}

func dataClean(data string) {
	re, _ := regexp.Compile(`<li>.*?</li>`)
	results := re.FindAllStringSubmatch(data, -1)
	for _, v := range results {
		dataReorganize(v[0])
	}
}

func dataReorganize(data string) {
	re, _ := regexp.Compile("<.*?[^>]>")
	result := re.ReplaceAllString(data, " ")
	translated = []string{}
	translated = append(translated, result)
}

func output(translated []string) {
	for _, v := range translated {
		fmt.Println(v)
	}
}

func Trans(word string) {

	defer end(start())

	page := getPage(word)
	re, _ := regexp.Compile(`<ul class='base-list switch_part' >.*?</ul>`)
	result := re.FindStringSubmatch(page)

	fmt.Println("\n- - - - - - - - - -")
	if result != nil {
		dataClean(result[0])
		output(translated)
	} else {
		fmt.Println("  Illegal word... (- - |||")
	}
	fmt.Println("\n")
}
