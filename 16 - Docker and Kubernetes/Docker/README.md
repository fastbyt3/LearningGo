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

## Docker Multi-stage builds

> [https://docs.docker.com/language/golang/build-images/#multi-stage-builds](Multi stage builds go docs)

- create a docker instance to build the executable
- create another instance -> place the executable in the new instance -> execute it

## Further development

Read: [https://docs.docker.com/language/golang/develop/](Containers for development docs)
