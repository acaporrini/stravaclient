package stravaclient

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
)

const baseUrl = "https://www.strava.com/api/v3"

type WebApi struct {
}

func (a *WebApi) GetAthlete(accessToken string) (io.Reader, error) {

	reqURL, _ := url.Parse(baseUrl + "/athlete")
	req := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {"Bearer " + accessToken},
		},
	}

	// send an HTTP request using `req` object
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res.Body, nil
}

func Openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
