# OurBranch::Custom::SSMParameterValue

An example of how to build a CloudFormation resource provider in Go.

This resource lets you refer to the latest value of an SSM parameter in your templates:

```yml
Resources:
  ESHostParameterValue:
    Type: OurBranch::Custom::SSMParameterValue
    Properties:
      ParameterName: es_host

  SomeLambda:
    Type: AWS::Serverless::Function
    Properties:
      Environment:
        Variables:
          ES_HOST: !GetAtt ESHostParameterValue.Value
```

This is just an example, don't use this in your apps. If you want to use SSM parameters in your production templates use [CloudFormation's built-in support for them](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html#aws-ssm-parameter-types).

## How to deploy

- Set up your environment:
  - Python 3.6 or above
  - Go 1.11 or above
  - AWS CLI

* Create a Python virtualenv for this project (not mandatory but definitely recommended):
  - `python3 -m venv pyvenv`
  - `source pyvenv/bin/activate`

- Install the Python dependencies (CloudFormation CLI and CloudFormation CLI Go plugin)
  - `pip install -r requirements.txt`

* `make clean`
* `cfn validate` to validate the schema.
* `cfn generate` to generate the Go files needed to build the provider from the schema.
* `make build` to build the provider code.
* `cfn submit -v` to register the resource in your CloudFormation registry.

If you update the provider, submit it again, and then set the default version to the new version that you just submitted: `aws cloudformation set-type-default-version --type RESOURCE --type-name "OurBranch::Custom::ParameterValue" --version-id "00000002"`

You can see what versions you have with `aws cloudformation list-type-versions --type "RESOURCE" --type-name "OurBranch::Custom::SSMParameterValue"`.

## More info

- [AWS documentation on creating custom resource providers](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html)
- [CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli)
- [CloudFormation CLI Go Plugin (includes another example of a GitHub repository CloudFormation resource)](https://github.com/aws-cloudformation/cloudformation-cli-go-plugin/)
