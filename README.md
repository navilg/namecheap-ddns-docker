# Namecheap DDNS docker client

## When to use ?

If your server do not have static IP, i.e. When Public IP of your server / router keep changing, This will automatically update new IP to Namecheap record.

## Why to use it ?

* Easy to setup

* Lightweight

* No cronjob configuration required

* Logs for visibility

* Open source

## How to use it

* Change nameserver of you domain to BasicDNS or PremiumDNS or FreeDNS. `Namecheap Account -> Domain List -> Manage -> Domain -> NAMESERVERS -> Choose BasicDNS from dropdown`

* Enable Dynamic DNS for your domain from Namecheap. `Namecheap Account -> Domain List -> Manage -> Advanced DNS -> Dynamic DNS -> Toggle Status`

* Copy the Dynamic DNS password which is generated after enabling Dynamic DNS. Keep it safe and handy.

* Add a record of type `A + Dynamic DNS` with required host name.

* Install docker on server.

* Run below command

Suppose, You need DDNS for `server.example.com`

```
# For Linux running on amd64 arch
docker run --name server.example.com -d --restart unless-stopped -e NC_HOST='server' -e NC_DOMAIN='example.com' -e NC_PASS='DynamicDDNSPa2w0rd' linuxshots/namecheap-ddns:1.0.0

# For linux running on arm64 arch
docker run --name server.example.com -d --restart unless-stopped -e NC_HOST='server' -e NC_DOMAIN='example.com' -e NC_PASS='DynamicDDNSPa2w0rd' linuxshots/namecheap-ddns:arm64v8-linux-1.0.0
```

Here, 
`NC_HOST` is host name added in Namecheap record.

`NC_DOMAIN` is your domain name.

`NC_PASS` is your Dynamic DDNS password which is generated from Namecheap.

* Check the log

```
docker logs server.example.com
```

* To stop, start and remove DDNS.

```
docker stop server.example.com # To stop
docker start server.example.com # To start after its stopped
docker rm server.example.com -f # To remove
```

## Build your own image

To build your own image

* Clone this repo

* Run docker build

```
# Replace OS and ARCH values with valid values of GOlang environment variables GOOS and GOARCH.
docker build --build-arg OS=linux --build-arg ARCH=amd64 --build-arg VERSION=1.0.0 -t linuxshots/namecheap-ddns:1.0.0 . 
```

NOTE: This sets the TTL to Automatic i.e. 30 minutes. Currently, There is no way provided by Namecheap to set custom TTL in Dynamic DDNS.