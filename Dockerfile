# build container
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bg-back

# ## production container
FROM scratch

WORKDIR /app

COPY --from=build /app/.env.example /app/.env
COPY --from=build /bg-back /app/bg-back

EXPOSE 8080

CMD ["/app/bg-back"]