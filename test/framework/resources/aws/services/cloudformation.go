// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://nholuongut.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package services

import (
	"context"
	"fmt"

	"github.com/nholuongut/amazon-vpc-cni-k8s/test/framework/utils"
	"github.com/nholuongut/nholuongut-sdk-go/service/cloudformation/cloudformationiface"

	"github.com/nholuongut/nholuongut-sdk-go/nholuongut"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut/session"
	"github.com/nholuongut/nholuongut-sdk-go/service/cloudformation"
	"k8s.io/apimachinery/pkg/util/wait"
)

type CloudFormation interface {
	WaitTillStackCreated(stackName string, stackParams []*cloudformation.Parameter, templateBody string) (*cloudformation.DescribeStacksOutput, error)
	WaitTillStackDeleted(stackName string) error
}

type defaultCloudFormation struct {
	cloudformationiface.CloudFormationAPI
}

func NewCloudFormation(session *session.Session) CloudFormation {
	return &defaultCloudFormation{
		CloudFormationAPI: cloudformation.New(session),
	}
}

func (d *defaultCloudFormation) WaitTillStackCreated(stackName string, stackParams []*cloudformation.Parameter, templateBody string) (*cloudformation.DescribeStacksOutput, error) {
	createStackInput := &cloudformation.CreateStackInput{
		Parameters:   stackParams,
		StackName:    nholuongut.String(stackName),
		TemplateBody: nholuongut.String(templateBody),
		Capabilities: nholuongut.StringSlice([]string{cloudformation.CapabilityCapabilityIam}),
	}

	_, err := d.CloudFormationAPI.CreateStack(createStackInput)
	if err != nil {
		return nil, err
	}

	describeStackInput := &cloudformation.DescribeStacksInput{
		StackName: nholuongut.String(stackName),
	}

	var describeStackOutput *cloudformation.DescribeStacksOutput
	err = wait.PollImmediateUntil(utils.PollIntervalLong, func() (done bool, err error) {
		describeStackOutput, err = d.CloudFormationAPI.DescribeStacks(describeStackInput)
		if err != nil {
			return true, err
		}
		if *describeStackOutput.Stacks[0].StackStatus == "CREATE_COMPLETE" {
			return true, nil
		}
		return false, nil
	}, context.Background().Done())

	return describeStackOutput, err
}

func (d *defaultCloudFormation) WaitTillStackDeleted(stackName string) error {
	deleteStackInput := &cloudformation.DeleteStackInput{
		StackName: nholuongut.String(stackName),
	}
	_, err := d.CloudFormationAPI.DeleteStack(deleteStackInput)
	if err != nil {
		return fmt.Errorf("failed to delete stack %s: %v", stackName, err)
	}

	describeStackInput := &cloudformation.DescribeStacksInput{
		StackName: nholuongut.String(stackName),
	}

	var describeStackOutput *cloudformation.DescribeStacksOutput
	err = wait.PollImmediateUntil(utils.PollIntervalLong, func() (done bool, err error) {
		describeStackOutput, err = d.CloudFormationAPI.DescribeStacks(describeStackInput)
		if err != nil {
			return true, err
		}
		if *describeStackOutput.Stacks[0].StackStatus == "DELETE_COMPLETE" {
			return true, nil
		}
		return false, nil
	}, context.Background().Done())

	return nil
}
