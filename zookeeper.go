package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var zkBase string = "/instance-tracker"
var flags = int32(0)
var acl = zk.WorldACL(zk.PermAll)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *zk.Conn {
	zksStr := os.Getenv("ZK_SERVERS")
	zks := strings.Split(zksStr, ",")
	conn, _, err := zk.Connect(zks, time.Second)
	check(err)
	return conn
}

func checkPath(conn *zk.Conn, path string) {
	exists, _, err := conn.Exists(path)
	if exists {
		if path == "/" {
			path = os.Getenv("ZK_BASE")
			if path == "" {
				path = zkBase
			}
			_, err := conn.Create(path, []byte{}, flags, acl)
			check(err)
		}
		return
	}
	checkPath(conn, filepath.Dir(path))
	_, err = conn.Create(path, []byte{}, flags, acl)
	check(err)
	return
}

func updateId(appId, taskId string) {
	replaced := false
	newContents := ""
	zkPath := filepath.Join(zkBase, appId)
	conn := connect()
	defer conn.Close()
	checkPath(conn, zkPath)
	lock := zk.NewLock(conn, zkPath, acl)
	lock.Lock()
	defer lock.Unlock()
	contents, stat, err := conn.Get(zkPath)
	if len(contents) == 0 {
		_, err := conn.Set(zkPath, []byte(taskId), stat.Version)
		check(err)
		return
	}
	log.Print("Original contents: ", string(contents))
	sContents := removeInactiveIds(string(contents), appId)
	ids := strings.Split(sContents, ",")
	for i, id := range ids {
		if id == "NA" {
			ids[i] = taskId
			replaced = true
			break
		}
	}
	if replaced {
		newContents = strings.Join(ids, ",")
	} else {
		newContents = strings.Join([]string{sContents, taskId}, ",")
	}
	log.Print("Contents: ", sContents)
	log.Print("New Contents: ", newContents)
	stat, err = conn.Set(zkPath, []byte(newContents), stat.Version)
	check(err)
	return
}

func removeId(appId, taskId string) {
	zkPath := filepath.Join(zkBase, appId)
	conn := connect()
	defer conn.Close()
	exists, _, err := conn.Exists(zkPath)
	check(err)
	if !exists {
		return
	}
	lock := zk.NewLock(conn, zkPath, acl)
	lock.Lock()
	defer lock.Unlock()
	contents, stat, err := conn.Get(zkPath)
	if len(contents) == 0 {
		return
	}
	newContents := strings.Replace(string(contents), taskId, "NA", 1)
	_, err = conn.Set(zkPath, []byte(newContents), stat.Version)
	check(err)
	return
}

func removeInactiveIds(contents string, appId string) string {
	log.Print("Entered removeInactiveIds")
	inst := getInstances(appId)
	ids := getIds(appId)
	for _, id := range strings.Split(contents, ",") {
		log.Print("Checking id: ", id)
		if !sListContains(ids, id) {
			log.Print("Replacing id: ", id)
			contents = strings.Replace(contents, id, "NA", 1)
		}
	}
	if len(strings.Split(contents, ",")) > inst + 1 {
		contents = strings.Join(strings.Split(contents, ",")[:inst + 1], ",")
	}
	log.Print("Returned: ", contents)
	return contents
}

func sListContains(l []string, s string) bool {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}
