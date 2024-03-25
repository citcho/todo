module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  providers = {
    aws = aws.todo
  }

  name           = "${local.project}-${local.env}-vpc"
  azs            = formatlist("${local.region}%s", ["a", "c", "d"])
  public_subnets = ["172.31.32.0/20", "172.31.0.0/20", "172.31.16.0/20"]

  create_vpc                      = false
  create_database_subnet_group    = false
  create_egress_only_igw          = false
  create_elasticache_subnet_group = false
  create_igw                      = false
  create_redshift_subnet_group    = false

  manage_default_vpc               = true
  manage_default_network_acl       = true
  manage_default_route_table       = true
  manage_default_security_group    = true
  default_vpc_enable_dns_hostnames = true
  default_security_group_name      = "default"
  map_public_ip_on_launch          = true

  tags = {
    Name = "${local.project}-${local.env}-vpc"
  }
}
