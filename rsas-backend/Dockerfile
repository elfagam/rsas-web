# build stage
FROM golang:latest AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /rsas-backend cmd/api/main.go

# production stage
FROM alpine:latest AS production-stage
RUN apk add --no-cache tzdata
WORKDIR /
COPY --from=build-stage /rsas-backend /rsas-backend
COPY .env .env
EXPOSE 8000
CMD ["/rsas-backend"]
