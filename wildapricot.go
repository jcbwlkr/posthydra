package posthydra

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type WildApricotConfig struct {
	Key       string
	AccountId int
}

// TokenAuthHeader encodes the API key on a WildApricotConfig to the
// Authorization header needed to gain an access token
func (c *WildApricotConfig) TokenAuthHeader() string {
	raw := []byte(fmt.Sprintf("APIKEY:%s", c.Key))
	encoded := base64.StdEncoding.EncodeToString(raw)
	return fmt.Sprintf("Basic %s", encoded)
}

type WildApricotClient struct {
	Config *WildApricotConfig
	token  string
}

func NewWildApricotClient(c *Config) *WildApricotClient {
	return &WildApricotClient{Config: &c.WildApricot}
}

// AcquireToken must be called first to gain access to other API methods
func (w *WildApricotClient) AcquireToken() error {
	c := new(http.Client)
	var req *http.Request
	var res *http.Response
	var err error

	scope := strings.NewReader("grant_type=client_credentials&scope=events")
	req, err = http.NewRequest("POST", "https://oauth.wildapricot.org/auth/token", scope)
	req.Header.Set("Authorization", w.Config.TokenAuthHeader())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err = c.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("Authentication request failed with status %s", res.Status)
	}

	js := make(map[string]interface{})

	if err = json.NewDecoder(res.Body).Decode(&js); err != nil {
		return err
	}
	if _, ok := js["access_token"]; !ok {
		return fmt.Errorf("No access token in WildApricot response")
	}
	w.token = js["access_token"].(string)

	return nil
}

// Read performs the work of fetching events from WildApricot's API
func (w *WildApricotClient) Read() ([]*Event, error) {
	// Our return values
	e := make([]*Event, 0)
	var err error

	// Our http objects
	c := new(http.Client)
	var req *http.Request
	var res *http.Response

	err = w.AcquireToken()
	if err != nil {
		return e, err
	}

	// With the newly acquired token in hand, make a request for Events
	url := fmt.Sprintf("https://api.wildapricot.org/v2/Accounts/%d/events", w.Config.AccountId)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return e, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.token))

	res, err = c.Do(req)
	if err != nil {
		return e, err
	}
	if res.StatusCode != 200 {
		return e, fmt.Errorf("Events request failed with status %s", res.Status)
	}

	var waRes WildApricotResponse

	if err = json.NewDecoder(res.Body).Decode(&waRes); err != nil {
		return e, err
	}

	for _, event := range waRes.Events {
		e = append(e, &Event{
			Title:    event.Name,
			Start:    event.StartDate,
			End:      event.EndDate,
			URL:      event.URL,
			Location: event.Location,
		})
	}

	return e, nil
}

// WildApricotResponse defines the response from the WildApricot Events API.
// Used for unmarshalling the response to JSON.
type WildApricotResponse struct {
	Events []struct {
		StartDate                   string
		EndDate                     string
		Location                    string
		RegistrationEnabled         bool
		RegistrationsLimit          int
		PendingRegistrationsCount   int
		ConfirmedRegistrationsCount int
		CheckedInAttendeesNumber    int
		Tags                        []string
		AccessLevel                 string
		ID                          int
		URL                         string
		Name                        string
	}
}
