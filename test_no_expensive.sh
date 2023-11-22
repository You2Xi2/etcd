#!/bin/bash

# Specify the etcd benchmark parameters
NUM_EXPENSIVE_READS=1
NUM_WRITES=20
TOTAL_WRITES=$((NUM_WRITES * 10))
NUM_CHEAP_READS=100
TOTAL_CHEAP_READS=$((NUM_CHEAP_READS * 10))

NUM_KEYS=100000
KEY_SIZE=256
VALUE_SIZE=1024

ENDPOINT="http://localhost:2379"

n=$1
# Please config and make dir before test
OUTPUT_DIR=./output/without_expensive_${n}
mkdir ${OUTPUT_DIR}
# CMD=./bin/benchmark_autocancel
CMD=./bin/benchmark

export ETCDCTL_API=3

# Perform the read benchmark in the background and direct output to a file
# echo "Starting the expensive read benchmark..."
# ${CMD} range 1 2 \
#     --endpoints=$ENDPOINT \
#     --rate=1 \
#     --total=11 \
#     > $OUTPUT_DIR/expensive_read.log 2>&1 &

# Perform the read benchmark in the background and direct output to a file
echo "Starting the cheap read benchmark..."
${CMD} range 14980 14981\
    --endpoints=$ENDPOINT \
    --rate=$NUM_CHEAP_READS \
    --total=$TOTAL_CHEAP_READS \
    --clients=64 \
    > $OUTPUT_DIR/cheap_read.log 2>&1 &

# Perform the write benchmark in the background and direct output to a file
echo "Starting the cheap write benchmark..."
${CMD} put \
    --endpoints=$ENDPOINT \
    --rate=$NUM_WRITES \
    --total=$TOTAL_WRITES \
    --clients=64 \
    > $OUTPUT_DIR/cheap_write.log 2>&1 &
