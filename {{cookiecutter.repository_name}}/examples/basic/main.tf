locals {
  tags = {
    terraform_id = var.namespace
    Namespace    = var.namespace
    Environment  = var.stage
    Stage        = var.stage
  }
}

module "main" {
  source = "git://{{cookiecutter.repository}}.com/{{ cookiecutter.repository_owner }}/{{ cookiecutter.repository_name }}.git?ref={{ cookiecutter.version }}"
  name   = {{ cookiecutter.repository_name }}
}
