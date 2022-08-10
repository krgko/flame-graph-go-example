#!/bin/bash

echo "Start Generating Flame Graph..."
./flamegraph/stackcollapse-go.pl cpu.txt | ./flamegraph/flamegraph.pl > flame.svg
echo "Done !"