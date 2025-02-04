// Code generated by "go generate". Please don't change this file directly.

package kendra

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// DataSource_ConfluenceBlogToIndexFieldMapping AWS CloudFormation Resource (AWS::Kendra::DataSource.ConfluenceBlogToIndexFieldMapping)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html
type DataSource_ConfluenceBlogToIndexFieldMapping struct {

	// DataSourceFieldName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-datasourcefieldname
	DataSourceFieldName string `json:"DataSourceFieldName"`

	// DateFieldFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-datefieldformat
	DateFieldFormat *string `json:"DateFieldFormat,omitempty"`

	// IndexFieldName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.html#cfn-kendra-datasource-confluenceblogtoindexfieldmapping-indexfieldname
	IndexFieldName string `json:"IndexFieldName"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *DataSource_ConfluenceBlogToIndexFieldMapping) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.ConfluenceBlogToIndexFieldMapping"
}
