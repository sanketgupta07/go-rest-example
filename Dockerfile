## We specify the base image we need for our
## go application
FROM golang:latest as build-env
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
FROM gcr.io/distroless/base
COPY --from=build-env /app/main /
CMD ["/main"]
## docker build -t go-rest-example .
## docker run -it --name test -p 8080:8080 go-rest-example
# PS C:\Users\sanket> curl http://localhost:8080


# StatusCode        : 200
# StatusDescription : OK
# Content           : Welcome home!
# RawContent        : HTTP/1.1 200 OK
#                     Content-Length: 13
#                     Content-Type: text/plain; charset=utf-8
#                     Date: Tue, 01 Sep 2020 18:37:56 GMT

#                     Welcome home!
# Forms             : {}
# Headers           : {[Content-Length, 13], [Content-Type, text/plain; charset=utf-8], [Date, Tue, 01 Sep 2020 18:37:56
#                     GMT]}
# Images            : {}
# InputFields       : {}
# Links             : {}
# ParsedHtml        : mshtml.HTMLDocumentClass
# RawContentLength  : 13

