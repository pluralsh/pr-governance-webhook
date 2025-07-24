FROM golang:1.24-alpine3.21 AS builder

ARG TARGETARCH

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY /cmd/ cmd/
COPY /pkg pkg/
COPY /api api/


# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} GO111MODULE=on go build -a -o pr-governance-webhook cmd/*.go


FROM alpine:3.21
WORKDIR /workspace

COPY --from=builder /workspace/pr-governance-webhook .

USER 65532:65532


ENTRYPOINT ["/workspace/pr-governance-webhook"]