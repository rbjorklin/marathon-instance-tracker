package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func EventRecv(w http.ResponseWriter, r *http.Request) {
	var statusUpdate StatusUpdate
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &statusUpdate); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	log.Print(statusUpdate)
	if statusUpdate.EventType == "status_update_event" {
		switch taskStatus := statusUpdate.TaskStatus; taskStatus {
		case "TASK_RUNNING":
			// 2 needs to to be changed to instances from GET /v2/apps/{AppId}
			updateId(statusUpdate.AppId, statusUpdate.TaskId)
		case "TASK_FINISHED":
			removeId(statusUpdate.AppId, statusUpdate.TaskId)
		case "TASK_FAILED":
			removeId(statusUpdate.AppId, statusUpdate.TaskId)
		case "TASK_KILLED":
			removeId(statusUpdate.AppId, statusUpdate.TaskId)
		case "TASK_LOST":
			removeId(statusUpdate.AppId, statusUpdate.TaskId)
		}
	}
}
