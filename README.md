# Go HTTPS Reverse Proxy

A simple and efficient HTTPS reverse proxy written in Go. This tool allows you to expose your local HTTP services through an HTTPS endpoint, making it particularly useful for development and testing scenarios where HTTPS is required.

## Features

-   Expose local HTTP services via HTTPS
-   Support for self-signed certificates
-   Configurable target and proxy ports
-   Simple command-line interface

## Prerequisites

-   Go 1.22 or higher
-   SSL certificate and key files (included sample self-signed certificates)

## Installation

```bash
go build -o reverse-proxy main.go
```

### Command Line Arguments

-   `-port`: (Required) The target port of your local HTTP service
-   `-proxy`: (Optional) The port for the HTTPS proxy (default: 8443)

## SSL Certificates

The repository includes sample self-signed certificates (`my.crt` and `my.key`) for development purposes. For production use, please replace these with your own valid SSL certificates.

To generate new self-signed certificates for development:

```bash
openssl req -x509 -newkey rsa:4096 -keyout my.key -out my.crt -days 365 -nodes
```
