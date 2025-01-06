package lobbyapi

import (
	"bytes"
	"errors"
	"strings"
	"text/template"

	lua "github.com/yuin/gopher-lua"
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

	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(luaScript); err != nil {
		return err
	}

	table := L.Get(-1).(*lua.LTable)
	L.Pop(1)
	details.Details.Day = int(table.RawGetString("day").(lua.LNumber))
	details.Details.DayElapsedInSeason = int(table.RawGetString("dayselapsedinseason").(lua.LNumber))
	details.Details.DaysLeftInSeason = int(table.RawGetString("daysleftinseason").(lua.LNumber))

	return nil
}

// parse players info from lua script
func parsePlayersInfo(luaScript string) ([]Player, error) {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(luaScript); err != nil {
		return nil, err
	}

	var players []Player

	table := L.Get(-1).(*lua.LTable)
	table.ForEach(func(idx lua.LValue, value lua.LValue) {
		playerTable := value.(*lua.LTable)
		players = append(players, Player{
			Name:    playerTable.RawGetString("name").String(),
			Prefab:  playerTable.RawGetString("prefab").String(),
			SteamId: playerTable.RawGetString("netid").String(),
			Colour:  playerTable.RawGetString("colour").String(),
			Level:   int(playerTable.RawGetString("eventlevel").(lua.LNumber)),
		})
	})

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
