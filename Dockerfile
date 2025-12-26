# syntax=docker.io/docker/dockerfile:1.19
# We use two different build environments to avoid installing npm inside docker
# The first one has npm and we use it to minify some files (e.g., css)
FROM node:25-trixie AS builder_npm

WORKDIR /usr/src/app
RUN npm install --save-dev @protobuf-ts/plugin protoc
RUN npm install --save-dev typescript ts-loader
RUN npm install --save-dev webpack webpack-cli
RUN npm install --save-dev css-loader style-loader
RUN npm install --save-dev @types/node
RUN npm install --save-dev @webtui/css
RUN npm install --save-dev @webtui/theme-catppuccin

# Copy everything except the local Node stuff.
COPY --exclude=node_modules --exclude=package.json --exclude=package-lock.json . .

# Compile the protobufs into the proto directory
RUN npx protoc --ts_out proto/ --proto_path proto proto/rc.proto

# Compile the TypeScript
# RUN tsc --project ./tsconfig.json --outDir generated/js

# Build the js+css bundle using webpack
RUN npx tsc -v
RUN npx webpack-cli -c ./webpack.config.js
RUN ls dist/bundle.js

# The second build environment has golang and we use it to build the app
# (which embeds some of the files we minified from npm above)
FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app

# Install the protocol buffer compiler
RUN apk update && apk add protoc

# Install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy everything else from the previous builder stage/s workdir, it's ok.
COPY --from=builder_npm /usr/src/app .

# Generate the minified static HTML files.
RUN go run ./templates/templatizer.go

# Run the protobuf compiler
RUN protoc --go_out=. --go_opt=paths=source_relative proto/rc.proto
RUN ls *.go
RUN ls proto
RUN cat proto/rc.pb.go

# Compile the actual app, which will embed the above files.
RUN go build -v -o /run-app .

FROM alpine:3.23

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
