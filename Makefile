all: build

build: proto certs

proto:
	protoc --go_out=plugins=grpc:. api/*.proto
	python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. api/*.proto
	@tree -hrtC api

certs:
	# create CA
	certstrap --depot-path certs init --common-name "gRPC Root CA"
	# create server cert request
	certstrap --depot-path certs request-cert --domain localhost
	# create client cert request
	certstrap --depot-path certs request-cert --cn grpc_client
	# sign server and client cert requests
	certstrap --depot-path certs sign --CA "gRPC Root CA" localhost
	certstrap --depot-path certs sign --CA "gRPC Root CA" grpc_client
	@tree -hrC certs

clean:
	rm -rf certs
	rm -f api/*.go api/*.py

server:
	TLS_ROOT_CA=$$(cat ../out/Reckon_Root_CA.crt) \
	TLS_SERVER_CERT=$$(cat ../out/localhost.crt) \
	TLS_SERVER_KEY=$$(cat ../out/localhost.key) \
	GRPC_SERVER_ADDRESS=localhost:3003 \
		go run server.go

client:
	TLS_ROOT_CA=$$(cat ../out/Reckon_Root_CA.crt) \
	TLS_CLIENT_CERT=$$(cat ../out/StatsModel.crt) \
	TLS_CLIENT_KEY=$$(cat ../out/StatsModel.key) \
	GRPC_SERVER_ADDRESS=localhost:3003 \
		python client.py
