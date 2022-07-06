#!/bin/bash

cd .. && make install && cd ./examples

rm .terraform.lock.hcl
rm terraform.tfstate*

terraform init
terraform plan
terraform apply --auto-approve
