# Build the application from source
FROM golang:1.21 AS build-stage
WORKDIR /client
COPY client/* .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11
COPY --from=build-stage /app /app
# EXPOSE 55555
ENTRYPOINT ["/app"]
