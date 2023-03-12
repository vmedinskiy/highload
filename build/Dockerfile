FROM golang:1.19.7 AS build
WORKDIR /hgh
COPY . .
RUN make build_bin

FROM debian:buster-slim
WORKDIR /
COPY --from=build /hgh/bin /
RUN chmod +x /socnet
CMD /socnet version