resource "aws_cloudfront_distribution" "client_distribution" {
  provider = aws.todo
  origin {
    domain_name              = aws_s3_bucket.client_origin.bucket_regional_domain_name
    origin_access_control_id = aws_cloudfront_origin_access_control.default.id
    origin_id                = aws_s3_bucket.client_origin.id
  }

  enabled             = true
  default_root_object = "index.html"

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    cache_policy_id  = data.aws_cloudfront_cache_policy.CachingDisabled.id
    target_origin_id = aws_s3_bucket.client_origin.id

    viewer_protocol_policy = "allow-all"
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["JP"]
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  custom_error_response {
    error_caching_min_ttl = 10
    error_code            = 403
    response_code         = 200
    response_page_path    = "/"
  }
}

resource "aws_cloudfront_origin_access_control" "default" {
  provider                          = aws.todo
  name                              = "${local.project}_${local.env}_cloudfront_origin_access_control"
  description                       = "${local.project}_${local.env}_cloudfront_origin_access_control"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

data "aws_cloudfront_cache_policy" "CachingDisabled" {
  provider = aws.todo
  name     = "Managed-CachingDisabled"
}
