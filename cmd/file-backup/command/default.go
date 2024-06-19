package command

import (
	"time"

	"github.com/pefish/file-backup/pkg/global"

	"github.com/pefish/file-backup/pkg/backup"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	go_time "github.com/pefish/go-time"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) Config() interface{} {
	return &global.GlobalConfig
}

func (dc *DefaultCommand) Data() interface{} {
	return &global.GlobalData
}

func (dc *DefaultCommand) Init(command *commander.Commander) error {
	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}

	return nil
}

func (dc *DefaultCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *DefaultCommand) Start(command *commander.Commander) error {
	fromDir := command.Args["from-dir"]
	toDir := command.Args["to-dir"]

	timer := time.NewTimer(0)
	for {
		select {
		case <-timer.C:
			if float64(go_time.TimeInstance.CurrentTimestamp()-global.GlobalData.LastBackupTimestamp) < global.GlobalConfig.IntervalHours*3600000 {
				timer.Reset(10 * time.Second)
				continue
			}
			err := backup.NewBackupTool().Backup(fromDir, toDir)
			if err != nil {
				return err
			}

			global.GlobalData.LastBackupTimestamp = go_time.TimeInstance.CurrentTimestamp()
			timer.Reset(10 * time.Second)
		case <-command.Ctx.Done():
			return nil
		}
	}
	return nil
}
