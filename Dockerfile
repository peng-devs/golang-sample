# phase of building application
FROM golang:alpine AS BUILD

RUN adduser -D -g '' sample

WORKDIR /workspace
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


# phase of building docker image
FROM scratch

COPY --from=BUILD /etc/passwd /etc/passwd
COPY --from=BUILD /workspace/app /app

EXPOSE 8080
USER sample

ENTRYPOINT ["/app"]
