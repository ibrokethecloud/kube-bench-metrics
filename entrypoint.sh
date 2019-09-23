#!/bin/sh -x

## Eventually might add logic here for auto-detection #
## Cannot depent on auto-detection from kube-bench because
## it runs the master and node checks, and the json output
## gets overwritten.
## As a result these checks need to be run in a loop.

apiServer=$(ps -C api-server -o cmd --no-header )
etcdServer=$(ps -C etcd -o cmd --no-header )

if [ ! -z "$apiServer"  ] || [ ! -z "$etcdServer" ]
then
    export nodeType="master"
else 
    export nodeType="node"
fi
 
exec kube-bench-metrics "$@" --nodeType ${nodeType}