data "aws_ssm_parameter" "app_client_host" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/app/client_host"
}

data "aws_ssm_parameter" "app_client_port" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/app/client_port"
}

data "aws_ssm_parameter" "app_port" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/app/port"
}

data "aws_ssm_parameter" "db_name" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/db/name"
}

data "aws_ssm_parameter" "db_username" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/db/username"
}

data "aws_ssm_parameter" "db_password" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/db/password"
}

data "aws_ssm_parameter" "db_port" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/db/port"
}

data "aws_ssm_parameter" "app_timezone" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/app/timezone"
}

data "aws_ssm_parameter" "app_bundebug" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/app/bundebug"
}
