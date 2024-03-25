resource "aws_s3_bucket" "client_origin" {
  provider = aws.todo
  bucket = "${local.project}-${local.env}-client-origin"
}

resource "aws_s3_bucket_ownership_controls" "client_origin_ownership_controls" {
  provider = aws.todo
  bucket = aws_s3_bucket.client_origin.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_public_access_block" "client_origin_public_access_block" {
  provider = aws.todo
  bucket = aws_s3_bucket.client_origin.id

  block_public_acls       = false
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_acl" "client_origin_acl" {
  provider = aws.todo
  depends_on = [
    aws_s3_bucket_ownership_controls.client_origin_ownership_controls,
    aws_s3_bucket_public_access_block.client_origin_public_access_block,
  ]

  bucket = aws_s3_bucket.client_origin.id
  acl    = "public-read"
}

resource "aws_s3_bucket_policy" "allow_access_from_cloudfront" {
  provider = aws.todo
  bucket = aws_s3_bucket.client_origin.id
  policy = data.aws_iam_policy_document.allow_access_from_cloudfront.json
}

data "aws_iam_policy_document" "allow_access_from_cloudfront" {
  provider = aws.todo
  statement {
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }

    actions = [
      "s3:GetObject",
    ]

    resources = [
      "${aws_s3_bucket.client_origin.arn}/*",
    ]

    condition {
      test     = "StringEquals"
      variable = "AWS:SourceArn"
      values   = [aws_cloudfront_distribution.client_distribution.arn]
    }
  }
}
