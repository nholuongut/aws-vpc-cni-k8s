#!/bin/bash

set -eo pipefail

pushd ./test/integration

echo "Running integration test with the following configuration:
KUBECONFIG: $KUBECONFIG
CLUSTER_NAME: $CLUSTER_NAME
nholuongut_REGION: $nholuongut_REGION"


if [[ -n "${ENDPOINT}" ]]; then
  ENDPOINT_FLAG="--endpoint $ENDPOINT"
fi

while getopts "f:" arg; do
  case $arg in
    f) FOCUS=$OPTARG;
    echo "FOCUS: $FOCUS";
  esac
done

CLUSTER_INFO=$(nholuongut eks describe-cluster --name $CLUSTER_NAME --region $nholuongut_REGION $ENDPOINT_FLAG)

VPC_ID=$(echo $CLUSTER_INFO | jq -r '.cluster.resourcesVpcConfig.vpcId')
SERVICE_ROLE_ARN=$(echo $CLUSTER_INFO | jq -r '.cluster.roleArn')
ROLE_NAME=${SERVICE_ROLE_ARN##*/}
TEST_FAILED=false
 
START=$SECONDS

for dir in */  
do

    cd "${dir%*/}"  
    if [ -n "$FOCUS" ]; then
        ( ginkgo --focus="$FOCUS" -v -r -- --cluster-kubeconfig=$KUBECONFIG --cluster-name=$CLUSTER_NAME --nholuongut-region=$nholuongut_REGION --nholuongut-vpc-id=$VPC_ID ) || TEST_FAILED=true
    else
        ( ginkgo -v -r -- --cluster-kubeconfig=$KUBECONFIG --cluster-name=$CLUSTER_NAME --nholuongut-region=$nholuongut_REGION --nholuongut-vpc-id=$VPC_ID ) || TEST_FAILED=true
    fi
    cd ..    
     
done

popd

# If any of the test failed, exit with non zero exit code
if [ $TEST_FAILED = true ]; then
  exit 1
fi
