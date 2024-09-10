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
docker run --name server.example.com -d --restart unless-stopped -e NC_HOST='server' -e NC_DOMAIN='example.com' -e NC_PASS='DynamicDDNSPa2w0rd' linuxshots/namecheap-ddns
```

Here, 
`NC_HOST` is host name added in Namecheap record.

`NC_DOMAIN` is your domain name.

`NC_PASS` is your Dynamic DDNS password which is generated from Namecheap.

You can also use an additional optional env variable:

`CUSTOM_IPCHECK_URL` is an optional variable where you can specify any URL that is used to get the current IP. The program uses HTTP GET method without additional parameters and expects a valid JSON with `ip` field.

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

### Alternative Configuration using Docker Compose

If you prefer to configure and manage your Docker containers using Docker Compose, follow the steps outlined below. This approach simplifies the configuration and management by allowing you to define all necessary parameters in a `docker-compose.yml` file.

1. Ensure that Docker and Docker Compose are installed on your server.

2. Create a `docker-compose.yml` file with the following content. Replace the placeholder values with your actual configuration:

    ```yaml
    services:
      ddns:
        image: linuxshots/namecheap-ddns
        container_name: server.example.com
        environment:
          - NC_HOST=server
          - NC_DOMAIN=example.com
          - NC_PASS=DynamicDDNSPa2w0rd
        restart: unless-stopped
    ```

3. Run the following command in the directory containing your `docker-compose.yml` file to start the container:

    ```sh
    docker-compose up -d
    ```

4. Check the logs to ensure the service is running correctly:

    ```sh
    docker-compose logs ddns
    ```

5. To manage the container (stop, start, remove), you can use Docker Compose commands:

    ```sh
    docker-compose stop ddns     # To stop the service
    docker-compose start ddns    # To start the service again after it’s stopped
    docker-compose down          # To stop and remove the container
    ```

By using Docker Compose, you simplify the deployment process and gain better control over your Docker services. This method is particularly useful for managing multiple containers and configurations in a unified manner.


## Build your own image

To build your own image

* Clone this repo

* Run docker build

```
# Replace OS and ARCH values with valid values of GOlang environment variables GOOS and GOARCH.
docker build --build-arg VERSION=1.0.0 -t linuxshots/namecheap-ddns . 
```

NOTE: This sets the TTL to Automatic i.e. 30 minutes. Currently, There is no way provided by Namecheap to set custom TTL in Dynamic DDNS.