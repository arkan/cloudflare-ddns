# cloudflare-ddns

This small application will keep a Cloudflare DNS record up to date with the Docker public IP address.

## Usage

### Basic
```
CF_HOST=xx.yy.com CF_API_KEY=xxxxx -e CF_API_EMAIL=your@email.com go run cmd/cloudflare-ddns/main.go
```

## Docker

```
docker run -e CF_HOST=xx.yy.com -e CF_API_KEY=xxxxx -e CF_API_EMAIL=your@email.com arkan/cloudflare-ddns:latest
```

## unRAID

You can use my [unraid-templates](https://github.com/arkan/unraid-templates).


#Copyright

Copyright © 2016 Florian Bertholin

See the [LICENSE](./LICENSE) (MIT) file for more details.

