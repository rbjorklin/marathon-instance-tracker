package main

type StatusUpdate struct {
	EventType  string `json:"eventType"`
	Timestamp  string `json:"timestamp"`
	SlaveId    string `json:"slaveId"`
	TaskId     string `json:"taskId"`
	TaskStatus string `json:"taskStatus"`
	AppId      string `json:"appId"`
	Host       string `json:"host"`
	Ports      []int  `json:"ports"`
	Version    string `json:"version"`
}
