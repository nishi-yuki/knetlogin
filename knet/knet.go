package knet

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const NetworkError = "Network Error"

func IsInternetAvailable() bool {
	res, err := http.Get("http://detectportal.firefox.com/success.txt")
	if err != nil {
		return false
	}
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	result := string(body)
	if strings.Contains(result, "success") {
		return true
	}
	return false
}

func Login(uid string, pass string) error {
	var (
		res *http.Response
		doc *goquery.Document
		err error
	)

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}

	client := &http.Client{
		Jar: jar,
	}

	res, err = client.Get("https://netauth.cis.kit.ac.jp/auth/login.php")
	if err != nil {
		return err
	}

	postUrl := res.Request.URL.String()

	/* id password -> SAML	*/
	payload := url.Values{}
	payload.Add("j_username", uid)
	payload.Add("j_password", pass)
	payload.Add("_eventId_proceed", "")
	res, err = client.PostForm(postUrl, payload)
	if err != nil {
		return err
	}

	/* SAMLの解析 */
	doc, _ = goquery.NewDocumentFromResponse(res)
	postUrl, payload, _ = formParser(doc.Find("form"))
	//log.Println(doc.Text())
	//log.Println(postUrl)

	/* SAML -> login.php uid pwd */
	res, err = client.PostForm(postUrl, payload)
	if err != nil {
		return err
	}

	doc, _ = goquery.NewDocumentFromResponse(res)
	postUrl, payload, _ = formParser(doc.Find("form"))
	//log.Println(doc.Text())
	//log.Println(postUrl)
	//log.Println(payload)

	res, err = client.PostForm(postUrl, payload)
	if err != nil {
		return err
	}

	doc, _ = goquery.NewDocumentFromResponse(res)
	//log.Println(doc.Text())
	return nil
}

func formParser(form *goquery.Selection) (string, url.Values, bool) { //エラーの返し方がわからんのでこうした．直しといて．
	postUrl, e := form.Attr("action")
	if !e {
		return "", nil, true
	}
	data := url.Values{}
	form.Find("input").Each(func(_ int, input *goquery.Selection) {
		name, en  := input.Attr("name")
		value, ev := input.Attr("value")
		if !en || !ev {
			return
		}
		data.Add(name, value)
	})
	return postUrl, data, false
}

