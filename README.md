# ubersnap-middle-backend-test

This is the documentation for take-home test project for Ubersnap Middle Backend Position.

The requirement is based on this [Gist](https://gist.github.com/RofieSagara/64e797a707fc7ac5a24c2f3bd4e930ca)

This program is writen in Go, with [Fiber](https://gofiber.io/) as the Web Framework, and uses FFMPEG for the image processing functionality.

## Table Of Contents

<!--toc:start-->

- [ubersnap-middle-backend-test](#ubersnap-middle-backend-test)
  - [Table Of Contents](#table-of-contents)
  - [Disclaimer](#disclaimer)
  - [installation](#installation)
    - [native](#native)
    - [docker](#docker)
  - [api-testing](#api-testing)
  - [caveats](#caveats)
  <!--toc:end-->

## Disclaimer

I use Linux for making this project, so I can't write a guide on Windows/MacOS, but for MacOS I believe the steps is almost similar.

## installation

This project needs few dependencies to be installed, such as:

- Go (the version I use is 1.22.1)
- FFMPEG
- Docker or Podman (optional if you want to run it in a container)

Also before you compile it please add the following environment variables to your `.env` file:

- copy the .env.example to .env
- update the values

The .env variables needed is just PORT at the moment, but in the future is a good practice to have sensitive variables in the .env file instead of hard coding them in the code.

### native

Please install Go and FFMPEG according to your OS, Then go to the root of the project and run

`go build main.go.`

It will download all the dependencies and compile the project.

Then once you see the `main` binary, run it with `./main`

### docker

For containerizing the project, you need to use Docker or Podman installed.

In your terminal, run:

`podman build . --tag "<IMAGE-NAME>"`

It will download base image, and the dependencies and compile the project.

`podman run -p 8080:8080 --env-file=.env <IMAGE-NAME>`

That will port forward the port 8080, passed to the .env file and run the project in the container.

## api-testing

For testing purposes, I provide [Hoppscotch](https://hoppscotch.io/) json file to be used for testing, it's a Postman alternative that I personally used.

For importing the json to Hoppscotch, please go to this [docs](docs/Importing-Hoppscotch-Json-file-into-the-workspace.pdf)

I provided the necessarry `Request Body` and `Request Header` for testing purposes.

## caveats

This project is not perfect, there are few things that I would like to improve.

- Improving the route handler, to make it neater.
- The output image is written in `/tmp/`, which is not ideal, for the production code use S3 or any other cloud storage to serve it.
- This project isn't tested in Cloud, so it's not production ready.
- This docs is written in rush, if I have more time I can write it in a more structured way.
- This project is not reviewed by anyone, so I can't tell what else I need to Improve.
