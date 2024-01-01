package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
)

/*
Author: Bahtyar
Date: 2024-01-01 14:30:28
Description: 监测文件变动脚本,暂时不用
*/

type Tconf struct {
	Path string
}

func main() {
	f := "config.toml"
	if _, err := os.Stat(f); err != nil {
		log.Fatalln(0, err)
	}

	var config Tconf
	var _, err = toml.DecodeFile(f, &config)
	if err != nil {
		log.Fatalln(2, err)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	var lastWriteEventTime time.Time
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if time.Since(lastWriteEventTime).Seconds() < 1 {
					continue
				}
				lastWriteEventTime = time.Now()
				// 打印监听事件
				log.Println("event:", event)
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
			}
		}
	}()
	// 监听当前目录
	err = watcher.Add(config.Path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func push() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "an old falcon")
	}()

	// out_, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
