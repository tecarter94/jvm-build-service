# Build the manager binary
FROM registry.access.redhat.com/ubi9/go-toolset:1.21.9-1.1716478616@sha256:9b4a17e28f4a2836c3e443a769d1d51de056b9a1628a974fde644dd8fa5dffab as builder

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY vendor/ vendor/
COPY pkg/ pkg/
COPY cmd/ cmd/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jvmbuildservice cmd/controller/main.go

# Use ubi-minimal as minimal base image to package the manager binary
# Refer to https://catalog.redhat.com/software/containers/ubi8/ubi-minimal/5c359a62bed8bd75a2c3fba8 for more details
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4-949.1714662671@sha256:2636170dc55a0931d013014a72ae26c0c2521d4b61a28354b3e2e5369fa335a3
COPY --from=builder /opt/app-root/src/jvmbuildservice /
USER 65532:65532

ENTRYPOINT ["/jvmbuildservice"]
