FROM golang:1.18-rc-buster

ENV USER vscode

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    less \
    && apt-get clean

RUN adduser ${USER}
USER ${USER}

RUN go install github.com/posener/complete/gocomplete@latest \
    && gocomplete -install -y \
    && go install github.com/766b/go-outliner@latest
