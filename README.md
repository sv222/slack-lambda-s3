# Upload files to AWS S3 via Slack with Lambda Function

This project contains an AWS Lambda function written in Go that responds to a Slack event API request containing a file upload, copies the file to an S3 bucket, and returns a success message.

## Usage

### 1. Clone the repository

```sh
git clone https://github.com/sv222/slack-lambda-s3.git
cd slack-lambda-s3
```

### 2. Create an S3 bucket

```sh
cd terraform
terraform init
terraform apply \
  -var="bucket_name=my-bucket-name" \
  -var="region=us-west-2"
```

Replace my-bucket-name with the desired name of the S3 bucket, and replace us-west-2 with the desired AWS region.

### 3. Configure the Slack app and AWS Lambda function

1. Create a new Slack app with an event subscription for file_shared.
2. Create a new AWS Lambda function with the slack-s3-lambda.zip deployment package.
3. Set the following environment variables for the Lambda function:

```sh
AWS_REGION: the AWS region in which the S3 bucket was created
SLACK_BOT_TOKEN: the Slack bot token for the app
```

### 4. Deploy the Lambda function

```sh
cd lambda
GOOS=linux go build -o main
zip slack-lambda-s3.zip main
aws lambda update-function-code \
  --function-name=my-lambda-function \
  --zip-file fileb://slack-lambda-s3.zip
```

Replace my-lambda-function with the name of the Lambda function.

### 5. Test the Lambda function

- Upload a file to a Slack channel that the Slack app is a member of.
- Verify that the file was copied to the S3 bucket and that the Lambda function returned a success message.

## Configuration

The Terraform configuration for the S3 bucket is located in the terraform directory. The following variables can be passed to Terraform:

```sh
bucket_name: The name of the S3 bucket to create.
region: The AWS region in which to create the S3 bucket.
```

The AWS Lambda function code is located in the lambda directory. The following environment variables must be set:

```sh
AWS_REGION: The AWS region in which the S3 bucket was created.
SLACK_BOT_TOKEN: The Slack bot token for the app.
```

## Contributing

Feel free to contribute to this project by submitting pull requests or reporting issues.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
