Golang PJSIP(Pjproject) Swig

1. Generate pjsua2.go pjsua2_wrap.cxx

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
#cgo CXXFLAGS: -I/Data/apphome/lib/static/include -g -O2 -Wno-delete-non-virtual-dtor
#cgo LDFLAGS: -L/Data/apphome/lib/applib -L/usr/local/opt/openssl/lib -lpjsua2-x86_64-apple-darwin17.7.0 -lstdc++ -lpjsua-x86_64-apple-darwin17.7.0 -lpjsip-ua-x86_64-apple-darwin17.7.0 -lpjsip-simple-x86_64-apple-darwin17.7.0 -lpjsip-x86_64-apple-darwin17.7.0 -lpjmedia-codec-x86_64-apple-darwin17.7.0 -lpjmedia-x86_64-apple-darwin17.7.0 -lpjmedia-videodev-x86_64-apple-darwin17.7.0 -lpjmedia-audiodev-x86_64-apple-darwin17.7.0 -lpjmedia-x86_64-apple-darwin17.7.0 -lpjnath-x86_64-apple-darwin17.7.0 -lpjlib-util-x86_64-apple-darwin17.7.0 -lsrtp-x86_64-apple-darwin17.7.0 -lresample-x86_64-apple-darwin17.7.0 -lpj-x86_64-apple-darwin17.7.0 -lssl -lcrypto -lm -lpthread -framework Foundation -framework AppKit

#define intgo swig_intgo
typedef void *swig_voidp;
```

3. Go build

```console
$ go clean -cache
$ go build -x
$ go install
```