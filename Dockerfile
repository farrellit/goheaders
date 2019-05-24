FROM golang as builder
ADD main.go main.go
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-extldflags "-static"' -o /goheaders main.go

FROM scratch
COPY --from=builder /goheaders /goheaders
ENTRYPOINT ["/goheaders"]
CMD []
