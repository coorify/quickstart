FROM node:18 as febuilder
WORKDIR /src
COPY . .
RUN cd web && yarn cache clean && yarn install && yarn run build

FROM golang:1.22 as bebuilder
WORKDIR /src
COPY . .
COPY --from=febuilder /src/web/dist /src/web/dist
RUN CGO_ENABLED=0 go build -o app main.go

FROM alpine:3.8
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=bebuilder /src/app .
COPY --from=bebuilder /src/config.yml .
CMD [ "/root/app" ]

