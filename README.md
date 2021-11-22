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

## acknowledgement

- [**techcraftlabs/mpesa**](https://github.com/techcraftlabs/mpesa)
- [**techcraftlabs/airtel**](https://github.com/techcraftlabs/airtel)
- [**techcraftlabs/tigopesa**](https://github.com/techcraftlabs/tigopesa)


Under the hood pesakit uses the mentioned libraries

## articles

- [**Pesakit: Developer companion for Mobile API integration development and testing stage.**](https://medium.com/@piusalfred/pesakit-developer-companion-for-mobile-api-integration-development-and-testing-stage-f90b744527b6)

## twitter

- [Thread - Pesakit](https://twitter.com/thepiusalfred/status/1460319910836555787?s=20)