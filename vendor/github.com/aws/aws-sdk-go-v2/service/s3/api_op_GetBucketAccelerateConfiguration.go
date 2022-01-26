// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This implementation of the GET action uses the accelerate subresource to return
// the Transfer Acceleration state of a bucket, which is either Enabled or
// Suspended. Amazon S3 Transfer Acceleration is a bucket-level feature that
// enables you to perform faster data transfers to and from Amazon S3. To use this
// operation, you must have permission to perform the s3:GetAccelerateConfiguration
// action. The bucket owner has this permission by default. The bucket owner can
// grant this permission to others. For more information about permissions, see
// Permissions Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html)
// in the Amazon S3 User Guide. You set the Transfer Acceleration state of an
// existing bucket to Enabled or Suspended by using the
// PutBucketAccelerateConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAccelerateConfiguration.html)
// operation. A GET accelerate request does not return a state value for a bucket
// that has no transfer acceleration state. A bucket has no Transfer Acceleration
// state if a state has never been set on the bucket. For more information about
// transfer acceleration, see Transfer Acceleration
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in
// the Amazon S3 User Guide. Related Resources
//
// * PutBucketAccelerateConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAccelerateConfiguration.html)
func (c *Client) GetBucketAccelerateConfiguration(ctx context.Context, params *GetBucketAccelerateConfigurationInput, optFns ...func(*Options)) (*GetBucketAccelerateConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketAccelerateConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketAccelerateConfiguration", params, optFns, c.addOperationGetBucketAccelerateConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketAccelerateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketAccelerateConfigurationInput struct {

	// The name of the bucket for which the accelerate configuration is retrieved.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type GetBucketAccelerateConfigurationOutput struct {

	// The accelerate configuration of the bucket.
	Status types.BucketAccelerateStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketAccelerateConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketAccelerateConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketAccelerateConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketAccelerateConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketAccelerateConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addGetBucketAccelerateConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketAccelerateConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketAccelerateConfiguration",
	}
}

// getGetBucketAccelerateConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getGetBucketAccelerateConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketAccelerateConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetBucketAccelerateConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetBucketAccelerateConfigurationBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
