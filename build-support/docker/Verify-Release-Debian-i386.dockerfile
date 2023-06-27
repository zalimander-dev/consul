# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

FROM i386/debian:bullseye
ARG TARGETARCH=i386
RUN apt-add-repository "deb [arch=${TARGETARCH}] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
RUN apt update && apt install -y software-properties-common curl gnupg
RUN curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add -
ARG PACKAGE=consul \
ARG VERSION \
ARG SUFFIX=1
RUN apt-get update && apt-get install -y ${PACKAGE}=${VERSION}-${SUFFIX}
