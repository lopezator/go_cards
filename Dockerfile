# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/lopezator/go_cards

# Build the go_cards command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/lopezator/go_cards

# Run the go_cards command by default when the container starts.
ENTRYPOINT /go/bin/go_cards

# Run the custom command
CMD ["/go/src/github.com/lopezator/go_cards/bootstrap.sh"]

# Document that the service listens on port 8080.
EXPOSE 8080
EXPOSE 5432