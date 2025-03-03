# Pike

![alt text](pike.jfif "Pike")

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/jameswoolfenden/pike/graphs/commit-activity)
[![Build Status](https://github.com/JamesWoolfenden/pike/workflows/CI/badge.svg?branch=master)](https://github.com/JamesWoolfenden/pike)
[![Latest Release](https://img.shields.io/github/release/JamesWoolfenden/pike.svg)](https://github.com/JamesWoolfenden/pike/releases/latest)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/JamesWoolfenden/pike.svg?label=latest)](https://github.com/JamesWoolfenden/pike/releases/latest)
![Terraform Version](https://img.shields.io/badge/tf-%3E%3D0.14.0-blue.svg)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![checkov](https://img.shields.io/badge/checkov-verified-brightgreen)](https://www.checkov.io/)
[![Github All Releases](https://img.shields.io/github/downloads/jameswoolfenden/pike/total.svg)](https://github.com/JamesWoolfenden/pike/releases)

Pike is a tool, to determine the minimum permissions required to run a TF/IAC run:

Pike currently supports Terraform and supports multiple providers (AWS, GCP, AZURE),
Azure is the newest with AWS having the most supported resources <https://github.com/JamesWoolfenden/pike/tree/master/src/mapping>.
Feel free to submit PR or Issue if you find an issue or even better add new resources, and then I'll take a look at merging it ASAP.

**CAVEAT** The policies and roles are to get you started, there are no conditions and resources are all wildcards (for AWS)
this is **definitely not best practice**, you need will to modify these permissions to minimum required by adding these constrictions, however I also added the ability (in AWS so far) to generate short-lived credentials for your build and remotely (REMOTE) supply and invoke your builds (INVOKE).

Ideally I would like to do this for you, but these policies are determined statically, and we would need to determine the resource names that will be created and know your intentions.

## Table of Contents

<!--toc:start-->
- [Pike](#pike)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
    - [MacOS](#macos)
    - [Windows](#windows)
    - [Docker](#docker)
  - [Usage](#usage)
    - [Scan](#scan)
    - [Output](#output)
    - [Make](#make)
    - [Invoke](#invoke)
    - [Apply](#apply)
    - [Remote](#remote)
    - [Readme](#readme)
  - [Compare](#compare)
  - [Help](#help)
  - [Building](#building)
  - [Extending](#extending)
    - [Add Import mapping file](#add-import-mapping-file)
    - [Add to provider Scan](#add-to-provider-scan)
  - [Related Tools](#related-tools)
<!--toc:end-->

## Install

Download the latest binary here:

<https://github.com/JamesWoolfenden/pike/releases>

Install from code:

- Clone repo
- Run `go install`

Install remotely:

```shell
go install  github.com/jameswoolfenden/pike@latest
```

### MacOS

```shell
brew tap jameswoolfenden/homebrew-tap
brew install jameswoolfenden/tap/pike
```

### Windows

```shell
choco install pike
```

### Docker

```shell
docker pull jameswoolfenden/pike
docker run --tty --volume /local/path/to/tf:/tf jameswoolfenden/pike scan -d /tf
```

<https://hub.docker.com/repository/docker/jameswoolfenden/pike>

## Usage

### Scan

To scan a directory of Terraform file:

```shell
./pike scan -d .\terraform\
{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Action": [
            "ec2:MonitorInstances",
            "ec2:UnmonitorInstances",
            "ec2:DescribeInstances",
            "ec2:DescribeTags",
            "ec2:DescribeInstanceAttribute",
            "ec2:DescribeVolumes",
            "ec2:DescribeInstanceTypes",
            "ec2:RunInstances",
            "ec2:DescribeInstanceCreditSpecifications",
            "ec2:StopInstances",
            "ec2:StartInstances",
            "ec2:ModifyInstanceAttribute",
            "ec2:TerminateInstances",
            "ec2:AuthorizeSecurityGroupIngress",
            "ec2:AuthorizeSecurityGroupEgress",
            "ec2:CreateSecurityGroup",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeAccountAttributes",
            "ec2:DescribeNetworkInterfaces",
            "ec2:DeleteSecurityGroup",
            "ec2:RevokeSecurityGroupEgress"
        ],
        "Resource": "*"
    }
}
```

You can also generate the policy as Terraform instead:

```bash
$pike scan -o terraform -d ../modules/aws/terraform-aws-activemq
resource "aws_iam_policy" "terraformXVlBzgba" {
  name        = "terraformXVlBzgba"
  path        = "/"
  description = "Add Description"

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "ec2:AuthorizeSecurityGroupEgress",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CreateNetworkInterface",
                "ec2:CreateNetworkInterfacePermission",
                "ec2:CreateSecurityGroup",
                "ec2:CreateTags",
                "ec2:DeleteNetworkInterface",
                "ec2:DeleteNetworkInterfacePermission",
                "ec2:DeleteSecurityGroup",
                "ec2:DeleteTags",
                "ec2:DescribeAccountAttributes",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeVpcs",
                "ec2:DetachNetworkInterface",
                "ec2:RevokeSecurityGroupEgress",
                "ec2:RevokeSecurityGroupIngress"
            ],
            "Resource": "*"
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": [
                "kms:CreateKey",
                "kms:DescribeKey",
                "kms:EnableKeyRotation",
                "kms:GetKeyPolicy",
                "kms:GetKeyRotationStatus",
                "kms:ListResourceTags",
                "kms:ScheduleKeyDeletion",
                "kms:TagResource",
                "kms:UntagResource"
            ],
            "Resource": "*"
        },
        {
            "Sid": "VisualEditor2",
            "Effect": "Allow",
            "Action": [
                "mq:CreateBroker",
                "mq:CreateConfiguration",
                "mq:CreateTags",
                "mq:CreateUser",
                "mq:DeleteBroker",
                "mq:DeleteTags",
                "mq:DeleteUser",
                "mq:DescribeBroker",
                "mq:DescribeConfiguration",
                "mq:DescribeConfigurationRevision",
                "mq:DescribeUser",
                "mq:RebootBroker",
                "mq:UpdateBroker",
                "mq:UpdateConfiguration",
                "mq:UpdateUser"
            ],
            "Resource": "*"
        }
    ]
})
}
```

### Output

If you select the -w flag, pike will write out the role/policy required to build your project into the .pike folder:

```bash
$pike scan -w -i -d .
2022/09/17 13:50:51 terraform init at .
2022/09/17 13:50:51 downloaded ip
```

The .pike folder will contain:

``` shell
aws_iam_role.terraform_pike.tf
pike.generated_policy.tf
```

Which you can deploy using terraform to create the role/policy to build your IAC project.

### Make

You can now deploy the policy you need directly (AWS only so far):

```bash
$pike make -d ../modules/aws/terraform-aws-apigateway/

2022/09/18 08:53:41 terraform init at ..\modules\aws\terraform-aws-apigateway\
2022/09/18 08:53:41 modules not found at ..\modules\aws\terraform-aws-apigateway\
2022/09/18 08:53:49 aws role create/updated arn:aws:iam::680235478471:role/terraform_pike_20220918071439382800000002
 arn:aws:iam::680235478471:role/terraform_pike_20220918071439382800000002
```

This new verb returns the ARN of the role created, and you can find the Terraform used in your .pike folder.

### Invoke

Invoke is currently for triggering GitHub actions, if supplied with the workflow (defaults to main.yaml), repository and
branch (defaults to main) flags, it will trigger the dispatch event.

You'll need to include the dispatch event in your workflow:

```yaml
on:
  workflow_dispatch:
  push:
    branches:
      - master
```

To authenticate the GitHub Api you will need to set you GitHub Personal Access Token as the environment variable
*GITHUB_TOKEN*

To Invoke a workflow it is then:

```shell
pike invoke -workflow master.yml -branch master -repository JamesWoolfenden/terraform-aws-s3
```

I created Invoke to be used in tandem with the new remote command which supplies temporary credentials to a workflow.

**Note The gitHub API is rate limited usually 5000 calls per hour.

```shell
pike make -d ./module/aws/terraform-aws-s3/example/examplea
```

### Apply

Apply is an extension to make and will apply the policy and role and use that role to create your infrastructure:

```shell
pike apply -d ./module/aws/terraform-aws-s3/example/examplea -region eu-west-2
```

It is intended for testing and developing the permissions for Pike itself

### Remote

Remote uses the core code of make and apply, to write temporary AWS credentials(only so far) into your workflow.

```shell
pike remote -d ./module/aws/terraform-aws-s3/example/examplea -region eu-west-2 -repository terraform-aws-s3
```

### Readme

Pike can now be used to update a projects README.md file:

./pike readme -o terraform -d ..\modules\aws\terraform-aws-activemq\

This looks in the readme for the deliminators:

```html
<!-- BEGINNING OF PRE-COMMIT-PIKE DOCS HOOK -->
<!-- END OF PRE-COMMIT-PIKE DOCS HOOK -->
```

and replaces is either with json or Terraform like so:

```markdown
This is the policy required to build this project:

The Policy required is

{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Action": [
            "mq:CreateTags",
            "mq:DeleteTags",
            "ec2:DescribeInternetGateways",
            "ec2:DescribeAccountAttributes",
            "ec2:DescribeVpcs",
            "ec2:DescribeSubnets",
            "ec2:DescribeSecurityGroups",
            "ec2:CreateNetworkInterface",
            "ec2:CreateNetworkInterfacePermission",
            "ec2:DeleteNetworkInterfacePermission",
            "ec2:DetachNetworkInterface",
            "ec2:DeleteNetworkInterface",
            "mq:CreateBroker",
            "mq:DescribeBroker",
            "mq:DescribeUser",
            "mq:UpdateBroker",
            "mq:DeleteBroker",
            "mq:CreateConfiguration",
            "mq:UpdateConfiguration",
            "mq:DescribeConfiguration",
            "mq:DescribeConfigurationRevision",
            "mq:RebootBroker",
            "ec2:CreateTags",
            "ec2:DeleteTags",
            "ec2:CreateSecurityGroup",
            "ec2:DescribeNetworkInterfaces",
            "ec2:DeleteSecurityGroup",
            "ec2:RevokeSecurityGroupEgress",
            "kms:TagResource",
            "kms:UntagResource",
            "kms:EnableKeyRotation",
            "kms:CreateKey",
            "kms:DescribeKey",
            "kms:GetKeyPolicy",
            "kms:GetKeyRotationStatus",
            "kms:ListResourceTags",
            "kms:ScheduleKeyDeletion"
        ],
        "Resource": "*"
    }
}
```

You can see an example here <https://github.com/jamesWoolfenden/terraform-aws-activemq#policy>.

## Compare

Want to check your deployed IAM policy against your IAC requirement?

>$./pike compare -d ../modules/aws/terraform-aws-appsync -a arn:aws:iam::680235478471:policy/basic

```markdown
IAM Policy arn:aws:iam::680235478471:policy/basic versus Local ../modules/aws/terraform-aws-appsync
 {
   "Statement": [
     0: {
       "Action": [
-        0: "kinesisvideo:CreateStream"
+        0: "firehose:CreateDeliveryStream"
+        0: "firehose:CreateDeliveryStream"
+        1: "firehose:DeleteDeliveryStream"
+        2: "firehose:DescribeDeliveryStream"
+        3: "firehose:ListTagsForDeliveryStream"
+        4: "iam:AttachRolePolicy"
+        5: "iam:CreateRole"
+        6: "iam:DeleteRole"
+        7: "iam:DetachRolePolicy"
+        8: "iam:GetRole"
+        9: "iam:ListAttachedRolePolicies"
+        10: "iam:ListInstanceProfilesForRole"
+        11: "iam:ListRolePolicies"
+        12: "iam:PassRole"
+        13: "iam:TagRole"
+        14: "kms:CreateKey"
+        15: "kms:DescribeKey"
+        16: "kms:EnableKeyRotation"
+        17: "kms:GetKeyPolicy"
+        18: "kms:GetKeyRotationStatus"
+        19: "kms:ListResourceTags"
+        20: "kms:ScheduleKeyDeletion"
+        21: "logs:AssociateKmsKey"
+        22: "logs:CreateLogGroup"
+        23: "logs:DeleteLogGroup"
+        24: "logs:DeleteRetentionPolicy"
+        25: "logs:DescribeLogGroups"
+        26: "logs:DisassociateKmsKey"
+        27: "logs:ListTagsLogGroup"
+        28: "logs:PutRetentionPolicy"
+        29: "s3:CreateBucket"
+        30: "s3:DeleteBucket"
+        31: "s3:GetAccelerateConfiguration"
+        32: "s3:GetBucketAcl"
+        33: "s3:GetBucketCORS"
+        34: "s3:GetBucketLogging"
+        35: "s3:GetBucketObjectLockConfiguration"
+        36: "s3:GetBucketPolicy"
+        37: "s3:GetBucketPublicAccessBlock"
+        38: "s3:GetBucketRequestPayment"
+        39: "s3:GetBucketTagging"
+        40: "s3:GetBucketVersioning"
+        41: "s3:GetBucketWebsite"
+        42: "s3:GetEncryptionConfiguration"
+        43: "s3:GetLifecycleConfiguration"
+        44: "s3:GetObject"
+        45: "s3:GetObjectAcl"
+        46: "s3:GetReplicationConfiguration"
+        47: "s3:ListAllMyBuckets"
+        48: "s3:ListBucket"
+        49: "s3:PutBucketAcl"
+        50: "s3:PutBucketPublicAccessBlock"
+        51: "s3:PutEncryptionConfiguration"
+        52: "wafv2:CreateWebACL"
+        53: "wafv2:DeleteWebACL"
+        54: "wafv2:GetWebACL"
       ],
       "Effect": "Allow",
       "Resource": "*",
-      "Sid": ""
+      "Sid": "VisualEditor0"
     }
   ],
   "Version": "2012-10-17"
 }

```

## Help

```bash
./pike -h
NAME:
   pike - Generate IAM policy from your IAC code

USAGE:
   pike [global options] command [command options] [arguments...]

VERSION:
   v0.2.1

AUTHOR:
   James Woolfenden <support@bridgecrew.io>

COMMANDS:
   apply, a    Create a policy and use it to instantiate the IAC
   compare, c  policy comparison of deployed versus IAC
   invoke, i   Triggers a gitHub action specified with the workflow flag
   make, m     make the policy/role required for this IAC to deploy
   readme, r   Looks in dir for a README.md and updates it with the Policy required to build the code
   remote, m   Create/Update the Policy and set credentials/secret for Github Action
   scan, s     scan a directory for IAM code
   version, v  Outputs the application version
   watch, w    Waits for policy update
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Building

```go
go build
```

or

```Make
Make build
```

## Extending

Determine and Create IAM mapping file,
working out the permissions required for your resource:
e.g. *aws_security_group.json*

```json
[
  {
    "apply": [
      "ec2:CreateSecurityGroup",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeAccountAttributes",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DeleteSecurityGroup",
      "ec2:RevokeSecurityGroupEgress"
    ],
    "attributes": {
      "ingress": [
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:AuthorizeSecurityGroupEgress"
      ],
      "tags": [
        "ec2:CreateTags",
        "ec2:DeleteTags"
      ]
    },
    "destroy": [
      "ec2:DeleteSecurityGroup"
    ],
    "modify": [],
    "plan": []
  }
]

```

### Add Import mapping file

Update **files.go** with:

```go
//go:embed aws_security_group.json
var securityGroup []byte
```

### Add to provider Scan

Once you have added the json import above you just need to update the lookup table,
so we can read it and get the permissions:

```go
func GetAWSResourcePermissions(result template) []interface{} {
    TFLookup := map[string]interface{}{
        "aws_s3_bucket":            awsS3Bucket,
        "aws_s3_bucket_acl":        awsS3BucketACL,
+         "aws_security_group":       awsSecurityGroup,

```

Also add an example Terraform file into the folder terraform/backups.

## Related Tools

<https://github.com/iann0036/iamlive>
