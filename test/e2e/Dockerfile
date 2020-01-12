FROM docker.pkg.github.com/jameshiew/kenv/kenv:latest
RUN go get -v github.com/mikefarah/yq/v2

FROM ubuntu:18.04
RUN apt-get update && apt-get install -y \
    ca-certificates \
    tree

RUN groupadd -r testuser && \
    useradd -m -s /bin/bash -r -g testuser testuser && \
    chown -R testuser:testuser /home/testuser
USER testuser
WORKDIR /home/testuser

COPY --from=0 /go/bin/yq /usr/local/bin/yq
COPY --from=0 /go/bin/kenv /usr/local/bin/kenv

COPY test/e2e/test.sh test.sh

ENTRYPOINT ["/bin/bash", "test.sh", "2>&1"]
