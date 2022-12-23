package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IPInfoResponse struct {
	IP       string `json:"ip,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	City     string `json:"city,omitempty"`
	Region   string `json:"region,omitempty"`
	Country  string `json:"country,omitempty"`
	Loc      string `json:"loc,omitempty"`
	Org      string `json:"org,omitempty"`
	Postal   string `json:"postal,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

func GetURLRequest(ip string) string {
	return fmt.Sprintf("https://ipinfo.io/%s/json", ip)
}

func CallAPI(ip string) (IPInfoResponse, error) {
	var response IPInfoResponse
	var err error
	var client = &http.Client{Timeout: 10 * time.Second}
	url := GetURLRequest(ip)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, err
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(request)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response, fmt.Errorf("status : %s", res.Status)
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return response, err
	}

	return response, nil
}
