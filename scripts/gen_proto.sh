#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  protoc --plugin="protoc-gen-go=${GOPATH}/bin/protoc-gen-go" --plugin="protoc-gen-go-grpc=${GOPATH}/bin/protoc-gen-go-grpc" -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done

# #!/bin/bash
# CURRENT_DIR=$1
# rm -rf ${CURRENT_DIR}/genproto
# for x in $(find ${CURRENT_DIR}/protos/* -type d); do
#   protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=${CURRENT_DIR} \
#    --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
# done