terraform {
  required_version = "= 1.7.5"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket  = "terraform-state-339713027168"
    key     = "terraform.tfstate"
    profile = "citcho_todo"
    region  = "ap-northeast-1"
    acl     = "bucket-owner-full-control"
  }
}

provider "aws" {
  alias   = "todo"
  region  = "ap-northeast-1"
  profile = "citcho_todo"
  default_tags {
    tags = {
      Project     = "todo",
      Environment = "prod",
      Terraform   = true,
    }
  }
}

provider "aws" {
  region = "us-east-1"
  alias  = "virginia"
  default_tags {
    tags = {
      Project     = "todo",
      Environment = "prod",
      Terraform   = true,
    }
  }
}
