#!/bin/bash

# Specify the etcd benchmark parameters
NUM_EXPENSIVE_READS=1
NUM_WRITES=20
NUM_CHEAP_READS=100

NUM_KEYS=100000
KEY_SIZE=256
VALUE_SIZE=1024

ENDPOINT="http://localhost:2379"

suffix=$1
# Please config and make dir before test
OUTPUT_DIR=./output/with_expensive_${suffix}
mkdir ${OUTPUT_DIR}
CMD=./bin/benchmark
number_of_times=10

# Run the commands in a loop
for ((i=1; i<=${number_of_times}; i++)); do
  # Expensive read benchmark
  echo "Starting the expensive read benchmark (iteration $i)..."
  ${CMD} range 1 2 \
      --endpoints=$ENDPOINT \
      --rate=1 \
      --total=1 \
      > "$OUTPUT_DIR/expensive_read_${i}.log" 2>&1 &

  # Cheap read benchmark
  echo "Starting the cheap read benchmark (iteration $i)..."
  ${CMD} range 14980 14981 \
      --endpoints=$ENDPOINT \
      --rate=$NUM_CHEAP_READS \
      --total=$NUM_CHEAP_READS \
      --clients=64 \
      > "$OUTPUT_DIR/cheap_read_${i}.log" 2>&1 &

  # Cheap write benchmark
  echo "Starting the cheap write benchmark (iteration $i)..."
  ${CMD} put \
      --endpoints=$ENDPOINT \
      --rate=$NUM_WRITES \
      --total=$NUM_WRITES \
      --clients=64 \
      > "$OUTPUT_DIR/cheap_write_${i}.log" 2>&1 &

  # Wait for 1 second before the next iteration
  sleep 1

  # If i is even, wait for 10 seconds
  if [ $((i % 2)) -eq 0 ]; then
    echo "Waiting for 5 seconds before the next iteration..."
    sleep 5
  else
    # Otherwise, wait for 1 second before the next iteration
    echo "Waiting for 1 second before the next iteration..."
    sleep 1
  fi
done

echo "Script completed."
