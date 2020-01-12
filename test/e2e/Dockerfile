# hadolint ignore=DL3007
FROM docker.pkg.github.com/jameshiew/kenv/kenv:latest AS kenv
RUN go get -v github.com/mikefarah/yq/v2

FROM ubuntu:18.04
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates=20180409 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd -r testuser && \
    useradd -m -s /bin/bash -r -g testuser testuser && \
    chown -R testuser:testuser /home/testuser
USER testuser
WORKDIR /home/testuser

COPY --from=kenv /go/bin/yq /usr/local/bin/yq
COPY --from=kenv /go/bin/kenv /usr/local/bin/kenv

COPY test.sh test.sh

ENTRYPOINT ["/bin/bash", "test.sh", "2>&1"]
