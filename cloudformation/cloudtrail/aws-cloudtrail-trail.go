// Code generated by "go generate". Please don't change this file directly.

package cloudtrail

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v6/cloudformation/policies"
	"github.com/awslabs/goformation/v6/cloudformation/tags"
)

// Trail AWS CloudFormation Resource (AWS::CloudTrail::Trail)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
type Trail struct {

	// CloudWatchLogsLogGroupArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsloggrouparn
	CloudWatchLogsLogGroupArn *string `json:"CloudWatchLogsLogGroupArn,omitempty"`

	// CloudWatchLogsRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsrolearn
	CloudWatchLogsRoleArn *string `json:"CloudWatchLogsRoleArn,omitempty"`

	// EnableLogFileValidation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-enablelogfilevalidation
	EnableLogFileValidation *bool `json:"EnableLogFileValidation,omitempty"`

	// EventSelectors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-eventselectors
	EventSelectors *[]Trail_EventSelector `json:"EventSelectors,omitempty"`

	// IncludeGlobalServiceEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-includeglobalserviceevents
	IncludeGlobalServiceEvents *bool `json:"IncludeGlobalServiceEvents,omitempty"`

	// InsightSelectors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-insightselectors
	InsightSelectors *[]Trail_InsightSelector `json:"InsightSelectors,omitempty"`

	// IsLogging AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-islogging
	IsLogging bool `json:"IsLogging"`

	// IsMultiRegionTrail AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-ismultiregiontrail
	IsMultiRegionTrail *bool `json:"IsMultiRegionTrail,omitempty"`

	// IsOrganizationTrail AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-isorganizationtrail
	IsOrganizationTrail *bool `json:"IsOrganizationTrail,omitempty"`

	// KMSKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-kmskeyid
	KMSKeyId *string `json:"KMSKeyId,omitempty"`

	// S3BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3bucketname
	S3BucketName string `json:"S3BucketName"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3keyprefix
	S3KeyPrefix *string `json:"S3KeyPrefix,omitempty"`

	// SnsTopicName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-snstopicname
	SnsTopicName *string `json:"SnsTopicName,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-tags
	Tags *[]tags.Tag `json:"Tags,omitempty"`

	// TrailName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-trailname
	TrailName *string `json:"TrailName,omitempty"`

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
func (r *Trail) AWSCloudFormationType() string {
	return "AWS::CloudTrail::Trail"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r Trail) MarshalJSON() ([]byte, error) {
	type Properties Trail
	return json.Marshal(&struct {
		Type                string
		Properties          Properties
		DependsOn           []string                     `json:"DependsOn,omitempty"`
		Metadata            map[string]interface{}       `json:"Metadata,omitempty"`
		DeletionPolicy      policies.DeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy policies.UpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Condition           string                       `json:"Condition,omitempty"`
	}{
		Type:                r.AWSCloudFormationType(),
		Properties:          (Properties)(r),
		DependsOn:           r.AWSCloudFormationDependsOn,
		Metadata:            r.AWSCloudFormationMetadata,
		DeletionPolicy:      r.AWSCloudFormationDeletionPolicy,
		UpdateReplacePolicy: r.AWSCloudFormationUpdateReplacePolicy,
		Condition:           r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *Trail) UnmarshalJSON(b []byte) error {
	type Properties Trail
	res := &struct {
		Type                string
		Properties          *Properties
		DependsOn           interface{}
		Metadata            map[string]interface{}
		DeletionPolicy      string
		UpdateReplacePolicy string
		Condition           string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = Trail(*res.Properties)
	}
	if res.DependsOn != nil {
		switch obj := res.DependsOn.(type) {
		case string:
			r.AWSCloudFormationDependsOn = []string{obj}
		case []interface{}:
			s := make([]string, 0, len(obj))
			for _, v := range obj {
				if value, ok := v.(string); ok {
					s = append(s, value)
				}
			}
			r.AWSCloudFormationDependsOn = s
		}
	}
	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.UpdateReplacePolicy != "" {
		r.AWSCloudFormationUpdateReplacePolicy = policies.UpdateReplacePolicy(res.UpdateReplacePolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}
