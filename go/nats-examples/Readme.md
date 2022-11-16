# NATS 테스트

NATS 에코서버를 구동하여 테스트 하는 방법입니다.

[nats-echo 어플리케이션](nats-echo/main.go)
```shell
cd nats-echo
go run main.go -s demo.nats.io -geo=true ncoin

Echo Service ID: [NATS Echo Service]
Echo Service listening on [ncoin]
Echo Service (Status) listening on [ncoin.status]
[#1] Echoing from [ncoin] to [_INBOX.LCbkwTDl566hdrD8ZwOUtb.48qwMurY]: "hellohi"
[#2] Echoing from [ncoin] to [_INBOX.sCf6CZ7Gr0aCwwmSSPUzr9.m3RfOmB3]: "ncoinhello"
```

[nats-req 어플리케이션](nats-req/main.go)
```shell
cd nats-req
go run main.go -s demo.nats.io ncoin ncoinhello

Published [ncoin] : 'ncoinhello'
Received  [_INBOX.TsyLZt2bEqN3ndksWuXi6i.HHSNMeNg] : '[Gyeonggi-do, KR]: "ncoinhello"'

```

프로젝트 초기 구성 및 실행
```shell
go mod init main
go mod tidy
go run main.go
```