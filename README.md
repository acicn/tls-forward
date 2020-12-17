# tls-forward

secure regular TCP connection with TLS

## Usage

**Environment Variables**

* `FORWARD_BIND` address to listen, default to `:6443`
* `FORWARD_TARGET` proxy target
* `FORWARD_TLS_CRT` tls certificate file, default to `/data/tls.crt`
* `FORWARD_TLS_KEY` tls certificate key file, default to `/data/tls.key`

## Credits

Guo Y.K., MIT License
