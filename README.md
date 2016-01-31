# Marathon instance tracker

**For your own safety don't use this in production without a ton of testing**

A quick and dirty solution to tracking Marathon instances.
It stores the ${MESOS_TASK_ID} in a comma (,) separated list under /instance-tracker/${MARATHON_APP_ID} in Zookeeper

## Get started

Build: go build -o instance-tracker *.go

Requires the following environment variables:

MARATHON_URL=http://yourmarathonurl.com:8080
PORT=9090
HOST=host_that_hosts_this_app
ZK_SERVERS=127.0.0.1:2181

$HOST, $MARATHON_URL and ZK_SERVERS need to be resolvable from Marathon as well as the server running this application.

Use something like [go-zkCli](https://github.com/go-zkcli/zkcli) to get the value from Zookeeper. See the included getMyId.sh for an example.

## TODO

* Sort out a working docker container
* Make sure it handles situations where znodes are locked
* Clean up the messy code
* Write some tests

## Credit where credit is due

* (This article)[http://thenewstack.io/make-a-restful-json-api-go/] by Cory Lanou helped me get started with creating a REST service.
* (This article)[https://mmcgrana.github.io/2014/05/getting-started-with-zookeeper-and-go.html] by Mark McGranaghan helped me get started with using Zookeper from Go.
