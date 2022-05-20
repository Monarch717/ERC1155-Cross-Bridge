FROM golang:1.16.10-alpine3.14

# Set up apk dependencies
ENV PACKAGES make git libc-dev bash gcc linux-headers eudev-dev curl ca-certificates

# Set working directory for the build
WORKDIR /opt/app

# Add source files
COPY . .

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache $PACKAGES && \
    make build

# Run as non-root user for security
USER 1000

# Run the app
CMD ./build/swap-backend --config-type local --config-path config.json
