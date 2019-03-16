FROM ubuntu:xenial

ENV PATH=/usr/lib/go-1.9/bin:$PATH

RUN /bin/bash -c '\
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.9 git make gcc libc-dev ca-certificates wget && \
  git clone https://github.com/metadium/go-metadium && \
  (cd go-metadium && make) && \
  mkdir -p /usr/local/metadium /usr/local/metadium/keystore && \
  cp -r go-metadium/build/bin go-metadium/build/conf /usr/local/metadium/ && \
  /usr/bin/wget -nv -O /usr/local/bin/solc https://github.com/ethereum/solidity/releases/download/v0.4.24/solc-static-linux && \
  chmod a+x /usr/local/bin/solc && \
  for i in 1 2 3 4; do \
    /usr/local/metadium/bin/gmet metadium new-account --password <(echo demo) --out /usr/local/metadium/keystore/account-$i; \
  done && \
  apt-get remove -y golang-1.9 git make gcc libc-dev wget && apt autoremove -y && apt-get clean && \
  rm -rf /go-metadium'

EXPOSE 10009
EXPOSE 10010

ENTRYPOINT ["/bin/bash"]
