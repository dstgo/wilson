package lobbyapi

import (
	"errors"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
)

// New returns a new instance of lobby client with klei token
func New(token string) *Client {
	return &Client{client: resty.New(), token: token}
}

func NewWith(token string, client *resty.Client) *Client {
	return &Client{client: client, token: token}
}

// Client is dst lobby http client, interact with lobby server and returns server information
type Client struct {
	client *resty.Client
	token  string
}

// GetCapableRegions returns a list of available regions that can be used in other api
// GET https://lobby-v2-cdn.klei.com/regioncapabilities-v2.json
func (c *Client) GetCapableRegions() (Regions, error) {
	response, err := c.client.R().Get(LobbyRegionURL)
	if err != nil {
		return Regions{}, err
	}

	// request failed
	if response.StatusCode() != http.StatusOK {
		return Regions{}, errors.New(string(response.Body()))
	}

	var regions Regions
	if err := sonic.Unmarshal(response.Body(), &regions); err != nil {
		return Regions{}, err
	}
	return regions, err
}

// GetLobbyServers returns a list of lobby servers with specified region and platform
// GET https://lobby-v2-cdn.klei.com/{region}-{platform}.json.gz
func (c *Client) GetLobbyServers(region string, platform string) (Servers, error) {
	url, err := parseURL(LobbyServersURL, map[string]any{
		"region":   region,
		"platform": platform,
	})
	if err != nil {
		return Servers{}, err
	}

	response, err := c.client.R().Get(url)
	if err != nil {
		return Servers{}, err
	}

	// request failed
	if response.StatusCode() != http.StatusOK {
		return Servers{}, errors.New(string(response.Body()))
	}

	var servers Servers
	if err := sonic.Unmarshal(response.Body(), &servers); err != nil {
		return Servers{}, err
	}

	for i, server := range servers.List {
		servers.List[i].Region = region
		if server.TagStr != "" {
			servers.List[i].Tags = strings.Split(server.TagStr, ",")
		}
	}

	return servers, nil
}

// GetServerDetails returns the details information for the specified server by rowId
// POST https://lobby-v2-{region}.klei.com/lobby/read
func (c *Client) GetServerDetails(region string, rowId string) (ServerDetails, error) {
	url, err := parseURL(LobbyDetailsURL, map[string]any{
		"region": region,
	})
	if err != nil {
		return ServerDetails{}, err
	}

	if len(c.token) == 0 {
		return ServerDetails{}, errors.New("klei token is required")
	}

	// prepare query body
	body := map[string]any{
		"__token":  c.token,
		"__gameId": GameId,
		"query": map[string]any{
			"__rowId": rowId,
		},
	}

	bytes, err := sonic.Marshal(body)
	if err != nil {
		return ServerDetails{}, err
	}

	// send request
	response, err := c.client.R().SetBody(bytes).Post(url)
	if err != nil {
		return ServerDetails{}, err
	}

	// request failed
	if response.StatusCode() != http.StatusOK {
		return ServerDetails{}, errors.New(string(response.Body()))
	}

	var detailResp struct {
		List []ServerDetails `json:"GET"`
	}

	if err := sonic.Unmarshal(response.Body(), &detailResp); err != nil {
		return ServerDetails{}, err
	}

	if len(detailResp.List) > 0 {
		details := detailResp.List[0]
		details.Region = region
		if details.TagStr != "" {
			details.Tags = strings.Split(details.TagStr, ",")
		}
		// parse lua script
		return parsedLuaDetails(details)
	}
	return ServerDetails{}, nil
}
