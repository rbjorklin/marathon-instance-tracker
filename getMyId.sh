#!/bin/bash

idCounter=1
for taskId in $(zkcli --servers srv-1,srv-2,srv-3 get /instance-tracker/${MARATHON_APP_ID} | tr , " ") ; do
    if [[ "${MESOS_TASK_ID}" == "${taskId}" ]] ; then
        echo ${idCounter}
        break
    fi
    idCounter=$((idCounter + 1))
done
