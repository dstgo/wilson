package lobbyapi

import (
	"errors"
	"testing"
)

func TestLobbyRegions(t *testing.T) {
	client := New("")
	regions, err := client.GetCapableRegions()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(regions)
}

func TestLobbyServersOk(t *testing.T) {
	client := New("")
	servers, err := client.GetLobbyServers("ap-east-1", Steam.String())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(servers)
}

func TestLobbyServersFailed(t *testing.T) {
	client := New("")
	servers, err := client.GetLobbyServers("unknown", Steam.String())
	if err == nil {
		t.Error(errors.New("error must be non-nil"))
		return
	}
	t.Log(servers)
	t.Log(err)
}

func TestServerDetails(t *testing.T) {
	client := New("klei Token")
	servers, err := client.GetServerDetails("ap-east-1", "KU_nnMF5SAo")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(servers)
}

func TestParsedURL(t *testing.T) {
	url, err := parseURL(LobbyServersURL, map[string]any{
		"region":   "ap-east-1",
		"platform": "Steam",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if url != "https://lobby-v2-cdn.klei.com/ap-east-1-Steam.json.gz" {
		t.Error("Invalid")
		return
	}

	t.Log(url)
}
