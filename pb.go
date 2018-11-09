package proxy_bonanza

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	API_VERSION = "1"
	BASE_PATH   = "https://api.proxybonanza.com"
)

type PB struct {
	client     *http.Client
	apiKey     string
	apiVersion string
}

// Return a new ProxyBonanza client for requests their API,
// using a custom http.Client and specified API version.
func NewCustomClientWithVersion(c *http.Client, apiKey, apiVersion string) *PB {
	return &PB{
		c,
		apiKey,
		apiVersion,
	}
}

// Return a new ProxyBonanza client for requests their API,
// using a custom http.Client.
func NewCustomClient(c *http.Client, apiKey string) *PB {
	return NewCustomClientWithVersion(c, apiKey, API_VERSION)
}

// Return a new ProxyBonanza client for requesting their API,
// using specified API version
func NewWithVersion(apiKey, apiVersion string) *PB {
	c := &http.Client{}

	return NewCustomClientWithVersion(c, apiKey, apiVersion)
}

// Return a new ProxyBonanza client for requesting their API,
// defaults to latest API version
func New(apiKey string) *PB {
	c := &http.Client{}

	return NewCustomClientWithVersion(c, apiKey, API_VERSION)
}

// List of active proxy plans in user account.
func (pb *PB) GetPlans() (plansResponse, error) {
	var pr plansResponse

	endpoint := "userpackages.json"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v%s/%s", BASE_PATH, pb.apiVersion, endpoint), nil)
	if err != nil {
		return pr, err
	}
	req.Header.Add("Authorization", pb.apiKey)
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := pb.client.Do(req)
	if err != nil {
		return pr, err
	}
	defer resp.Body.Close()

	if strings.Contains(strings.ToLower(resp.Header.Get("Content-Encoding")), "gzip") {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return pr, err
		}
		err = json.NewDecoder(gz).Decode(&pr)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&pr)
	}

	return pr, err
}

// Details of proxy plan including list of proxy IPs.
func (pb *PB) GetPlan(packageID int) (planResponse, error) {
	var pr planResponse

	endpoint := fmt.Sprintf("userpackages/%v.json", packageID)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v%s/%s", BASE_PATH, pb.apiVersion, endpoint), nil)
	if err != nil {
		return pr, err
	}
	req.Header.Add("Authorization", pb.apiKey)
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := pb.client.Do(req)
	if err != nil {
		return pr, err
	}
	defer resp.Body.Close()

	if strings.Contains(strings.ToLower(resp.Header.Get("Content-Encoding")), "gzip") {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return pr, err
		}
		err = json.NewDecoder(gz).Decode(&pr)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&pr)
	}

	return pr, err
}

// List of all authentication IPs in user account.
func (pb *PB) GetAuthIPS() (authIPSResponse, error) {
	var ar authIPSResponse

	endpoint := "authips.json"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v%s/%s", BASE_PATH, pb.apiVersion, endpoint), nil)
	if err != nil {
		return ar, err
	}
	req.Header.Add("Authorization", pb.apiKey)
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := pb.client.Do(req)
	if err != nil {
		return ar, err
	}
	defer resp.Body.Close()

	if strings.Contains(strings.ToLower(resp.Header.Get("Content-Encoding")), "gzip") {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return ar, err
		}
		err = json.NewDecoder(gz).Decode(&ar)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&ar)
	}

	return ar, err
}

// Add new authentication IP to proxy plan. POST parameters: ip, userpackage_id.
// Returns ID of created authentication IP.
func (pb *PB) AddAuthIP(authIP string, packageID int) (AuthIPID int, err error) {
	var ar addAuthIPResponse

	form := url.Values{}
	form.Add("ip", authIP)
	form.Add("userpackage_id", strconv.Itoa(packageID))

	endpoint := "authips.json"
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v%s/%s", BASE_PATH, pb.apiVersion, endpoint), strings.NewReader(form.Encode()))
	if err != nil {
		return ar.AuthIPID.ID, err
	}
	req.PostForm = form
	req.Header.Add("Authorization", pb.apiKey)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := pb.client.Do(req)
	if err != nil {
		return ar.AuthIPID.ID, err
	}
	defer resp.Body.Close()

	if strings.Contains(strings.ToLower(resp.Header.Get("Content-Encoding")), "gzip") {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return ar.AuthIPID.ID, err
		}
		err = json.NewDecoder(gz).Decode(&ar)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&ar)
	}

	if ar.AuthIPID == nil {
		return 0, err
	}

	return ar.AuthIPID.ID, err
}

// Remove authentication IP ID from proxy plan
// No error indicates successful.
func (pb *PB) RemoveAuthIP(authIPID int) error {
	r := struct {
		Success bool `json:"success"`
	}{}

	endpoint := fmt.Sprintf("authips/%v.json", authIPID)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v%s/%s", BASE_PATH, pb.apiVersion, endpoint), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", pb.apiKey)
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := pb.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if strings.Contains(strings.ToLower(resp.Header.Get("Content-Encoding")), "gzip") {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		err = json.NewDecoder(gz).Decode(&r)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&r)
	}

	if err == nil {
		if r.Success == false {
			err = errors.New("unexpected error: unable to remove auth ip")
		}
	}

	return err
}

// Data transfer usage stats for the last 30 days.
func (pb *PB) MonthsPackageStats(packageID int) (monthsStatsResponse, error) {
	var r monthsStatsResponse

	endpoint := fmt.Sprintf("userpackagedailystats/%v.json", packageID)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v%s/%s", BASE_PATH, pb.apiVersion, endpoint), nil)
	if err != nil {
		return r, err
	}
	req.Header.Add("Authorization", pb.apiKey)
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := pb.client.Do(req)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	if strings.Contains(strings.ToLower(resp.Header.Get("Content-Encoding")), "gzip") {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return r, err
		}
		err = json.NewDecoder(gz).Decode(&r)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&r)
	}

	return r, err
}

// Data transfer for the last 24 hours.
func (pb *PB) TodaysPackageStats(packageID int) (todaysStatsResponse, error) {
	var r todaysStatsResponse

	endpoint := fmt.Sprintf("userpackagehourlystats/%v.json", packageID)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v%s/%s", BASE_PATH, pb.apiVersion, endpoint), nil)
	if err != nil {
		return r, err
	}
	req.Header.Add("Authorization", pb.apiKey)
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := pb.client.Do(req)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	if strings.Contains(strings.ToLower(resp.Header.Get("Content-Encoding")), "gzip") {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return r, err
		}
		err = json.NewDecoder(gz).Decode(&r)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&r)
	}

	return r, err
}
