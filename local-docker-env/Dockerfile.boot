# Stage 1: Build stage
FROM golang:1.19 as builder

# Set environment variables
ENV PATH=/usr/local/go/bin:$PATH

# Update and upgrade the package list
RUN apt-get update && \
    apt-get upgrade -q -y

# Install required packages
RUN apt-get install -y --no-install-recommends \
    ca-certificates \
    make && \
    rm -rf /var/lib/apt/lists/*

# Define the location for custom certificates
ARG cert_location=/usr/local/share/ca-certificates

# Fetch and install certificates for proxy.golang.org
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null | \
    openssl x509 -outform PEM > ${cert_location}/proxy.golang.crt && \
    update-ca-certificates

# Set work directory
WORKDIR /gwemix

# Copy the source code
COPY . /gwemix

# Build wemix
RUN USE_ROCKSDB=NO make -f Makefile-local-env

# Clean up unnecessary packages and files after building
RUN apt-get remove -y \
    ca-certificates \
    make && \
    apt autoremove -y && \
    apt-get clean

# Stage 2: Runtime stage
FROM ubuntu:latest

# Update and upgrade the package list
RUN apt-get update && \
    apt-get upgrade -q -y

# Install required runtime packages
RUN apt-get install -y --no-install-recommends \
    g++ \
    libc-dev \
    bash \
    jq \
    wget \
    netcat-traditional && \
    rm -rf /var/lib/apt/lists/*

# Create directories for wemix
RUN mkdir -p /usr/local/wemix

# Set environment variables
ENV PATH=/usr/local/wemix/bin:$PATH

# Set work directory
WORKDIR /usr/local/wemix

# Copy the built binaries and configuration files from the builder stage
COPY --from=builder /gwemix/build /usr/local/wemix

# Copy config.json & key file
COPY local-docker-env/config.json ./conf/config.json
COPY local-docker-env/keystore/ ./keystore/
COPY local-docker-env/nodekey/ ./nodekey/

# Define variables to be used at build time
ARG NODE_NUM
ENV NODE_NUM=${NODE_NUM}

# Run set-nodekey.sh
RUN chmod a+x bin/set-nodekey.sh && \
    ./bin/set-nodekey.sh -a ${NODE_NUM}

# Clean up unnecessary files
RUN rm -rf nodekey

# Run init-boot.sh
RUN chmod a+x bin/init-boot.sh
CMD ./bin/init-boot.sh && tail -f /dev/null

# Clean up unnecessary packages
RUN apt-get remove -y \
    g++ \
    libc-dev && \
    apt autoremove -y && \
    apt-get clean

# Expose necessary ports
EXPOSE 8588 8589 8598

# # Define the entrypoint
# ENTRYPOINT ["bash"]