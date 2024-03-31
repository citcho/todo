resource "aws_acm_certificate" "tokyo_cert" {
  provider = aws.todo

  domain_name       = "api.${local.domain}"
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate" "virginia_cert" {
  provider = aws.virginia

  domain_name       = local.domain
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}
