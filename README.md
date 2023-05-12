# Lambda Invoker Cli

## About

Simple tool to allow you to easily invoke your lambdas to test them out from the command line.

## Prerequisites

[aws cli](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

## Setup

### Using Go (Easiest)

- `go install github.com/whutchinson98/lambda-invoker-cli@latest`

### Locally

- you will need go >= 1.19 installed and setup on your machine
- git clone the repository
- cd `lambda-invoker-cli`
- `make build`

### From Artifact

- You can download the executable artifact for your OS from the **Releases** of this repository

## Usage

Run `lambda-invoker-cli --help` for a full list of commands and parameters
