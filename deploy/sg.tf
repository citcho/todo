# ---------------------------------------------------------------------
# alb security group
# ---------------------------------------------------------------------
resource "aws_security_group" "alb_sg" {
  provider    = aws.todo
  name        = "${local.project}-${local.env}-alb-sg"
  description = "alb security group"
  vpc_id      = module.vpc.vpc_id
}

resource "aws_security_group_rule" "web_in_http" {
  provider          = aws.todo
  security_group_id = aws_security_group.alb_sg.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 80
  to_port           = 80
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "web_in_https" {
  provider          = aws.todo
  security_group_id = aws_security_group.alb_sg.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 443
  to_port           = 443
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "web_out_app" {
  provider                 = aws.todo
  security_group_id        = aws_security_group.alb_sg.id
  type                     = "egress"
  protocol                 = "tcp"
  from_port                = 80
  to_port                  = 80
  source_security_group_id = aws_security_group.app_sg.id
}

# ---------------------------------------------------------------------
# ecs security group
# ---------------------------------------------------------------------
resource "aws_security_group" "app_sg" {
  provider    = aws.todo
  name        = "${local.project}-${local.env}-app-sg"
  description = "application security group"
  vpc_id      = module.vpc.vpc_id
}

resource "aws_security_group_rule" "app_in_tcp80" {
  provider                 = aws.todo
  security_group_id        = aws_security_group.app_sg.id
  type                     = "ingress"
  protocol                 = "tcp"
  from_port                = 80
  to_port                  = 80
  source_security_group_id = aws_security_group.alb_sg.id
}

resource "aws_security_group_rule" "app_out_rds" {
  provider                 = aws.todo
  security_group_id        = aws_security_group.app_sg.id
  type                     = "egress"
  protocol                 = "tcp"
  from_port                = tonumber(data.aws_ssm_parameter.db_port.value)
  to_port                  = tonumber(data.aws_ssm_parameter.db_port.value)
  source_security_group_id = aws_security_group.rds_sg.id
}

resource "aws_security_group_rule" "app_out_tcp443" {
  provider          = aws.todo
  security_group_id = aws_security_group.app_sg.id
  type              = "egress"
  protocol          = "tcp"
  from_port         = 443
  to_port           = 443
  cidr_blocks       = ["0.0.0.0/0"]
}

# ---------------------------------------------------------------------
# rdb security group
# ---------------------------------------------------------------------
resource "aws_security_group" "rds_sg" {
  provider    = aws.todo
  name        = "${local.project}-${local.env}-rds-sg"
  description = "database security group"
  vpc_id      = module.vpc.vpc_id
}

resource "aws_security_group_rule" "db_in_app_sg" {
  provider                 = aws.todo
  security_group_id        = aws_security_group.rds_sg.id
  type                     = "ingress"
  protocol                 = "tcp"
  from_port                = tonumber(data.aws_ssm_parameter.db_port.value)
  to_port                  = tonumber(data.aws_ssm_parameter.db_port.value)
  source_security_group_id = aws_security_group.app_sg.id
}
