mocks:
	mockery -dir=./gobot -all
	mockery -dir=./http -all
	mockery -dir=./init -all
	mockery -dir=./log -all
	mockery -dir=./prometheus -all
	mockery -dir=./twillio -all
	mockery -dir=./vendor/github.com/prometheus/client_golang/api/prometheus -name Client
	mockery -dir=./vendor/github.com/hybridgroup/gobot -all


