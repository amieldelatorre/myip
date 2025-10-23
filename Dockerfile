FROM alpine:3.22 AS final
WORKDIR /app
RUN apk add --no-cache netcat-openbsd curl bind-tools


FROM golang:1.25.3 AS build
RUN useradd -ms /bin/sh -u 1001 app
USER app

WORKDIR /build
COPY --chown=app:app go.mod ./
RUN go mod download

COPY --chown=app:app ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/api ./cmd/api/


FROM final AS release
COPY --from=build /build/api /app/api
COPY templates /app/templates

EXPOSE 8080
CMD ["/app/api"]