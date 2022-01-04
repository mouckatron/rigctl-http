
build:
	cd cmd/rigctl-http && go build

integration: build
	rigctld &
	cmd/rigctl-http/rigctl-http &
	cd test/integration && go test || true
	-killall rigctld
	-killall rigctl-http
