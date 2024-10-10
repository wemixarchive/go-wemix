# Stage 1: Build stage
FROM golang:1.19 as builder

# Set environment variables
ENV PATH=/usr/local/go/bin:$PATH

# Update and upgrade the package list
RUN apt-get update && \
    apt-get upgrade -q -y

# Install required packages
RUN apt-get install -y --no-install-recommends \
    git \
    ca-certificates \
    openssl \
    make && \
    rm -rf /var/lib/apt/lists/*

# Define the location for custom certificates
ARG cert_location=/usr/local/share/ca-certificates

# Fetch and install certificates for github.com and proxy.golang.org
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null | \
    openssl x509 -outform PEM > ${cert_location}/github.crt && \
    update-ca-certificates
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null | \
    openssl x509 -outform PEM > ${cert_location}/proxy.golang.crt && \
    update-ca-certificates

# Clone the repository, install dependencies, and build the project
RUN git clone https://github.com/wemixarchive/go-wemix.git /go-wemix && \
    cd /go-wemix && \
    go mod download && \
    make

# Clean up unnecessary packages and files after building
RUN apt-get remove -y \
    git \
    ca-certificates \
    openssl \
    make && \
    apt autoremove -y && \
    apt-get clean

# Stage 2: Runtime stage
FROM ubuntu:22.04

# Set environment variables
ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

# Update and upgrade the package list
RUN apt-get update && \
    apt-get upgrade -q -y

# Install required runtime packages
RUN apt-get install -y --no-install-recommends \
    g++ \
    libc-dev \
    ca-certificates \
    bash \
    wget && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Create directories for wemix
RUN mkdir -p /usr/local/wemix /usr/local/wemix/keystore

# Set environment variables
ENV PATH=/usr/local/wemix/bin:$PATH

# Copy the built binaries and configuration files from the builder stage
COPY --from=builder /go-wemix/build /usr/local/wemix/

# Download and install solc
RUN wget -nv -O /usr/local/bin/solc https://github.com/ethereum/solidity/releases/download/v0.4.24/solc-static-linux && \
    chmod a+x /usr/local/bin/solc

# Create new accounts for wemix
RUN bash -c 'for i in 1 2 3 4; do \
    /usr/local/wemix/bin/gwemix wemix new-account --password <(echo demo) --out /usr/local/wemix/keystore/account-$i; \
done'

# Clean up unnecessary packages
RUN apt-get remove -y \
    g++ \
    libc-dev \
    wget && \
    apt autoremove -y && \
    apt-get clean

# Expose necessary ports
EXPOSE 8588 8589 8598

# Set the entrypoint
ENTRYPOINT ["/usr/local/wemix/bin/gwemix"]