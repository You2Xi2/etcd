#!/bin/bash

sudo cgcreate -g cpu:/cpulimited

sudo cgset -r cpu.cfs_quota_us=50000 cpulimited 

sudo cgexec -g cpu:cpulimited ./bin/etcd
