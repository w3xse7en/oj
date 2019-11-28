package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	didaTaskFile = "dida.task.txt"
)

func dida(prefix string, p problem) {
	// 不再往滴答清单中添加任务
	// task := p.didaTask(prefix)
	// mailToDida(task)
}

func saveLocal(task string) {
	ts, err := ioutil.ReadFile(didaTaskFile)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Panicf("无法读取 %s：%s\n", didaTaskFile, err)
		}
		f, _ := os.Create(didaTaskFile)
		f.Close()
	}

	ts = append(ts, []byte(task+"\n")...)

	err = ioutil.WriteFile(didaTaskFile, ts, 0755)
	if err != nil {
		log.Panicf("无法写入 %s: %s\n", didaTaskFile, err)
	}

	log.Printf("新建任务已经写入 %s，请手动添加到滴答清单", didaTaskFile)
}

var m = map[string]time.Duration{
	"#do": 15,
	"#re": 90,
	"#fa": 30,
}

func delay(task string) string {
	key := task[:3]
	if day, ok := m[key]; ok {
		task += time.Now().Add(time.Hour * 24 * day).Format("2006-01-02")
		m[key] += 2
	}
	return task
}
