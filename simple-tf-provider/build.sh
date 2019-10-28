#!/usr/bin/env bash

set -e

terraform_dir="infra"

go build -o $terraform_dir/terraform-provider-todo

cd ./$terraform_dir
terraform init

echo "Navigate to ./$terraform_dir to try your provider out!"


