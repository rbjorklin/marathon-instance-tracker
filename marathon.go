package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/* This feels dirty... */
func getIds(appId string) []string {
	var ids []string
	var base map[string]interface{}
	var apps map[string]interface{}
	var tasks []interface{}
	var taskInfo map[string]interface{}
	r, err := http.Get(os.Getenv("MARATHON_URL") + "/v2/apps" + appId)
	check(err)
	body, err := ioutil.ReadAll(r.Body)
	check(err)
	check(r.Body.Close())
	if err := json.Unmarshal(body, &base); err != nil {
		panic(err)
	}
	apps = base["app"].(map[string]interface{})
	tasks = apps["tasks"].([]interface{})
	for _, task := range tasks {
		taskInfo = task.(map[string]interface{})
		id := taskInfo["id"].(string)
		ids = append(ids, id)
	}
	return ids
}

func eventBusRegister() {
	var tmp io.Reader
	url := os.Getenv("MARATHON_URL") + "/v2/eventSubscriptions?callbackUrl=http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/events"
	r, err := http.Post(url, "", tmp)
	check(err)
	r.Body.Close()
}

func getInstances(appId string) int {
	var base map[string]interface{}
	var apps map[string]interface{}
	r, err := http.Get(os.Getenv("MARATHON_URL") + "/v2/apps" + appId)
	check(err)
	body, err := ioutil.ReadAll(r.Body)
	check(err)
	check(r.Body.Close())
	if err := json.Unmarshal(body, &base); err != nil {
		panic(err)
	}
	apps = base["app"].(map[string]interface{})
	return int(apps["instances"].(float64))
}
