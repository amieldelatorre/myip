FROM golang:1.22.3 AS build
RUN useradd -ms /bin/sh -u 1001 app
USER app

WORKDIR /build
COPY --chown=app:app go.mod ./
RUN go mod download

COPY --chown=app:app ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/api ./cmd/api/


FROM alpine as final
COPY --from=build /build/api /app/api

EXPOSE 8080
CMD ["/app/api"]