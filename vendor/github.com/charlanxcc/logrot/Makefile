# Makefile

all: logrot

logrot: bin/logrot
bin/logrot: bin src logrot.go cmd/Main.go
	GOPATH=${PWD} go build -o $@ cmd/Main.go


bin:
	@[ -d $@ ] || mkdir -p $@

clean:
	rm -rf bin
	rm -f src logrot

src:
	ln -sf . src
	ln -sf . logrot

# EOF
