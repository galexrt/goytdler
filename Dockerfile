FROM alpine:3.7
LABEL maintainer="Alexander Trost <galexrt@googlemail.com>"

ADD .build/linux-amd64/goytdler /bin/goytdler

RUN apk add --no-cache youtube-dl && \
    mkdir /goytdler-output

VOLUME ["/goytdler-output"]

WORKDIR "/"

ENTRYPOINT ["/bin/goytdler"]
