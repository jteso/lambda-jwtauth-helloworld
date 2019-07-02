## Pre-requisites

- Go 1.12
- serverless framework

## Getting Started

1. Initialize the go module (optional, only if you dont want to use the `$GOPATH`)

```
Usage: make init modname=<name_of_your_module>

Eg. make init modname=lamda-with-custom-authorizer
```

2. Deploy the solution:

**Important** Amend the `serverless.yml` file to reflect the stages that you want to enable. By the default, `dev` will be used.

```
make deploy stage=<stage>

Eg. make deploy stage=dev
```

3. Questions:

You can run: `make` to see all the operations available