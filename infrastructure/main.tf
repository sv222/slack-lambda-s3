provider "aws" {
  region = var.region
}

resource "aws_s3_bucket" "file_storage" {
  bucket = var.bucket_name
  acl    = "private"

  tags = {
    Name = var.bucket_name
  }
}

output "bucket_name" {
  value = aws_s3_bucket.file_storage.id
}
