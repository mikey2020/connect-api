FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/connect

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install connect

ENV SERVER=mongodb://flash:flash@ds059365.mlab.com:59365/connect
ENV DATABASE=connect

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/connect

CMD ["/go/bin/connect"]

EXPOSE 3000


