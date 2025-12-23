FROM golang:1.25.5-trixie AS builder

WORKDIR /usr/src/app
COPY . .

# Install dependencies (npm)
RUN apt-get update && apt-get install -y nodejs npm
RUN npm install clean-css-cli -g

# Generate the minified static HTML files.
RUN go run ./templates/templatizer.go

# Minify the CSS file in-place.
RUN cleancss -o statics/styles.css statics/styles.css

# Compile the actual app, which will embed the above files.
RUN go build -v -o /run-app .


FROM debian:trixie

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
