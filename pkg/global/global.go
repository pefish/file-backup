package global

import "github.com/pefish/go-commander"

type Config struct {
	commander.BasicConfig
	IntervalHours float64 `json:"interval-hours" default:"24" usage:"Set backup interval."`
}

type Data struct {
	LastBackupTimestamp uint64 `json:"last_backup_timestamp"`
}

var GlobalConfig Config
var GlobalData Data
