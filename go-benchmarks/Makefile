bench:
	cd slices && go test -bench . -benchmem

cpuprofile:
	cd slices && go test -bench . -benchmem -cpuprofile=cpuprofile.pprof

run_cpupprof:
	cd slices && go tool pprof cpuprofile.pprof

memprofile:
	cd slices && go test -bench . -benchmem -memprofile=memprofile.pprof

run_mempprof:
	cd slices && go tool pprof memprofile.pprof