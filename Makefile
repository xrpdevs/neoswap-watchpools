ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

.ONESHELL:
default:
	echo "Installing dependencies"

	echo "Building watchpools"
	go build
	strip test
	echo 'All done! Now try "make run" or "make install"'


run:
	go run test.go

.ONESHELL:
install:
	echo "Installing watchPools..."

	cp ./watchPools.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl enable watchPools.service
	systemctl start watchPools.service
	echo "All done!"
