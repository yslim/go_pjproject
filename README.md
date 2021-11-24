Golang PJSIP(pjproject) using Swig

1. Generate pjsua2.go pjsua2_wrap.cxx using swig

```console
$ cd $GOPATH/src
$ mkdir pjproject
$ cd pjproject
$ cp ${pjprject-src-dir}/pjsip-apps/src/swig/pjsua2.i .
$ cp ${pjprject-src-dir}/pjsip-apps/src/swig/symbols.i .
$ export CGO_CXXFLAGS="-I${pjproject-install-dir}/include"
$ swig -go -cgo -intgosize 64 $CGO_CXXFLAGS -c++ pjsua2.i
```

2. Insert cgo compile, link flags to pjsua2.go
* CXXFLAGS, LDFLAGS from pkgconfig/libpjproject.pc

```golang
package pjsua2

/*
#cgo CXXFLAGS: -I/Data/apphome/pjproject-2.10/include -g -O2 -DPJ_AUTOCONF=1  -O2 -DPJ_IS_BIG_ENDIAN=0 -DPJ_IS_LITTLE_ENDIAN=1
#cgo LDFLAGS: -L/Data/apphome/pjproject-2.10/lib -L/usr/local/opt/openssl/lib -L/usr/local -lpjsua2-x86_64-apple-darwin20.6.0 -lstdc++ -lpjsua-x86_64-apple-darwin20.6.0 -lpjsip-ua-x86_64-apple-darwin20.6.0 -lpjsip-simple-x86_64-apple-darwin20.6.0 -lpjsip-x86_64-apple-darwin20.6.0 -lpjmedia-codec-x86_64-apple-darwin20.6.0 -lpjmedia-x86_64-apple-darwin20.6.0 -lpjmedia-videodev-x86_64-apple-darwin20.6.0 -lpjmedia-audiodev-x86_64-apple-darwin20.6.0 -lpjmedia-x86_64-apple-darwin20.6.0 -lpjnath-x86_64-apple-darwin20.6.0 -lpjlib-util-x86_64-apple-darwin20.6.0  -lsrtp-x86_64-apple-darwin20.6.0 -lresample-x86_64-apple-darwin20.6.0  -lpj-x86_64-apple-darwin20.6.0 -lssl -lcrypto -lvpx -lm -lpthread  -framework Foundation -framework AppKit -lgnutls

#define intgo swig_intgo
typedef void *swig_voidp;
```

3. Go build

```console
$ go clean -cache
$ go build -x
$ go install
```