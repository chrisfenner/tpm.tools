FROM golang:1.25.5-trixie AS builder

WORKDIR /usr/src/app
COPY . .

# Generate the minified static HTML files.
RUN go run ./templates/templatizer.go

# Compile the actual app, which will embed the above files.
RUN go build -v -o /run-app .


FROM debian:trixie

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
