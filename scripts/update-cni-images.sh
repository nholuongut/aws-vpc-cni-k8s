#!/usr/bin/env bash

# script to set the image on nholuongut-node daemonset for running tests

# Usage: Set test images as ENV variables $AMAZON_K8S_CNI & $AMAZON_K8S_CNI_INIT
# Run script to update nholuongut-node daemonset images

set -e

SCRIPTS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source "$SCRIPTS_DIR/lib/k8s.sh"
: "${REGION=us-west-2}"

# Select CNI manifest based on regions
if [[ $REGION == "cn-north-1" || $REGION == "cn-northwest-1" ]]; then
    nholuongut_K8S_CNI_MANIFEST="$SCRIPTS_DIR/../config/master/nholuongut-k8s-cni-cn.yaml"
elif [[ $REGION == "us-gov-east-1" ]]; then
    nholuongut_K8S_CNI_MANIFEST="$SCRIPTS_DIR/../config/master/nholuongut-k8s-cni-us-gov-east-1.yaml"
elif [[ $REGION == "us-gov-west-1" ]]; then
    nholuongut_K8S_CNI_MANIFEST="$SCRIPTS_DIR/../config/master/nholuongut-k8s-cni-us-gov-west-1.yaml"
else
    nholuongut_K8S_CNI_MANIFEST="$SCRIPTS_DIR/../config/master/nholuongut-k8s-cni.yaml"
fi

MANIFEST_IMG_VERSION=`grep "image:" $nholuongut_K8S_CNI_MANIFEST | cut -d ":" -f3 | cut -d "\"" -f1 | head -1`
IMAGE_REPOSITORY=`grep "image:" $nholuongut_K8S_CNI_MANIFEST | cut -d ":" -f2 | cut -d "/" -f1 | head -1`

# Replace the images in nholuongut-k8s-cni.yaml with the tester images when environment variables are set
if [[ -z $AMAZON_K8S_CNI ]]; then
    echo "Using latest CNI image from nholuongut-k8s-cni manifest"
else
    echo "Replacing CNI image in nholuongut-k8s-cni manifest with $AMAZON_K8S_CNI"
    sed -i'.bak' "s,$IMAGE_REPOSITORY/amazon-k8s-cni:$MANIFEST_IMG_VERSION, $AMAZON_K8S_CNI," "$nholuongut_K8S_CNI_MANIFEST"
fi
if [[ -z $AMAZON_K8S_CNI_INIT ]]; then
    echo "Using latest CNI init image from nholuongut-k8s-cni manifest"
else
    echo "Replacing CNI init image in nholuongut-k8s-cni manifest with $AMAZON_K8S_CNI_INIT"
    sed -i'.bak' "s,$IMAGE_REPOSITORY/amazon-k8s-cni-init:$MANIFEST_IMG_VERSION, $AMAZON_K8S_CNI_INIT," "$nholuongut_K8S_CNI_MANIFEST"
fi

echo "Applying $nholuongut_K8S_CNI_MANIFEST manifest"
kubectl apply -f $nholuongut_K8S_CNI_MANIFEST

check_ds_rollout "nholuongut-node" "kube-system" "10m"
