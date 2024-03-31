resource "aws_cloudwatch_log_group" "app" {
  provider = aws.todo
  name     = "${local.project}-${local.env}-app"
}