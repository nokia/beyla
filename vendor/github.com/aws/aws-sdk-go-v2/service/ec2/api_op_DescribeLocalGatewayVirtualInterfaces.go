// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified local gateway virtual interfaces.
func (c *Client) DescribeLocalGatewayVirtualInterfaces(ctx context.Context, params *DescribeLocalGatewayVirtualInterfacesInput, optFns ...func(*Options)) (*DescribeLocalGatewayVirtualInterfacesOutput, error) {
	if params == nil {
		params = &DescribeLocalGatewayVirtualInterfacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocalGatewayVirtualInterfaces", params, optFns, c.addOperationDescribeLocalGatewayVirtualInterfacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocalGatewayVirtualInterfacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocalGatewayVirtualInterfacesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//   - local-address - The local address.
	//   - local-bgp-asn - The Border Gateway Protocol (BGP) Autonomous System Number
	//   (ASN) of the local gateway.
	//   - local-gateway-id - The ID of the local gateway.
	//   - local-gateway-virtual-interface-id - The ID of the virtual interface.
	//   - owner-id - The ID of the Amazon Web Services account that owns the local
	//   gateway virtual interface.
	//   - peer-address - The peer address.
	//   - peer-bgp-asn - The peer BGP ASN.
	//   - vlan - The ID of the VLAN.
	Filters []types.Filter

	// The IDs of the virtual interfaces.
	LocalGatewayVirtualInterfaceIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeLocalGatewayVirtualInterfacesOutput struct {

	// Information about the virtual interfaces.
	LocalGatewayVirtualInterfaces []types.LocalGatewayVirtualInterface

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocalGatewayVirtualInterfacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeLocalGatewayVirtualInterfaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLocalGatewayVirtualInterfaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLocalGatewayVirtualInterfaces"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocalGatewayVirtualInterfaces(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeLocalGatewayVirtualInterfacesAPIClient is a client that implements the
// DescribeLocalGatewayVirtualInterfaces operation.
type DescribeLocalGatewayVirtualInterfacesAPIClient interface {
	DescribeLocalGatewayVirtualInterfaces(context.Context, *DescribeLocalGatewayVirtualInterfacesInput, ...func(*Options)) (*DescribeLocalGatewayVirtualInterfacesOutput, error)
}

var _ DescribeLocalGatewayVirtualInterfacesAPIClient = (*Client)(nil)

// DescribeLocalGatewayVirtualInterfacesPaginatorOptions is the paginator options
// for DescribeLocalGatewayVirtualInterfaces
type DescribeLocalGatewayVirtualInterfacesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLocalGatewayVirtualInterfacesPaginator is a paginator for
// DescribeLocalGatewayVirtualInterfaces
type DescribeLocalGatewayVirtualInterfacesPaginator struct {
	options   DescribeLocalGatewayVirtualInterfacesPaginatorOptions
	client    DescribeLocalGatewayVirtualInterfacesAPIClient
	params    *DescribeLocalGatewayVirtualInterfacesInput
	nextToken *string
	firstPage bool
}

// NewDescribeLocalGatewayVirtualInterfacesPaginator returns a new
// DescribeLocalGatewayVirtualInterfacesPaginator
func NewDescribeLocalGatewayVirtualInterfacesPaginator(client DescribeLocalGatewayVirtualInterfacesAPIClient, params *DescribeLocalGatewayVirtualInterfacesInput, optFns ...func(*DescribeLocalGatewayVirtualInterfacesPaginatorOptions)) *DescribeLocalGatewayVirtualInterfacesPaginator {
	if params == nil {
		params = &DescribeLocalGatewayVirtualInterfacesInput{}
	}

	options := DescribeLocalGatewayVirtualInterfacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLocalGatewayVirtualInterfacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLocalGatewayVirtualInterfacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLocalGatewayVirtualInterfaces page.
func (p *DescribeLocalGatewayVirtualInterfacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLocalGatewayVirtualInterfacesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeLocalGatewayVirtualInterfaces(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeLocalGatewayVirtualInterfaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLocalGatewayVirtualInterfaces",
	}
}
