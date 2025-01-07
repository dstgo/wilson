package lobbyapi

import (
	"bytes"
	"errors"
	"strings"
	"text/template"

	"github.com/dstgo/wilson/framework/dontstarve/luax"
)

// parse url from template
func parseURL(url string, params map[string]any) (string, error) {
	tmpl, err := template.New("lobby").Parse(url)
	if err != nil {
		return "", err
	}
	buff := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buff, params); err != nil {
		return "", err
	}
	return buff.String(), nil
}

func parsedLuaDetails(details ServerDetails) (ServerDetails, error) {

	// days info
	err := parsedDaysInfo(details.Data, &details)
	if err != nil {
		return ServerDetails{}, err
	}

	// player info
	playersInfo, err := parsePlayersInfo(details.OnlinePlayers)
	if err != nil {
		return ServerDetails{}, err
	}
	details.Details.Players = playersInfo

	// mods info
	details.Details.ModsInfo = parseModsInfo(details.Mods)

	return details, nil
}

// parse days info from lua script
func parsedDaysInfo(luaScript string, details *ServerDetails) error {
	if details == nil {
		return errors.New("nil details")
	}

	vm := luax.NewVM()
	defer vm.Close()
	if err := vm.DoString(luaScript); err != nil {
		return err
	}

	table := luax.LValue(vm.Get(-1)).ToTable()

	vm.Pop(1)

	details.Details.Day = table.GetInt("key")
	details.Details.DayElapsedInSeason = table.GetInt("dayselapsedinseason")
	details.Details.DaysLeftInSeason = table.GetInt("daysleftinseason")

	return nil
}

// parse players info from lua script
func parsePlayersInfo(luaScript string) ([]Player, error) {
	vm := luax.NewVM()
	defer vm.Close()
	if err := vm.DoString(luaScript); err != nil {
		return nil, err
	}

	var players []Player

	playerArray := luax.LValueToArray(vm.Get(-1))
	for _, value := range playerArray {
		playerTable := luax.LValue(value).ToTable()
		if playerTable.MapLen() == 0 {
			continue
		}

		players = append(players, Player{
			Name:    playerTable.GetString("name"),
			Prefab:  playerTable.GetString("prefab"),
			SteamId: playerTable.GetString("netid"),
			Colour:  playerTable.GetString("colour"),
			Level:   playerTable.GetInt("eventlevel"),
		})
	}

	return players, nil
}

// parsed mods info from string slice, examples as follows:
//  [
//  	"workshop-374550642",
//  	"Increased Stack size",
//  	"1.62",
//  	"1.62",
//  	true,
//  	"workshop-2798599672",
//  	"六格装备栏（适配mod版）",
//  	"4.6.8.f",
//  	"4.6.8.f",
//  	true,
//  	"workshop-378160973",
//  	"Global Positions",
//  	"1.7.4",
//  	"1.7.4",
//  	true,
//  ]

func parseModsInfo(mods []any) []Mod {
	if len(mods) == 0 {
		return nil
	}

	var res []Mod

	for i := 0; i < len(mods); i++ {
		if strings.Contains(mods[i].(string), "workshop-") {
			modId := strings.Split(mods[i].(string), "-")[1]
			name := mods[i+1].(string)
			v1 := mods[i+2].(string)
			v2 := mods[i+3].(string)
			enable := mods[i+4].(bool)

			i = i + 4

			res = append(res, Mod{
				Id:       modId,
				Name:     name,
				Version1: v1,
				Version2: v2,
				Enabled:  enable,
			})
		}
	}

	return res
}

func PlatformDisplayName(region string, platform Platform) string {
	// WeGame only supported in CN
	if region == ApEast && platform == Rail {
		return "WeGame"
	}
	return platform.String()
}
