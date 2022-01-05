
build:
	cd cmd/rigctl-http && go build

integration: build
	rigctld &
	cmd/rigctl-http/rigctl-http -gin-debug &
	cd test/integration && go test || true
	-killall rigctld
	-killall rigctl-http

clean:
	find . -type f -name '*~' | xargs rm
