# pesakit
pesakit is a highly configurable commandline tool that comes in handy during testing and
development of systems that integrate with mobile money vendors. With pesakit you can send
C2B (pushpay) requests or B2C (disbursement) requests. You can do this on either production
or staging stage, it just depends on how you configure it. Meaning you can use pesakit in
real production env.


Supported Vendors: Tigo Pesa, Airtel Money and Vodacom MPESA. There is a possibility to use
the tool in countries that the vendors API supports e.g GHANA for MPESA. But the tool has been
tested for Tanzania only.

## commands
```bash
  pesakit disburse --phone=0784AAAAAA --amount=1000 --description=testing --id=BAGATATSVSNSUXNJ  
```

```bash
  pesakit push --phone=067AAAAAAA --amount=1000 --description=testing --reference=BAGATATSVSNSUXNJ    
```

## docker image

to build the image

```bash
make docker
```
then run,

```bash
docker run pesakit push --phone=0784956141 --amount=1000 --description=testing --id=BAGATsjksndhjSNSUXNJ  
```

you can also download the pre-built image from our github container registry

```bash
docker pull ghcr.io/techcraftlabs/pesakit:latest

```
run the image,

```bash
docker run --env-file .env ghcr.io/techcraftlabs/pesakit:latest
```

## note

Look at the file [ENV.md](ENV.md) to see how to set env vars for pesakit

## use as library
pesakit can be used in go projects as a library

```bash
go get github.com/pesakit/pesakit
```

```go
package demo

import (
	"github.com/pesakit/pesakit"
	"github.com/pesakit/pesakit/airtel"
	"github.com/pesakit/pesakit/mpesa"
	"github.com/pesakit/pesakit/tigo"
)

func main() {
	pc := pesakit.NewClient(&airtel.Client{}, &tigo.Config{}, &mpesa.Client{})
}


```
