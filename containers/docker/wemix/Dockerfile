FROM ubuntu:xenial

ENV PATH=/usr/lib/go-1.9/bin:$PATH

RUN /bin/bash -c '\
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.9 git make gcc libc-dev ca-certificates wget && \
  git clone https://github.com/wemix3/go-wemix && \
  (cd go-wemix && make) && \
  mkdir -p /usr/local/wemix /usr/local/wemix/keystore && \
  cp -r go-wemix/build/bin go-wemix/build/conf /usr/local/wemix/ && \
  /usr/bin/wget -nv -O /usr/local/bin/solc https://github.com/ethereum/solidity/releases/download/v0.4.24/solc-static-linux && \
  chmod a+x /usr/local/bin/solc && \
  for i in 1 2 3 4; do \
    /usr/local/wemix/bin/gwemix wemix new-account --password <(echo demo) --out /usr/local/wemix/keystore/account-$i; \
  done && \
  apt-get remove -y golang-1.9 git make gcc libc-dev wget && apt autoremove -y && apt-get clean && \
  rm -rf /go-wemix'

EXPOSE 10009
EXPOSE 10010

ENTRYPOINT ["/bin/bash"]
