package weather

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	host = "https://api.openweathermap.org/data/2.5"
)

type Coord struct {
	Lat float64
	Lon float64
}

type City struct {
	Id         int
	Name       string
	Coord      Coord
	Country    string
	Population int
	Timezone   int
	Sunrise    int
	Sunset     int
}

type ForecastList struct {
	Dt     int
	Pop    float64
	Dt_Txt string
}

type ForecastResponse struct {
	Cod     string
	Message int
	Cnt     int
	List    []ForecastList
	City    City
}

type Client struct {
	Client *http.Client
	AppId  string
}

func NewClient(client *http.Client, appId string) *Client {
	return &Client{
		client,
		appId,
	}
}

func DefaultClient(appId string) *Client {
	return NewClient(http.DefaultClient, appId)
}

func (c *Client) GetForecast(ctx context.Context, city string) (*ForecastResponse, error) {
	endpoint := fmt.Sprintf("%s/forecast", host)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	req.WithContext(ctx)

	q := req.URL.Query()
	q.Add("appid", c.AppId)
	q.Add("q", city)
	req.URL.RawQuery = q.Encode()

	body, err := get(c, req)
	if err != nil {
		return nil, err
	}

	response := &ForecastResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("decode forecast: %w", err)
	}

	return response, nil
}

func get(c *Client, req *http.Request) ([]byte, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("response error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	return body, nil
}
