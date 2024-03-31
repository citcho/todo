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
data "aws_ssm_parameter" "app_port" {
  provider = aws.todo
  name     = "/${local.project}/${local.env}/app/port"
}
