package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type HelloGoCdkStackProps struct {
	awscdk.StackProps
}

func NewHelloGoCdkStack(scope constructs.Construct, id string, props *HelloGoCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	awsssm.NewStringParameter(stack, jsii.String("Hello"), &awsssm.StringParameterProps{
		StringValue:   jsii.String("Hello, Go CDK!!"),
		ParameterName: jsii.String("HelloFromGoCDK"),
	})
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewHelloGoCdkStack(app, "HelloGoCdkStack", &HelloGoCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Region: jsii.String("ap-northeast-1"),
	}
}
