# Certificates

You need to generate certificates for this program to work. How to do it you can find here - [click](https://github.com/durudex/durudex-gateway/blob/main/certs/README.md).

**If you do not want to use tls connection change [configs/main](https://github.com/durudex/durudex-sample-service/blob/main/configs/main.yml)**:
```yml
grpc:
    tls:
        enable: false
```
