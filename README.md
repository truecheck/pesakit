# pesakit
mobile money dev kit written in golang supporting among other things collection and disbursement for tigo,airtel and mpesa


## configuration

To configure pesakit to pick up your configurations you need to create a `.env` file.
The file should contain variables as listed in the [ENV.md](ENV.md) file. If you are running as docker
make sure to pass the env file and if you are using pesakit as an executable put the env file in the current working directory.


## docker

build
```bash

docker build -t pesakit .

```

run

```bash

docker run -it --env-file=.env pesakit

```


## use a pre-built container

```bash

docker run -it --env-file=.env ghcr.io/pesakit/pesakit:latest --format=json config print 

```






```text
NAME:
   pesakit - commandline tool to test/interact with Mobile Money API

USAGE:
   pesakit [global options] command [command options] [arguments...]

VERSION:
   1.0.0

DESCRIPTION:
   pesakit is a highly configurable commandline tool that comes on handy during testing and
   development of systems that integrate with mobile money vendors. With pesakit you can send
   C2B (pushpay) requests or B2C (disbursement) requests. You can do this on either production
   or staging stage, it just depends on how you configure it. Meaning you can use pesakit in
   real production env.
   
   Supported Vendors: Tigo Pesa, Airtel Money and Vodacom MPESA. There is a possibility to use
   the tool in countries that the vendors API supports e.g GHANA for MPESA. But the tool has been
   tested for Tanzania only.

AUTHOR:
   Pius Alfred <me.pius1102@gmail.com>

COMMANDS:
   config         configurations management
   callbacks, cb  monitor callbacks from mno
   push           send push requests to msisdn
   disburse       send money to phone number from your account
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose       Enable verbose mode (default: false)
   --conf value    configuration file path
   --debug         Enable debug mode (default: false)
   --format value  print format (text, json, yaml)
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)

COPYRIGHT:
   MIT Licence, Creative Commons
   
   

```



## acknowledgement

- [**techcraftlabs/mpesa**](https://github.com/techcraftlabs/mpesa)
- [**techcraftlabs/airtel**](https://github.com/techcraftlabs/airtel)
- [**techcraftlabs/tigopesa**](https://github.com/techcraftlabs/tigopesa)


Under the hood pesakit uses the mentioned libraries