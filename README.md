# Simple AWS Lambda APIs using Go

This POC was made for testing purposes and uses Localstack as an alternative to AWS live servers. It creates two functions using AWS Lambda and API Gateway integration (POST and GET).

## Before running

Install Localstack (https://github.com/localstack/localstack)
Install Serverless (https://www.serverless.com/)
Install Serverless plugin for Localstack (inside project root folder):

    npm install -D serverless-localstack

Deploy (local stage only):

    make deploy

More about Serverless with Localstack: https://localstack.cloud/docs/how-to/serverless/

Make sure to change LOCALSTACK_API_ID env to match yoyr Localstack (more here https://github.com/localstack/localstack#invoking-api-gateway).
