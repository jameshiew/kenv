#!/usr/bin/env bash
#
# These tests should be run in a clean Ubuntu 18.04 environment

set -ex

kenv version | grep system

kenv init

[ -d "${HOME}"/.kube/ ]
[ -d "${HOME}"/.kube/kenv/ ]
[ -d "${HOME}"/.kube/kenv/shims/ ]
test -x ~/.kube/kenv/shims/kubectl
[ -d "${HOME}"/.kube/kenv/versions/ ]

eval "$(kenv init -)"

command -v kubectl
kubectl | grep "No version of kubectl is installed"

kenv version | grep system

kenv install 1.17.0

[ -d "${HOME}"/.kube/kenv/versions/1.17.0/ ]
test -x ~/.kube/kenv/versions/1.17.0/kubectl

kenv global 1.17.0
kenv global | grep 1.17.0
kenv version | grep 1.17.0
kubectl version --client -oyaml | yq read - clientVersion.gitVersion | grep 1.17.0

kenv local
kenv install 1.14.0
kenv local 1.14.0
kenv local | grep 1.14.0
kenv version | grep 1.14.0
kubectl version --client -oyaml | yq read - clientVersion.gitVersion | grep 1.14.0

kenv shell
kenv install 1.15.0
kenv shell 1.15.0
kenv shell | grep 1.15.0
kenv version | grep 1.15.0
kubectl version --client -oyaml | yq read - clientVersion.gitVersion | grep 1.15.0

echo "Tests passed!"
