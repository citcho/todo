module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  providers = {
    aws = aws.todo
  }

  name            = "${local.project}-${local.env}-vpc"
  azs             = formatlist("${local.region}%s", ["a", "c", "d"])
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  database_subnet_group_name = "${local.project}-${local.env}-db-subnet-group"

  create_database_subnet_group       = true
  create_database_subnet_route_table = true
  map_public_ip_on_launch            = true

  tags = {
    Name = "${local.project}-${local.env}-vpc"
  }
}
