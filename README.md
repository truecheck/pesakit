# pesakit
pesakit is a highly configurable commandline tool that comes on handy during testing and
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

## use as library
pesakit can be used in go projects as a library

```bash
go get github.com/techcraftlabs/pesakit
```

```go
package demo

import (
	"github.com/techcraftlabs/pesakit"
	"github.com/techcraftlabs/pesakit/airtel"
	"github.com/techcraftlabs/pesakit/mpesa"
	"github.com/techcraftlabs/pesakit/tigo"
)

func main() {
	pc := pesakit.NewClient(&airtel.Client{}, &tigo.Config{}, &mpesa.Client{})
}


```
