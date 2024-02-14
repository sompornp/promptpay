# Promptpay api (create & confirm)

## What is it?

Golang Gin based restful api for create and confirm promptpay qr code.

### Features
- Create promptpay request, input: phone or thai id
- Confirm promptpay request, input: uuid from above response
- Can set timeout
- Docker ready
- Unit tests

### Requirements

- Go (should be 1.21+)

### How to run

1. git clone https://github.com/sompornp/promptpay
2. go mod download
3. Rename or copy .env.example to .env. Take a look on each config value in the file or description in later section.
4. For docker run, use env.dev
5. Use makefile commands to run or manage the app
6. run `make run` to start the app
7. run `make dockerrun` to start the app in docker
8. run `make test` to run unit tests

## APIs

| Path | Method | Description                                                                  |
| ---- | ------ |------------------------------------------------------------------------------|
| /createPromptpay | POST | Create promptpay |
| /confirmPromptpay | POST | Confirm promptpay                                                            |

See postman or rest client vs code file for example requests in api_test folder.

## ENV Configs

dev.env is an example of environment file. The following are the environment variables that can be set.

TIMEOUT_IN_SECOND is required. It is the timeout for the promptpay response uuid in second.