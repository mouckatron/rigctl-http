build: build_ui embed_ui build_go

build_ui:
	cd ui && ng build --base-href /ui/

embed_ui:
	scripts/embed_ui.sh

build_go:
	cd cmd/rigctl-http && go build

integration: build
	rigctld &
	cmd/rigctl-http/rigctl-http -gin-debug &
	cd test/integration && go test || true
	-killall rigctld
	-killall rigctl-http

clean:
	find . -type f -name '*~' | xargs rm

openapi:
	ng-openapi-gen --input api/reference/rigctl-http.yaml --output ui/src/app/api
