FROM alpine:3.7
LABEL maintainer="Alexander Trost <galexrt@googlemail.com>"

RUN apk add --no-cache youtube-dl && \
    mkdir -p /goytdler-output

VOLUME ["/goytdler-output"]

WORKDIR "/goytdler-output"

ADD .build/linux-amd64/goytdler /bin/goytdler

ENTRYPOINT ["/bin/goytdler"]
