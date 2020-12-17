package main

import (
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func exit(err *error) {
	if *err != nil {
		log.Println("exited with error:", (*err).Error())
		os.Exit(1)
	}
}

var (
	envBind   = strings.TrimSpace(os.Getenv("FORWARD_BIND"))
	envTarget = strings.TrimSpace(os.Getenv("FORWARD_TARGET"))
	envTLSCrt = strings.TrimSpace(os.Getenv("FORWARD_TLS_CRT"))
	envTLSKey = strings.TrimSpace(os.Getenv("FORWARD_TLS_KEY"))
)

func main() {
	var err error
	defer exit(&err)

	if envBind == "" {
		envBind = ":6443"
	}

	if envTarget == "" {
		err = errors.New("missing environment $PROXY_TARGET")
		return
	}

	if envTLSCrt == "" {
		envTLSCrt = "/data/tls.crt"
	}
	if envTLSKey == "" {
		envTLSKey = "/data/tls.key"
	}

	var c tls.Certificate
	if c, err = tls.LoadX509KeyPair(envTLSCrt, envTLSKey); err != nil {
		return
	}

	var l net.Listener

	if l, err = tls.Listen("tcp", envBind, &tls.Config{
		Certificates: []tls.Certificate{c},
	}); err != nil {
		return
	}

	for {
		var c net.Conn
		if c, err = l.Accept(); err != nil {
			return
		}
		go handle(c, envTarget)
	}
}

func handle(c net.Conn, target string) {
	defer c.Close()

	nc, err := net.Dial("tcp", target)
	if err != nil {
		log.Println("failed to dial:", target, ", error:", err.Error())
		return
	}
	defer nc.Close()

	go io.Copy(c, nc)
	io.Copy(nc, c)
}
