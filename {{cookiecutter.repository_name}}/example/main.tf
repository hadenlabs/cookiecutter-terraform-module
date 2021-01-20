terraform {
  required_version = ">=0.13.0"
}

module "main" {
  source = "git://{{cookiecutter.repository}}.com/{{ cookiecutter.repository_owner }}/{{ cookiecutter.repository_name }}.git?ref={{ cookiecutter.version }}"
  name   = var.repository_name
}