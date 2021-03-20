# Rectangle Properties

This program is built to determine below properties between 2 rectangles.

1. Intersection
1. Containment
1. Adjacency
    1. Sub line Adjacency
    1. Proper Adjacency
    1. Partial Adjacency
    1. No Adjacency

## Technologies

1. GoLang (v1.14)

This program is written in [GoLang](https://golang.org/) and also distributed as Docker image.

## Run Program

> A Makefile is provided for easier usage

```sh
# See available options
$ make

# Install dependencies
$ make tidy

# Build binary
$ make build

# Run an example
$ make test-example
````

## Contribution

Read [Go setup guide](https://golang.org/doc/install) to configure local environment.

```sh
# Run Test
$ make test
```

## Run Using Docker image

```sh
# RUn using docker image
$ docker run niravrk/rectangle ./rectangle-assessment execute 0 0 10 10 10 0 7 7
```
