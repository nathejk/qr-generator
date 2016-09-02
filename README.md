# A QR code generator in a Docker container
Starts an HTTP server (listening on port 80) that generates QR codes. Once started, the service accepts the following two parameters:
* ```data```: (Required) The (URL encoded) string that should be encoded in the QR code
* ```size```: (Optional) The size of the resulting image (default: 250)

E.g. ```http://.../?data=Hello%2C%20world&size=300```

## Building
Type ```make``` to compile the go files or ```make build``` to build the docker container.

## References
* Barcode Library: https://github.com/boombuler/barcode
