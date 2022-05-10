# Certificates

You need to configure the tls configuration:
```yml
grpc:
  tls:
    ca-cert: "path to cert ca"
    cert: "path to cert"
    key: "path to cert key"
```

**If you do not want to use tls connection change**:
```yml
grpc:
  tls:
    enable: false
```
