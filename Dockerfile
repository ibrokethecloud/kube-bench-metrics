FROM golang:1.22 AS builder
RUN mkdir -p /src/github.com/ibrokethecloud/kube-bench-metrics
COPY . /src/github.com/ibrokethecloud/kube-bench-metrics
RUN cd /src/github.com/ibrokethecloud/kube-bench-metrics \
    && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o kube-bench-metrics .

## Using upstream aquasec kube-bench and layering it up
FROM aquasec/kube-bench:v0.7.2
COPY --from=builder /src/github.com/ibrokethecloud/kube-bench-metrics/kube-bench-metrics /usr/bin/kube-bench-metrics
COPY entrypoint.sh /entrypoint.sh
COPY helper_scripts/check_files_owner_in_dir.sh /usr/local/bin
COPY helper_scripts/check_files_permissions.sh /usr/local/bin

WORKDIR /opt/kube-bench
ENTRYPOINT [ "/entrypoint.sh" ] 
