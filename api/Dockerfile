FROM golang:latest AS go-build
ENV GOBIN=/go/bin
ADD . /go/src/api
WORKDIR /go/src/api
RUN CGO_ENABLED=0 GOOS=linux go install -i main.go





#Multistage is useful in order to reduce the size of the image
FROM busybox
COPY --from=go-build /go/bin/main main
EXPOSE 3000 
CMD ["./main"]