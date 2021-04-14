package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestProjectBasicSuccess(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "{{cookiecutter.repository_name}}-basic",
		Upgrade:      true,
		Vars: map[string]interface{}{
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
}
