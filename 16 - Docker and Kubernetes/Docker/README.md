# Using docker for deploying GO projects

### Sample Dockerfile

```bash
# syntax=docker/dockerfile:1
FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /wapp

EXPOSE 8080

ENTRYPOINT ["/wapp"]
```

### Building the image

```bash
docker build --tag go-docker-test .
```

If the file is not named `Dockerfile`,

```bash
docker build -f Dockerfile.test --tag docker-test
```

### Running the file

Since we need to create a portfwd tunnel we have to use `--publish` to create the tunnel

```bash
docker run --publish 8080:8080 go-docker-test
```
