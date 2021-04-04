FROM golang:1.16-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/banking cmd/server/main.go 

FROM scratch
COPY --from=build /bin/banking /bin/banking
ENTRYPOINT [ "/bin/banking" ]
