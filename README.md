[![CircleCI](https://circleci.com/gh/shoito/aws-insights.svg?style=svg)](https://circleci.com/gh/shoito/aws-insights)

# AWS Insights CLI

## Installation

Use `go get` to retrieve the CLI to add it to your `GOPATH` workspace, or
project's Go module dependencies.

    go get github.com/shoito/aws-insights

To update the CLI use `go get -u` to retrieve the latest version of the CLI.

	go get -u github.com/shoito/aws-insights

## Usage
Multiple services.
```sh
$ iaws --service=ec2,rds,s3 -o excel -f output.xlsx
```

Single service.
```sh
$ iaws ec2 -o excel -f output.xlsx
```

Help.
```sh
$ iaws -h
AWS Insights CLI

Usage:
  iaws [flags]
  iaws [command]

Available Commands:
  ec2         Amazon EC2 Insights Command
  help        Help about any command
  rds         Amazon RDS Insights Command
  s3          Amazon S3 Insights Command

Flags:
  -c, --config string         Configuration file (default "$HOME/.iaws.yaml")
  -f, --file string           Output file
  -h, --help                  help for iaws
  -o, --output string         The formatting style for command output (excel, pdf, json, ...) (default "excel")
      --profile string        Use a specific profile from your credential file
      --region string         The region to use. Overrides config/env settings (default "ap-northeast-1")
      --service stringArray   Services(ec2,rds,s3,...). (default "all")

Use "iaws [command] --help" for more information about a command.
```

```sh
$ iaws ec2 -h
Amazon EC2 Insights Command

Usage:
  iaws ec2 [flags]

Flags:
  -h, --help   help for ec2

Global Flags:
  -c, --config string    Configuration file (default "$HOME/.iaws.yaml")
  -f, --file string      Output file
  -o, --output string    The formatting style for command output (excel, pdf, json, ...) (default "excel")
      --profile string   Use a specific profile from your credential file
      --region string    The region to use. Overrides config/env settings (default "ap-northeast-1")
```

## Design
![Design](https://g.gravizo.com/svg?%40startuml%3B%0D%0A%0D%0Aactor%20User%0D%0Aparticipant%20%22cmd%22%20as%20C%0D%0Aparticipant%20%22service%22%20as%20S%0D%0Aparticipant%20%22writer%22%20as%20W%0D%0A%0D%0AUser%20-%3E%20C%3A%20exec%0D%0Aactivate%20C%0D%0A%0D%0AC%20-%3E%20S%3A%20Describe%20Request%0D%0Aactivate%20S%0D%0A%0D%0AS%20--%3E%20C%3A%20Describe%20Response%0D%0Adeactivate%20S%0D%0A%0D%0AC%20-%3E%20W%3A%20Write%20Request%0D%0Aactivate%20W%0D%0A%0D%0AW%20--%3E%20C%3A%20Write%20Response%0D%0Adeactivate%20W%0D%0A%0D%0AC%20--%3E%20User%3A%0D%0Adeactivate%20C%0D%0A%0D%0A%40enduml)

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md)

## License
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)