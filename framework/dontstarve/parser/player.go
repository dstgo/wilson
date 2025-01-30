package dstparser

import (
	"strconv"
	"strings"

	"github.com/dstgo/wilson/framework/pkg/strs"
)

// ParsePlayerTxt parse player.txt file, returns slice of klei ID
func ParsePlayerTxt(content string) []string {
	if len(content) == 0 {
		return nil
	}
	return strings.Split(content, "\n")
}

// ToPlayerTxt converts kleiIDs to player.txt
func ToPlayerTxt(kleiIDs []string) ([]byte, error) {
	joinStr := strings.Join(kleiIDs, "\n")
	return strs.StringToBytes(joinStr), nil
}

// ChatLog represents a chat record in server_chat_log.txt
type ChatLog struct {
	Time   string `mapstructure:"time"`
	Type   string `mapstructure:"type"`
	KleiId string `mapstructure:"klei_id"`
	Name   string `mapstructure:"player_name"`
	Msg    string `mapstructure:"msg"`
}

func trim(s, l, r string) string {
	s = strings.TrimLeft(s, l)
	s = strings.TrimRight(s, r)
	return s
}

func surround(bs string, l, r string) bool {
	return strings.HasPrefix(bs, l) && strings.HasSuffix(bs, r)
}

func ParseServerChatLogs(content []byte) ([]ChatLog, error) {
	if len(content) == 0 {
		return nil, nil
	}

	logsStr := strs.BytesToString(content)

	var logs []ChatLog
	lines := strings.Split(logsStr, "\n")
	for _, line := range lines {
		var log ChatLog
		records := strings.Split(line, ": ")
		if len(records) > 0 {
			log.Time = trim(records[0], "[", "]")
		}

		if len(records) > 1 {
			record := strings.Join(records[1:], ": ")

			sql := strings.Index(record, "[")
			sqr := strings.Index(record, "]")

			// log type
			if sqr > sql {
				log.Type = strings.TrimSpace(record[sql+1 : sqr])
			}

			if log.Type == "Vote Announcement" {
				log.Msg = strings.TrimSpace(record[sqr+1:])
				logs = append(logs, log)
				continue
			}

			bkl := strings.Index(record, "(")
			bkr := strings.Index(record, ")")

			// klei id
			if bkr > bkl {
				kid := strings.TrimSpace(record[bkl+1 : bkr])
				if strings.HasPrefix(kid, "KU_") {
					log.KleiId = kid
				}
			}

			idx := sqr
			if bkr > 0 && len(log.KleiId) > 0 {
				idx = bkr
			}

			// name
			fields := strings.Split(strings.TrimSpace(record[idx+1:]), " ")
			if len(fields) > 0 {
				log.Name = strings.TrimRight(fields[0], ":")
			}

			// msg
			if len(fields) > 1 {
				log.Msg = strconv.Quote(strings.TrimSpace(strings.Join(fields[1:], " ")))
			}
		}

		logs = append(logs, log)
	}

	return logs, nil
}
