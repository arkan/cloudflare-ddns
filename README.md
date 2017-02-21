# cloudflare-ddns

This small application will keep a Cloudflare DNS record up to date with the Docker public IP address.

## Usage

### Basic
```
CF_HOST=xx.yy.com CF_API_KEY=xxxxx -e CF_API_EMAIL=your@email.com go run cmd/cloudflare-ddns/main.go
```

## Docker

```
docker run -e CF_HOST=xx.yy.com -e CF_API_KEY=xxxxx -e CF_API_EMAIL=your@email.com berndinox/cloudflare-ddns:latest
```

## Compose v3 
```
version: '3'
services:
  dns:
    image: berndinox/cloudflare-ddns:latest
    environment:
      - CF_HOST=${DOMAIN}
      - CF_API_KEY=${KEY}
      - CF_API_EMAIL=${MAIL}
    deploy:
      replicas: 1
```


Added CURL to perform health checks from inside the container
eg.:
```
    healthcheck:
      test: ["CMD", "curl", "-f", "http://${DOMAIN}"]
      interval: 20s
      timeout: 10s
      retries: 3
```


Soure:
Copyright © 2016 Florian Bertholin
See the [LICENSE](./LICENSE) (MIT) file for more details.

