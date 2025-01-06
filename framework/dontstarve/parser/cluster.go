package dstparser

import (
	"bytes"

	"gopkg.in/ini.v1"
)

// ClusterConfig represents cluster.ini
type ClusterConfig struct {
	GamePlay GamePlay       `ini:"GAMEPLAY"`
	NetWork  ClusterNetwork `ini:"NETWORK"`
	Misc     Misc           `ini:"MISC"`
	Shard    ClusterShard   `ini:"SHARD"`
	Steam    ClusterSteam   `ini:"STEAM"`
}

// ServerConfig represents server.ini
type ServerConfig struct {
	Network ServerNetwork `ini:"NETWORK"`
	Shard   ServerShard   `ini:"SHARD"`
	Steam   ServerSteam   `ini:"STEAM"`
	Account ServerAccount `ini:"ACCOUNT"`
}

type ClusterNetwork struct {
	// basic information
	WhiteListSlots     int    `ini:"whitelist_slots"`
	ClusterPassword    string `ini:"cluster_password"`
	ClusterName        string `ini:"cluster_name"`
	ClusterDescription string `ini:"cluster_description"`
	ClusterLanguage    string `ini:"cluster_language"`
	ClusterCloudId     string `ini:"cluster_cloud_id,omitempty"`

	// network
	Offline   bool `ini:"offline_cluster"`
	TrickRate int  `ini:"tick_rate,omitempty"`
	LanOnly   bool `ini:"lan_only_cluster"`
}

type GamePlay struct {
	MaxPlayers     int    `ini:"max_players"`
	Pvp            bool   `ini:"pvp"`
	GameMode       string `ini:"game_mode"`
	PauseWhenEmpty bool   `ini:"pause_when_empty"`
	VoteEnable     bool   `ini:"vote_enabled"`
	VoteKickEnable bool   `ini:"vote_kick_enabled"`
}

type Misc struct {
	MaxSnapShots  int  `ini:"max_snapshots"`
	ConsoleEnable bool `ini:"console_enabled"`
}

type ClusterShard struct {
	ShardEnable bool   `ini:"shard_enabled"`
	BindIP      string `ini:"bind_ip"`
	ClusterKey  string `ini:"cluster_key"`
	MasterIp    string `ini:"master_ip"`
	MasterPort  int    `ini:"master_port"`
}

type ClusterSteam struct {
	GroupOnly  bool   `ini:"steam_group_only"`
	GroupId    string `ini:"steam_group_id"`
	GroupAdmin bool   `ini:"steam_group_admins"`
}

type ServerNetwork struct {
	ServerPort int `ini:"server_port"`
}

type ServerSteam struct {
	MasterServerPort   int `ini:"master_server_port"`
	AuthenticationPort int `ini:"authentication_port"`
}

type ServerAccount struct {
	EncodeUserPath bool `ini:"encode_user_path"`
}

type ServerShard struct {
	ID       string `ini:"id,omitempty"`
	Name     string `ini:"name,omitempty"`
	IsMaster bool   `ini:"is_master"`
}

// ParseClusterInI parses the cluster.ini to ClusterConfig struct type
func ParseClusterInI(data []byte) (ClusterConfig, error) {
	var config ClusterConfig
	err := ini.MapTo(&config, data)
	if err != nil {
		return ClusterConfig{}, err
	}
	return config, nil
}

// ParseServerInI parses the server.ini to ServerConfig struct type
func ParseServerInI(data []byte) (ServerConfig, error) {
	var config ServerConfig
	err := ini.MapTo(&config, data)
	if err != nil {
		return ServerConfig{}, err
	}
	return config, nil
}

// ToClusterInI converts ClusterConfig to cluster.ini
func ToClusterInI(config ClusterConfig) ([]byte, error) {
	empty := ini.Empty()
	err := empty.ReflectFrom(&config)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(nil)
	if _, err := empty.WriteToIndent(buffer, "\t"); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// ToServerInI converts ServerConfig to server.ini
func ToServerInI(config ServerConfig) ([]byte, error) {
	empty := ini.Empty()
	err := empty.ReflectFrom(&config)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(nil)
	if _, err := empty.WriteToIndent(buffer, "\t"); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
