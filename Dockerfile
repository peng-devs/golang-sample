# stage of building application
FROM golang:alpine AS BUILD

RUN adduser -D -g '' sample

WORKDIR /workspace
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


# stage of building docker image
FROM scratch

# copy user data cuz scratch image don't hav adduser cmd
COPY --from=BUILD /etc/passwd /etc/passwd
USER sample

COPY --from=BUILD /workspace/app /app
EXPOSE 8080

ENTRYPOINT ["/app"]
