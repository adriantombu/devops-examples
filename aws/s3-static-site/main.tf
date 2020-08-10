provider "aws" {
  version = "~> 2.69"
  region  = "eu-west-1"
}

resource "aws_s3_bucket" "b" {
  bucket        = "s3-static-site.devops.otso.fr"
  arn           = "arn:aws:s3:::s3-static-site.devops.otso.fr"
  acl           = "public-read"
  policy        = file("policy.json")
  force_destroy = true

  website {
    index_document = "index.html"
  }
}
