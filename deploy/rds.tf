resource "aws_db_instance" "app_db" {
  provider                  = aws.todo
  allocated_storage         = 10
  availability_zone         = "${local.region}a"
  db_subnet_group_name      = aws_db_subnet_group.app_db_subnet_group.name
  vpc_security_group_ids    = [aws_security_group.rds_sg.id]
  db_name                   = data.aws_ssm_parameter.db_name.value
  identifier                = "${local.project}-${local.env}-db"
  engine                    = "mysql"
  engine_version            = "8.0"
  instance_class            = "db.t4g.micro"
  username                  = data.aws_ssm_parameter.db_username.value
  password                  = data.aws_ssm_parameter.db_password.value
  port                      = tonumber(data.aws_ssm_parameter.db_port.value)
  parameter_group_name      = aws_db_parameter_group.app_db_parameter_group.name
  skip_final_snapshot       = false
  final_snapshot_identifier = "${local.project}-${local.env}-db-final-snapshot"
}

resource "aws_db_subnet_group" "app_db_subnet_group" {
  provider   = aws.todo
  name       = "${local.project}-${local.env}-db-subnet-group"
  subnet_ids = [module.vpc.private_subnets[0], module.vpc.private_subnets[1], module.vpc.private_subnets[2]]

  tags = {
    Name = "${local.project} DB subnet group"
  }
}

resource "aws_db_parameter_group" "app_db_parameter_group" {
  provider = aws.todo
  name     = "${local.project}-${local.env}-db-parameter-group"
  family   = "mysql8.0"

  parameter {
    name  = "character_set_server"
    value = "utf8mb4"
  }

  parameter {
    name  = "character_set_client"
    value = "utf8mb4"
  }
}