data "aws_route53_zone" "primary" {
  provider = aws.todo
  name     = "citcho.com"
}

resource "aws_route53_record" "app" {
  provider = aws.todo
  zone_id  = data.aws_route53_zone.primary.zone_id
  name     = "api.${local.domain}"
  type     = "A"
  alias {
    name                   = aws_alb.alb.dns_name
    zone_id                = aws_alb.alb.zone_id
    evaluate_target_health = true
  }
}

resource "aws_route53_record" "web" {
  provider = aws.todo
  zone_id  = data.aws_route53_zone.primary.zone_id
  name     = local.domain
  type     = "A"
  alias {
    name                   = aws_cloudfront_distribution.client_distribution.domain_name
    zone_id                = aws_cloudfront_distribution.client_distribution.hosted_zone_id
    evaluate_target_health = true
  }
}