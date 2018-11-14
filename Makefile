all: build

build:
	cd cmd/client && go fmt && go build
	cd tools && go fmt && go build -o crd-generator

GENERATED_MANIFEST_DIR=manifests/generated

generate:
	hack/update-codegen.sh
	mkdir -p manifests/generated
	tools/crd-generator -crd-type net-state > $(GENERATED_MANIFEST_DIR)/net-state-crd.yaml
	tools/crd-generator -crd-type net-conf > $(GENERATED_MANIFEST_DIR)/net-conf-crd.yaml

test:
	hack/test.sh

dep:
	dep ensure -v

clean-dep:
	rm -f ./Gopkg.lock
	rm -rf ./vendor

clean-generate:
	rm -f pkg/apis/nmstate.io/v1/zz_generated.deepcopy.go
	rm -rf pkg/client
	rm -rf $(GENERATED_MANIFEST_DIR)
