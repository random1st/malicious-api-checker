# Malicious API checker

Check a curated list of URLs(domain names) from DynamoDB table in order to notify about security problems with them.


## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Setup process
#### Credentials
* [Safe Browsing API](https://developers.google.com/safe-browsing/v4/get-started)
* [WebRisk API](https://cloud.google.com/web-risk/docs/quickstart)
* [Slack WebHook URL](https://api.slack.com/messaging/webhooks)

### Local pre-requirements
Create file `env.json` like this
```json
{
  "MaliciousAPIChecker": {
    "GOOGLE_WEBRISK_API_KEY": "<api_key>",
    "GOOGLE_SAFEBROWSING_API_KEY": "<api_key>",
    "DYNAMODB_TABLE": "<table name>",
    "SLACK_WEBHOOK_URL": "<webjook url>"
  }
}
```
You can skip setup either WebRisk or SafeBrowsing API in order to use only one of them



### Installing dependencies & building the target 

In this example we use the built-in `sam build` to automatically download all the dependencies and package our build target.   
Read more about [SAM Build here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-build.html) 

The `sam build` command is wrapped inside of the `Makefile`. To execute this simply run
 
```shell
make
```

### Local development

**Invoking function locally through local API Gateway**

```bash
make invoke
```



## Packaging and deployment

AWS Lambda Golang runtime requires a flat folder with the executable generated on build step. SAM will use `CodeUri` property to know where to look up for the application:

To deploy your application for the first time, run the following in your shell:

```bash
make deploy
```

The command will package and deploy your application to AWS, with a series of prompts:

* **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
* **AWS Region**: The AWS region you want to deploy your app to.
* **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
* **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required permissions. To deploy an AWS CloudFormation stack which creates or modifies IAM roles, the `CAPABILITY_IAM` value for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
* **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

###ATTENTION! After first do not apply changes. Add to `samconfig.toml` parameters

```
parameter_overrides = "GoogleWebriskApiKey=\"<api_key>\" GoogleSafebrowsingApiKey=\"<api_key>\" DynamoDbTable=\"<table_name>\" SlackWebhookUrl=\"<webhook_url>\""
```

# Appendix

### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

A quickstart way would be to use Homebrew, chocolatey or your linux package manager.

#### Homebrew (Mac)

Issue the following command from the terminal:

```shell
brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
brew update
brew upgrade golang
```

#### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
choco upgrade golang
```

## Bringing to the next level

Here are a few ideas that you can use to get more acquainted as to how this overall process works:

* Create an additional API resource (e.g. /hello/{proxy+}) and return the name requested through this new path
* Update unit test to capture that
* Package & Deploy

Next, you can use the following resources to know more about beyond hello world samples and how others structure their Serverless applications:

* [AWS Serverless Application Repository](https://aws.amazon.com/serverless/serverlessrepo/)
