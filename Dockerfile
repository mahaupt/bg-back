# build container
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /bg-back

## production container
FROM scratch

WORKDIR /

COPY --from=build /app/.env.example /.env
COPY --from=build /bg-back /bg-back

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/bg-back"]