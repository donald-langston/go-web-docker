FROM golang:1.15.7-buster
ENV G0111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/mathapp
ARG GROUP_ID
ARG USER_ID
RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
RUN echo 'export GOPATH="$HOME/go"' >> ~/.profile
RUN echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile
#USER $APP_USER
WORKDIR $APP_HOME
COPY ./src .
RUN go mod init mathapp
RUN go get github.com/beego/beego/v2
RUN go get github.com/beego/bee/v2
RUN go mod tidy
RUN go mod vendor
USER $APP_USER
EXPOSE 8010
CMD ["bee", "run"]