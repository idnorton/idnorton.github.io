---
layout: default
title:  "Terraforming your AWS parent account"
---
In the modern AWS world, it's not uncommon to have multiple AWS accounts
under a parent billing account for different environments such as dev, qa
and production. Let's look at how we get this started.

All of these steps are assuming Terraform v12.

## Bootstrap

The following items are initial bootstrap and don't rely on shared state
or DynamoDB for locking. At this point we are assuming that the bootstrap
steps are being followed by one person and nobody else will be working on
the account at the same time.

### IAM user and key

Before we can begin working with the account, we need to manually create
an IAM user and key pair. Allocate the IAM user the policy
`arn:aws:iam::aws:policy/AdministratorAccess` which will give us access to
all resources for the next steps.

This may be your own user, but from here on in we assume that the
credentials are in your AWS credentials file with the profile name
`parent`.

### S3 Bucket for state

If you're in a shared environment with more than a single user (which is
likely if you're looking at creating multiple accounts) now is the time to
create an S3 bucket for us to store shared state in.

```
module "shared_state" {
  source = "terraform-aws-modules/s3-bucket/aws"
  bucket = "shared-state"

  acl                     = "private"
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm     = "aws:kms"
      }
    }
  }

  versioning = {
    enabled = true
  }
}
```

At this point we are assuming the user account being used has admin rights
so we're not going to talk about permissions on the bucket. By specifying
the sse_algorithm but not the `kms_master_key_id` we get the default AWS
key `aws/s3` is used which may not be the key you wish to use in production.

### DynamoDB

Next we need a DynamoDB for tracking locks and making sure we don't have
multiple people working on the same things at the same times.

```
module "dynamodb_table" {
  source   = "terraform-aws-modules/dynamodb-table/aws"

  name     = "terraform-locks"
  hash_key = "LockID"

  read_capacity  = 20
  write_capacity = 20

  attributes = [
    {
      name = "LockID"
      type = "S"
    }
  ]
}
```

At this point we are assuming the user account being used has admin rights
so we're not going to talk about permissions on the DynamoDB.

## See also

* [https://www.terraform.io/docs/backends/types/s3.html](https://www.terraform.io/docs/backends/types/s3.html)
