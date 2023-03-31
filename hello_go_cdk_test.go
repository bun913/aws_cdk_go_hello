package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
	"github.com/bradleyjkemp/cupaloy/v2"
)

// Resource properties check test
func TestHelloGoCdkStack(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewHelloGoCdkStack(app, "MyStack", nil)
	template := assertions.Template_FromStack(stack, nil)
	// template should have a resource of type AWS::SSM::Parameter
	template.HasResourceProperties(jsii.String("AWS::SSM::Parameter"), map[string]interface{}{
		"Name": "HelloFromGoCDK",
	})
}

// Snap Shot test
func TestHelloGoCdkStackSnapshot(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewHelloGoCdkStack(app, "MyStack", nil)
	template := assertions.Template_FromStack(stack, nil)
	cupaloy.SnapshotT(t, template.ToJSON())
}
