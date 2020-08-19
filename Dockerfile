FROM golang AS build

ADD ./ /opt/build/GOSecretProject/
WORKDIR /opt/build/GOSecretProject/cmd
RUN go build main.go

FROM ubuntu:18.04 AS release

EXPOSE 5432

EXPOSE 8001

WORKDIR /opt/build/GOSecretProject/cmd/
COPY --from=build /opt/build/GOSecretProject/cmd/ ./
ADD
CMD ./main