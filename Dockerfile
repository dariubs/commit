FROM golang:1.10

LABEL maintainer="poshtehani@gmail.com"

ARG app_env=development
ENV COMMIT_ENV=$app_env

# install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN PATH="/go/bin:$PATH"

COPY . /go/src/github.com/dariubs/commit
WORKDIR /go/src/github.com/dariubs/commit
RUN dep ensure
RUN go install github.com/dariubs/commit

WORKDIR /go/src/github.com/dariubs/commit
CMD /go/bin/commit
