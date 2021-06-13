ARG STEP_1_IMAGE=alpine:3.13

FROM ${STEP_1_IMAGE} AS STEP_1

# Install Jupyter and gophernotes.
RUN set -x \
 && apk update \
 && apk --no-cache \
     --repository http://dl-4.alpinelinux.org/alpine/v3.13/community \
     --repository http://dl-4.alpinelinux.org/alpine/v3.13/main \
     --arch=x86_64 add \
     ca-certificates \
     gcc \
     git \
     libffi-dev \
     make \
     musl-dev \
     pkgconfig \
     py3-pyzmq \
     py3-pip \
     python3-dev \
     su-exec \
     wget \
     zeromq-dev \
     zip \
 && ln -s /usr/bin/python3.8 /usr/bin/python

## install Go
RUN apk --update-cache --allow-untrusted \
     --repository http://dl-4.alpinelinux.org/alpine/edge/community \
     --arch=x86_64 add \
     go

## jupyter notebook
RUN ln -s /usr/include/locale.h /usr/include/xlocale.h \
 && pip3 install jupyter notebook tornado ipykernel

## install gophernotes
RUN mkdir -p /go/src/github.com/gopherdata \
 && cd /go/src/github.com/gopherdata \
 && git clone https://github.com/gopherdata/gophernotes.git \
 && cd /go/src/github.com/gopherdata/gophernotes \
 && GOPATH=/go GO111MODULE=on go install . \
 && cp /go/bin/gophernotes /usr/local/bin/ \
 && mkdir -p ~/.local/share/jupyter/kernels/gophernotes \
 && cp -r ./kernel/* ~/.local/share/jupyter/kernels/gophernotes \
 && cd - \
 ## clean
 && find /usr/lib/python3.8 -name __pycache__ | xargs rm -r \
 && rm -rf \
     /root/.[acpw]* \
     ipaexg00301* \
 && rm -rf /var/cache/apk/*

## Install terraform
RUN cd /usr/bin \
 && wget https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip \
 && unzip terraform_1.0.0_linux_amd64.zip \
 && rm -rf terraform_1.0.0_linux_amd64.zip

# Set GOPATH.
ENV GOPATH /go

## Intall terraform provider pluging to generate documetnation
RUN go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs \
 && mv /go/bin/tfplugindocs /usr/bin/tfplugindocs

EXPOSE 8888
CMD [ "jupyter", "notebook", "--no-browser", "--allow-root", "--ip=0.0.0.0", "--NotebookApp.token=''"]
