
export GOPATH=$(PWD)/thirds


APPS=datanote2

all: $(APPS)

datanote2:
	go build -o $@  src/main.go

check: test_log
test_log:
	cd $(PWD)/src/common/log && go test
clean:
	rm -fr $(APPS)

