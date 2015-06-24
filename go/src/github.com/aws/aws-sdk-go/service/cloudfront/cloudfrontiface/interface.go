// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudfrontiface provides an interface for the Amazon CloudFront.
package cloudfrontiface

import (
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

// CloudFrontAPI is the interface type for cloudfront.CloudFront.
type CloudFrontAPI interface {
	CreateCloudFrontOriginAccessIdentity(*cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error)

	CreateDistribution(*cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error)

	CreateInvalidation(*cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error)

	CreateStreamingDistribution(*cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error)

	DeleteCloudFrontOriginAccessIdentity(*cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error)

	DeleteDistribution(*cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error)

	DeleteStreamingDistribution(*cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error)

	GetCloudFrontOriginAccessIdentity(*cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)

	GetCloudFrontOriginAccessIdentityConfig(*cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)

	GetDistribution(*cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error)

	GetDistributionConfig(*cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error)

	GetInvalidation(*cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error)

	GetStreamingDistribution(*cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error)

	GetStreamingDistributionConfig(*cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error)

	ListCloudFrontOriginAccessIdentities(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)

	ListDistributions(*cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error)

	ListInvalidations(*cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error)

	ListStreamingDistributions(*cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error)

	UpdateCloudFrontOriginAccessIdentity(*cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error)

	UpdateDistribution(*cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error)

	UpdateStreamingDistribution(*cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error)
}
