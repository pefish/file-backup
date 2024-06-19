package backup

import (
	"os/exec"

	go_shell "github.com/pefish/go-shell"
)

type BackupToolType struct {
}

func NewBackupTool() *BackupToolType {
	return &BackupToolType{}
}

func (b *BackupToolType) Backup(fromDir, toDir string) error {
	cmd := exec.Command("/bin/bash", "-c", `
#!/bin/bash
set -euxo pipefail

# 设置源文件夹和目标文件夹
SRC_FOLDER="`+fromDir+`"
DEST_FOLDER="`+toDir+`"

# 获取当前时间，并格式化为"年月日时分秒"
CURRENT_TIME=$(date +"%Y%m%d%H%M%S")

# 生成备份文件的名称
BACKUP_FILENAME="backup_${CURRENT_TIME}.tar.gz"
BACKUP_FILEPATH="${DEST_FOLDER}/${BACKUP_FILENAME}"

# 确保目标文件夹存在
if [ ! -d "$DEST_FOLDER" ]; then
    mkdir -p "$DEST_FOLDER"
fi

# 创建压缩文件
tar -czf "$BACKUP_FILEPATH" -C "$SRC_FOLDER" .

echo "Backup created at $BACKUP_FILEPATH"
	`)
	err := go_shell.ExecInConsole(cmd)
	if err != nil {
		return err
	}

	return nil
}
