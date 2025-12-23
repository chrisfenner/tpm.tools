# We use two different build environments to avoid installing npm inside docker
# The first one has npm and we use it to minify some files (e.g., css)
FROM node:25-trixie AS builder_npm

WORKDIR /usr/src/app
COPY . .
RUN npm install clean-css-cli -g

# Minify the CSS file in-place.
RUN cleancss -o statics/styles.css statics/styles.css

# The second build environment has golang and we use it to build the app
# (which embeds some of the files we minified from npm above)
FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy everything else from the previous builder stage/s workdir, it's ok.
COPY --from=builder_npm /usr/src/app .

# Generate the minified static HTML files.
RUN go run ./templates/templatizer.go

# Compile the actual app, which will embed the above files.
RUN go build -v -o /run-app .

FROM alpine:3.23

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
