// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package lambda

const (

	// ErrCodeCodeStorageExceededException for service response error code
	// "CodeStorageExceededException".
	//
	// You have exceeded your maximum total code size per account. Limits (http://docs.aws.amazon.com/lambda/latest/dg/limits.html)
	ErrCodeCodeStorageExceededException = "CodeStorageExceededException"

	// ErrCodeEC2AccessDeniedException for service response error code
	// "EC2AccessDeniedException".
	ErrCodeEC2AccessDeniedException = "EC2AccessDeniedException"

	// ErrCodeEC2ThrottledException for service response error code
	// "EC2ThrottledException".
	//
	// AWS Lambda was throttled by Amazon EC2 during Lambda function initialization
	// using the execution role provided for the Lambda function.
	ErrCodeEC2ThrottledException = "EC2ThrottledException"

	// ErrCodeEC2UnexpectedException for service response error code
	// "EC2UnexpectedException".
	//
	// AWS Lambda received an unexpected EC2 client exception while setting up for
	// the Lambda function.
	ErrCodeEC2UnexpectedException = "EC2UnexpectedException"

	// ErrCodeENILimitReachedException for service response error code
	// "ENILimitReachedException".
	//
	// AWS Lambda was not able to create an Elastic Network Interface (ENI) in the
	// VPC, specified as part of Lambda function configuration, because the limit
	// for network interfaces has been reached.
	ErrCodeENILimitReachedException = "ENILimitReachedException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// One of the parameters in the request is invalid. For example, if you provided
	// an IAM role for AWS Lambda to assume in the CreateFunction or the UpdateFunctionConfiguration
	// API, that AWS Lambda is unable to assume you will get this exception.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidRequestContentException for service response error code
	// "InvalidRequestContentException".
	//
	// The request body could not be parsed as JSON.
	ErrCodeInvalidRequestContentException = "InvalidRequestContentException"

	// ErrCodeInvalidSecurityGroupIDException for service response error code
	// "InvalidSecurityGroupIDException".
	//
	// The Security Group ID provided in the Lambda function VPC configuration is
	// invalid.
	ErrCodeInvalidSecurityGroupIDException = "InvalidSecurityGroupIDException"

	// ErrCodeInvalidSubnetIDException for service response error code
	// "InvalidSubnetIDException".
	//
	// The Subnet ID provided in the Lambda function VPC configuration is invalid.
	ErrCodeInvalidSubnetIDException = "InvalidSubnetIDException"

	// ErrCodeInvalidZipFileException for service response error code
	// "InvalidZipFileException".
	//
	// AWS Lambda could not unzip the function zip file.
	ErrCodeInvalidZipFileException = "InvalidZipFileException"

	// ErrCodeKMSAccessDeniedException for service response error code
	// "KMSAccessDeniedException".
	//
	// Lambda was unable to decrypt the environment variables because KMS access
	// was denied. Check the Lambda function's KMS permissions.
	ErrCodeKMSAccessDeniedException = "KMSAccessDeniedException"

	// ErrCodeKMSDisabledException for service response error code
	// "KMSDisabledException".
	//
	// Lambda was unable to decrypt the environment variables because the KMS key
	// used is disabled. Check the Lambda function's KMS key settings.
	ErrCodeKMSDisabledException = "KMSDisabledException"

	// ErrCodeKMSInvalidStateException for service response error code
	// "KMSInvalidStateException".
	//
	// Lambda was unable to decrypt the environment variables because the KMS key
	// used is in an invalid state for Decrypt. Check the function's KMS key settings.
	ErrCodeKMSInvalidStateException = "KMSInvalidStateException"

	// ErrCodeKMSNotFoundException for service response error code
	// "KMSNotFoundException".
	//
	// Lambda was unable to decrypt the environment variables because the KMS key
	// was not found. Check the function's KMS key settings.
	ErrCodeKMSNotFoundException = "KMSNotFoundException"

	// ErrCodePolicyLengthExceededException for service response error code
	// "PolicyLengthExceededException".
	//
	// Lambda function access policy is limited to 20 KB.
	ErrCodePolicyLengthExceededException = "PolicyLengthExceededException"

	// ErrCodeRequestTooLargeException for service response error code
	// "RequestTooLargeException".
	//
	// The request payload exceeded the Invoke request body JSON input limit. For
	// more information, see Limits (http://docs.aws.amazon.com/lambda/latest/dg/limits.html).
	ErrCodeRequestTooLargeException = "RequestTooLargeException"

	// ErrCodeResourceConflictException for service response error code
	// "ResourceConflictException".
	//
	// The resource already exists.
	ErrCodeResourceConflictException = "ResourceConflictException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource (for example, a Lambda function or access policy statement)
	// specified in the request does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceException for service response error code
	// "ServiceException".
	//
	// The AWS Lambda service encountered an internal error.
	ErrCodeServiceException = "ServiceException"

	// ErrCodeSubnetIPAddressLimitReachedException for service response error code
	// "SubnetIPAddressLimitReachedException".
	//
	// AWS Lambda was not able to set up VPC access for the Lambda function because
	// one or more configured subnets has no available IP addresses.
	ErrCodeSubnetIPAddressLimitReachedException = "SubnetIPAddressLimitReachedException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnsupportedMediaTypeException for service response error code
	// "UnsupportedMediaTypeException".
	//
	// The content type of the Invoke request body is not JSON.
	ErrCodeUnsupportedMediaTypeException = "UnsupportedMediaTypeException"
)
