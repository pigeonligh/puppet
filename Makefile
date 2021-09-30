BIN_DIR=_output/bin

all: puppet control

init:
	mkdir -p ${BIN_DIR}

puppet: init
	go build -v -o=${BIN_DIR}/puppet ./cmd/puppet/

control: init
	go build -v -o=${BIN_DIR}/control ./cmd/control/

clean:
	rm -rf _output/
