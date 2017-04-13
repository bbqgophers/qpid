MOCK_PACKAGES = \
	gobot \
	http \
	log \
	prometheus \
	twillio \
	vendor/github.com/prometheus/client_golang/api/prometheus \
	vendor/github.com/hybridgroup/gobot

build-mocks:
	for dir in $(MOCK_PACKAGES); do \
		mockery -dir=./$$dir/ -all -case underscore -note "*DO NOT EDIT* Auto-generated via mockery"; \
	done

deps:
	glide update
