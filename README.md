**Golang PJSIP(Pjproject) Swig**

1. Generate pjsua2.go pjsua2\_wrap.cxx

	$ cd $GOPATH/src
	$ mkdir pjproject
	$ cd pjproject
	$ cp $pjprject-src-dir/pjsip-apps/src/swig/pjsua2.i .
	$ cp $pjprject-src-dir/pjsip-apps/src/swig/symbols.* .
	$ export CGO_CXXFLAGS="-I${pjproject-install-dir}/include"
	$ swig -go -cgo -intgosize 64 $CGO_CXXFLAGS -c++ pjsua2.i
