
t=`date +%s`
mkdir -p benchmarks
dir=${1:-./}

go test -v \
    -count=5 \
    -run=. \
    -bench=. \
    ${dir}/... | tee -a benchmarks/test-$t.txt

benchstat benchmarks/*
