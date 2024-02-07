FROM golang:1.21.7-bullseye as env-builder

RUN go install github.com/posener/complete/gocomplete@latest \
    && go install github.com/766b/go-outliner@latest \
    && go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/josharian/impl@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install golang.org/x/tools/gopls@latest

FROM golang:1.21.7-bullseye as dev

ENV USER vscode

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    less \
    && apt-get clean

RUN adduser ${USER}
USER ${USER}

# If you don't needs tools, comment out this lines.
COPY --from=env-builder /go/bin /go/bin
RUN gocomplete -install -y
