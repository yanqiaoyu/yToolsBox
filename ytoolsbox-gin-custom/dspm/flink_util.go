package dspm

import (
	"fmt"
	"github.com/buger/jsonparser"
	"log"
	"main/cons"
	"time"
)

var maxRetryTimes = 3
var DetectJobName = "RARE-detect"

func CancelAllJobs(url string) {
	flinkUrl := fmt.Sprintf("%s/jobs/overview", url)
	var body, statusCode, err = HttpGet(flinkUrl)
	if err != nil || statusCode != 200 {
		log.Printf("flinkCancelAllDetectJob err: %v\n", err)
		return
	}

	_, err = jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			log.Printf("err: %v\n", err)
			return
		}

		jid, err := jsonparser.GetString(value, "jid")
		if err != nil {
			log.Printf("parse json jid: %v\n", err)
			return
		}
		state, err := jsonparser.GetString(value, "state")
		if err != nil {
			log.Printf("parse json state: %v", err)
			return
		}

		if state != "CANCELED" {
			flinkJobIdCancel(url, jid)
		}
	}, "jobs")

}

func flinkJobIdCancel(url string, jobId string) error {

	flinkUrl := fmt.Sprintf("%s/jobs/%s?mode=cancel", url, jobId)

	log.Printf("flinkJobCancel cancel jobId: %s, url: %s\n", jobId, flinkUrl)
	var err error
	var statusCode int

	for i := 0; i < maxRetryTimes; i++ {
		_, statusCode, err = HttpPatch(flinkUrl)
		if err == nil {
			break
		}
	}
	if err != nil || statusCode != 200 {
		return fmt.Errorf("cancel job:%s failed: [%v %d]", flinkUrl, err, statusCode)
	}

	return nil
}

func GetSuccessFlinkTask(url string) map[string]interface{} {
	taskMap := make(map[string]interface{})
	flinkUrl := fmt.Sprintf("%s/jobs/overview", url)
	var body, statusCode, err = HttpGet(flinkUrl)
	log.Printf("flinkUrl:%s\n", flinkUrl)
	if err != nil || statusCode != 200 {
		log.Printf("GetSuccessFlinkTask err: %v\n", err)
		return taskMap
	}

	_, err = jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			log.Printf("err: %v\n", err)
			return
		}

		name, err := jsonparser.GetString(value, "name")
		if err != nil {
			log.Printf("parse json name: %v\n", err)
			return
		}

		state, err := jsonparser.GetString(value, "state")
		if err != nil {
			log.Printf("parse json state: %v\n", err)
			return
		}

		if name == DetectJobName && state == "RUNNING" {
			taskMap[DetectJobName] = true
		}
	}, "jobs")
	return taskMap
}

func RestartFlink() {
	flinkUrl := fmt.Sprintf("http://%s:%d", cons.DspmAddr, cons.FlinkPort)
	CancelAllJobs(flinkUrl)

	resetNum := 0
	for true {
		taskMap := GetSuccessFlinkTask(flinkUrl)
		if len(taskMap) == 1 {
			break
		}
		resetNum++
		log.Printf("reset:%d; 当前任务:%s", resetNum, taskMap)
		time.Sleep(10 * time.Second)
	}
	log.Printf("restart flink task success")
}
