# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: geth android ios evm all test clean rocksdb etcd
.PHONY: gmet-linux

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

# USE_ROCKSDB
# - undefined | "NO": Do not use
# - "YES": build a static lib from rocksdb directory, and use that one
# - "EXISTING": use existing rocksdb shared lib.
ifndef USE_ROCKSDB
  ifeq ($(shell uname), Linux)
    USE_ROCKSDB = YES
  else
    USE_ROCKSDB = NO
  endif
endif
ifneq ($(shell uname), Linux)
  USE_ROCKSDB = NO
endif

ifneq ($(USE_ROCKSDB), NO)
ROCKSDB_DIR=$(shell pwd)/rocksdb
ROCKSDB_TAG=-tags rocksdb
endif

metadium: etcd gmet logrot
	@[ -d build/conf ] || mkdir -p build/conf
	@cp -p metadium/scripts/gmet.sh metadium/scripts/solc.sh build/bin/
	@cp -p metadium/scripts/config.json.example		\
		metadium/scripts/genesis-template.json		\
		metadium/contracts/MetadiumGovernance.js	\
		metadium/scripts/deploy-governance.js		\
		build/conf/
	@(cd build; tar cfz metadium.tar.gz bin conf)
	@echo "Done building build/metadium.tar.gz"

gmet: etcd rocksdb metadium/governance_abi.go
ifeq ($(USE_ROCKSDB), NO)
	$(GORUN) build/ci.go install $(ROCKSDB_TAG) ./cmd/gmet
else
	CGO_CFLAGS=-I$(ROCKSDB_DIR)/include \
		CGO_LDFLAGS="-L$(ROCKSDB_DIR) -lrocksdb -lm -lstdc++ $(shell awk '/PLATFORM_LDFLAGS/ {sub("PLATFORM_LDFLAGS=", ""); print} /JEMALLOC=1/ {print "-ljemalloc"}' < $(ROCKSDB_DIR)/make_config.mk)" \
		$(GORUN) build/ci.go install $(ROCKSDB_TAG) ./cmd/gmet
endif
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gmet\" to launch gmet."

logrot:
	$(GORUN) build/ci.go install ./cmd/logrot

geth:
	$(GORUN) build/ci.go install ./cmd/geth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geth\" to launch geth."

dbbench: rocksdb
ifeq ($(USE_ROCKSDB), NO)
	$(GORUN) build/ci.go install $(ROCKSDB_TAG) ./cmd/dbbench
else
	CGO_CFLAGS=-I$(ROCKSDB_DIR)/include \
		CGO_LDFLAGS="-L$(ROCKSDB_DIR) -lrocksdb -lm -lstdc++ $(shell awk '/PLATFORM_LDFLAGS/ {sub("PLATFORM_LDFLAGS=", ""); print} /JEMALLOC=1/ {print "-ljemalloc"}' < $(ROCKSDB_DIR)/make_config.mk)" \
		$(GORUN) build/ci.go install $(ROCKSDB_TAG) ./cmd/dbbench
endif

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/geth.aar\" to use the library."
	@echo "Import \"$(GOBIN)/geth-sources.jar\" to add javadocs"
	@echo "For more info see https://stackoverflow.com/questions/20994336/android-studio-how-to-attach-javadoc"

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Geth.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: metadium/governance_abi.go ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/* build/conf metadium/admin_abi.go metadium/governance_abi.go
	@ROCKSDB_DIR=$(ROCKSDB_DIR);			\
	if [ -e $${ROCKSDB_DIR}/Makefile ]; then	\
		cd $${ROCKSDB_DIR};			\
		make clean;				\
	fi

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go install golang.org/x/tools/cmd/stringer@latest
	env GOBIN= go install github.com/kevinburke/go-bindata/go-bindata@latest
	env GOBIN= go install github.com/fjl/gencodec@latest
	env GOBIN= go install github.com/golang/protobuf/protoc-gen-go@latest
	env GOBIN= go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

gmet-linux: etcd
	@docker --version > /dev/null 2>&1;				\
	if [ ! $$? = 0 ]; then						\
		echo "Docker not found.";				\
	else								\
		docker build -t meta/builder:local			\
			-f Dockerfile.metadium . &&			\
		docker run -e HOME=/tmp --rm -v $(shell pwd):/data	\
			-w /data meta/builder:local			\
			make USE_ROCKSDB=$(USE_ROCKSDB);		\
	fi

ifneq ($(USE_ROCKSDB), YES)
rocksdb:
else
rocksdb:
	@[ ! -e rocksdb/.git ] && git submodule update --init rocksdb;	\
	cd $(ROCKSDB_DIR) && make -j8 static_lib;
endif

etcd:
	@if [ ! -e etcd/.git ]; then			\
		git submodule update --init etcd;	\
	fi

AWK_CODE='								\
BEGIN { print "package metadium"; bin = 0; name = ""; abi = ""; }	\
/^{/ { bin = 1; abi = ""; name = ""; }					\
/^}/ { bin = 0; abi = abi "}"; print "var " name "Abi = `" abi "`"; }	\
{									\
  if (bin == 1) {							\
    abi = abi $$0;							\
    if ($$1 == "\"contractName\":") {					\
      name = $$2;							\
      gsub(",|\"", "", name);						\
    }									\
  }									\
}'

metadium/admin_abi.go: metadium/contracts/MetadiumAdmin-template.sol build/bin/solc
	@PATH=${PATH}:build/bin metadium/scripts/solc.sh -f abi $< /tmp/junk.$$$$; \
	cat /tmp/junk.$$$$ | awk $(AWK_CODE) > $@;	\
	rm -f /tmp/junk.$$$$;

AWK_CODE_2='								     \
BEGIN { print "package metadium\n"; }					     \
/^var Registry_contract/ {						     \
  sub("^var[^(]*\\(","",$$0); sub("\\);$$","",$$0);			     \
  n = "Registry";							     \
  print "var " n "Abi = `{ \"contractName\": \"" n "\", \"abi\": " $$0 "}`"; \
}									     \
/^var Staking_contract/ {						     \
  sub("^var[^(]*\\(","",$$0); sub("\\);$$","",$$0);			     \
  n = "Staking";							     \
  print "var " n "Abi = `{ \"contractName\": \"" n "\", \"abi\": " $$0 "}`"; \
}									     \
/^var EnvStorageImp_contract/ {						     \
  sub("^var[^(]*\\(","",$$0); sub("\\);$$","",$$0);			     \
  n = "EnvStorageImp";							     \
  print "var " n "Abi = `{ \"contractName\": \"" n "\", \"abi\": " $$0 "}`"; \
}									     \
/^var Gov_contract/ {							     \
  sub("^var[^(]*\\(","",$$0); sub("\\);$$","",$$0);			     \
  n = "Gov";								     \
  print "var " n "Abi = `{ \"contractName\": \"" n "\", \"abi\": " $$0 "}`"; \
}'

metadium/governance_abi.go: metadium/contracts/MetadiumGovernance.js
	@cat $< | awk $(AWK_CODE_2) > $@

ifneq ($(shell uname), Linux)

build/bin/solc:
	@test 1

else

SOLC_URL=https://github.com/ethereum/solidity/releases/download/v0.4.24/solc-static-linux
build/bin/solc:
	@[ -d build/bin ] || mkdir -p build/bin;		\
	if [ ! -x build/bin/solc ]; then			\
		if which curl > /dev/null 2>&1; then		\
			curl -Ls -o build/bin/solc $(SOLC_URL);	\
			chmod +x build/bin/solc;		\
		elif which wget > /dev/null 2>&1; then		\
			wget -nv -o build/bin/solc $(SOLC_URL);	\
			chmod +x build/bin/solc;		\
		fi						\
	fi

endif
