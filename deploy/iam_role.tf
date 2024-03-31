data "aws_iam_policy" "s3_full_access" {
  provider = aws.todo
  arn      = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}


resource "aws_iam_role" "client_github_sync_s3_role" {
  provider    = aws.todo
  name        = "${local.project}_${local.env}_client_github_sync_s3_role"
  description = "github actions role."

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" : "Allow",
        "Principal" : {
          "Federated" : aws_iam_openid_connect_provider.repo.arn
        },
        "Action" : "sts:AssumeRoleWithWebIdentity",
        "Condition" : {
          "StringEquals" : {
            "token.actions.githubusercontent.com:aud" : "sts.amazonaws.com"
          },
          "StringLike" : {
            "token.actions.githubusercontent.com:sub" : "repo:citcho/hexisa_go_nal_todo:*"
          }
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "s3_full_access" {
  provider   = aws.todo
  policy_arn = data.aws_iam_policy.s3_full_access.arn
  role       = aws_iam_role.client_github_sync_s3_role.name
}

resource "aws_iam_openid_connect_provider" "repo" {
  provider = aws.todo
  url      = "https://token.actions.githubusercontent.com"

  client_id_list = [
    "sts.amazonaws.com",
  ]

  thumbprint_list = ["1b511abead59c6ce207077c0bf0e0043b1382612"]
}
