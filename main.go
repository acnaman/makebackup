package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

func main() {
	curdir := path.Dir(os.Args[0])
	fmt.Println(curdir)

	rootDir := getRootDir()
	newBackDir := getNewBackupPath(rootDir)
	os.MkdirAll(newBackDir, 0777)
	makeBackup(newBackDir)
}

func getRootDir() string {
	return "backup"
}

func getNewBackupPath(rootdir string) string {
	var newDir = rootdir + "/" + getTodayString()
	newDirSeed := newDir
	for i := 0; ; i++ {
		_, err := os.Stat(newDir)
		fmt.Println(newDir)
		// 存在しないファイルパスを見つけるまで続ける
		if err != nil {
			break
		} else {
			newDir = fmt.Sprintf("%s_%03d", newDirSeed, i)
		}
	}
	return newDir
}

func getTodayString() string {
	t := time.Now()
	return fmt.Sprintf("%4d%02d%02d", t.Year(), t.Month(), t.Day())
}

func makeBackup(backupdir string) {
	strcommand := "xcopy runtime " + backupdir
	println(strcommand)
	cmd := exec.Command(strcommand)
	cmd.Start()
}
