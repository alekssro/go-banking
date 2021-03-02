FROM golang:1.15-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/banking

FROM scratch
COPY --from=build /bin/banking /bin/banking
ENTRYPOINT [ "/bin/banking" ]
