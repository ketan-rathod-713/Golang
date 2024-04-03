VERSION ?= 1

generateAuthProtoCode:
	# First Create Directory Inside proto
	cd ./../proto && mkdir auth-server && cd .. && cd auth-server && protoc --go_out=./../proto/auth-server/ --go-grpc_out=./../proto/auth-server/  ./protobuf/v1/auth.proto
	# This will build auth package inside proto folder and will generate code in that folder only