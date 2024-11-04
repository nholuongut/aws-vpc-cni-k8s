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
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut/session"
	"github.com/nholuongut/nholuongut-sdk-go/service/iam"
	"github.com/nholuongut/nholuongut-sdk-go/service/iam/iamiface"
)

type PolicyDocument struct {
	Version   string
	Statement []StatementEntry
}

type StatementEntry struct {
	Effect   string
	Action   []string
	Resource string
}

type IAM interface {
	AttachRolePolicy(policyArn string, roleName string) error
	DetachRolePolicy(policyARN string, roleName string) error
	CreatePolicy(policyName string, policyDocument string) (*iam.CreatePolicyOutput, error)
	DeletePolicy(policyARN string) error
	GetInstanceProfile(instanceProfileName string) (*iam.GetInstanceProfileOutput, error)
	GetRolePolicy(policyName string, role string) (*iam.GetRolePolicyOutput, error)
	PutRolePolicy(policyDocument string, policyName string, roleName string) error
	ListPolicies(scope string) (*iam.ListPoliciesOutput, error)
}

type defaultIAM struct {
	iamiface.IAMAPI
}

func (d *defaultIAM) AttachRolePolicy(policyARN string, roleName string) error {
	attachRolePolicyInput := &iam.AttachRolePolicyInput{
		PolicyArn: nholuongut.String(policyARN),
		RoleName:  nholuongut.String(roleName),
	}
	_, err := d.IAMAPI.AttachRolePolicy(attachRolePolicyInput)
	return err
}

func (d *defaultIAM) DetachRolePolicy(policyARN string, roleName string) error {
	detachRolePolicyInput := &iam.DetachRolePolicyInput{
		PolicyArn: nholuongut.String(policyARN),
		RoleName:  nholuongut.String(roleName),
	}
	_, err := d.IAMAPI.DetachRolePolicy(detachRolePolicyInput)
	return err
}

func (d *defaultIAM) CreatePolicy(policyName string, policyDocument string) (*iam.CreatePolicyOutput, error) {
	createPolicyInput := &iam.CreatePolicyInput{
		PolicyDocument: nholuongut.String(policyDocument),
		PolicyName:     nholuongut.String(policyName),
	}
	return d.IAMAPI.CreatePolicy(createPolicyInput)
}

func (d *defaultIAM) DeletePolicy(policyARN string) error {
	deletePolicyInput := &iam.DeletePolicyInput{
		PolicyArn: nholuongut.String(policyARN),
	}
	_, err := d.IAMAPI.DeletePolicy(deletePolicyInput)
	return err
}

func (d *defaultIAM) GetRolePolicy(role string, policyName string) (*iam.GetRolePolicyOutput, error) {
	rolePolicyInput := &iam.GetRolePolicyInput{
		RoleName:   nholuongut.String(role),
		PolicyName: nholuongut.String(policyName),
	}
	return d.IAMAPI.GetRolePolicy(rolePolicyInput)
}

func (d *defaultIAM) PutRolePolicy(policyDocument string, policyName string, roleName string) error {
	policyInput := &iam.PutRolePolicyInput{
		PolicyDocument: nholuongut.String(policyDocument),
		PolicyName:     nholuongut.String(policyName),
		RoleName:       nholuongut.String(roleName),
	}
	_, err := d.IAMAPI.PutRolePolicy(policyInput)
	return err
}

func (d *defaultIAM) GetInstanceProfile(instanceProfileName string) (*iam.GetInstanceProfileOutput, error) {
	getInstanceProfileInput := &iam.GetInstanceProfileInput{
		InstanceProfileName: nholuongut.String(instanceProfileName),
	}
	return d.IAMAPI.GetInstanceProfile(getInstanceProfileInput)
}

func (d *defaultIAM) ListPolicies(scope string) (*iam.ListPoliciesOutput, error) {
	listPolicyInput := &iam.ListPoliciesInput{
		Scope: nholuongut.String(scope),
	}
	return d.IAMAPI.ListPolicies(listPolicyInput)
}

func NewIAM(session *session.Session) IAM {
	return &defaultIAM{
		IAMAPI: iam.New(session),
	}
}
