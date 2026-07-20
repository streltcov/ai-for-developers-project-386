FROM golang:1.22-alpine AS build

WORKDIR /app

COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api/ ./

RUN CGO_ENABLED=0 go build -o /server ./cmd/server

FROM alpine:3.19

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /server .

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["./server"]
