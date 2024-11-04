#!/usr/bin/env bash

check_nholuongut_credentials() {
    nholuongut sts get-caller-identity --query "Account" ||
        ( echo "No nholuongut credentials found. Please run \`nholuongut configure\` to set up the CLI for your credentials." && exit 1)
}

ensure_ecr_repo() {
    echo "Ensuring that $2 exists for account $1"
    local __registry_account_id="$1"
    local __repo_name="$2"
    if ! `nholuongut ecr describe-repositories --registry-id "$__registry_account_id" --repository-names "$__repo_name" >/dev/null 2>&1`; then
        echo "creating ECR repo with name $__repo_name in registry account $__registry_account_id"
        nholuongut ecr create-repository --repository-name "$__repo_name"
    fi
}

emit_cloudwatch_metric() {
    nholuongut cloudwatch put-metric-data --metric-name $1 --namespace TestExecution --unit None --value $2 --region $nholuongut_DEFAULT_REGION
}

