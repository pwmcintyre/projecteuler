
t=`date +%s`
mkdir -p benchmarks

go test -v \
    -count=5 \
    -run=. \
    -bench=. \
    ./... | tee -a benchmarks/test-$t.txt

benchstat benchmarks/*
