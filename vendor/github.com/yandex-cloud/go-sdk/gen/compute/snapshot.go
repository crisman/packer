// Code generated by sdkgen. DO NOT EDIT.

//nolint
package compute

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// SnapshotServiceClient is a compute.SnapshotServiceClient with
// lazy GRPC connection initialization.
type SnapshotServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements compute.SnapshotServiceClient
func (c *SnapshotServiceClient) Create(ctx context.Context, in *compute.CreateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements compute.SnapshotServiceClient
func (c *SnapshotServiceClient) Delete(ctx context.Context, in *compute.DeleteSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements compute.SnapshotServiceClient
func (c *SnapshotServiceClient) Get(ctx context.Context, in *compute.GetSnapshotRequest, opts ...grpc.CallOption) (*compute.Snapshot, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotServiceClient(conn).Get(ctx, in, opts...)
}

// List implements compute.SnapshotServiceClient
func (c *SnapshotServiceClient) List(ctx context.Context, in *compute.ListSnapshotsRequest, opts ...grpc.CallOption) (*compute.ListSnapshotsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotServiceClient(conn).List(ctx, in, opts...)
}

type SnapshotIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *SnapshotServiceClient
	request *compute.ListSnapshotsRequest

	items []*compute.Snapshot
}

func (c *SnapshotServiceClient) SnapshotIterator(ctx context.Context, folderId string, opts ...grpc.CallOption) *SnapshotIterator {
	return &SnapshotIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &compute.ListSnapshotsRequest{
			FolderId: folderId,
			PageSize: 1000,
		},
	}
}

func (it *SnapshotIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Snapshots
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SnapshotIterator) Value() *compute.Snapshot {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SnapshotIterator) Error() error {
	return it.err
}

// ListOperations implements compute.SnapshotServiceClient
func (c *SnapshotServiceClient) ListOperations(ctx context.Context, in *compute.ListSnapshotOperationsRequest, opts ...grpc.CallOption) (*compute.ListSnapshotOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotServiceClient(conn).ListOperations(ctx, in, opts...)
}

type SnapshotOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *SnapshotServiceClient
	request *compute.ListSnapshotOperationsRequest

	items []*operation.Operation
}

func (c *SnapshotServiceClient) SnapshotOperationsIterator(ctx context.Context, snapshotId string, opts ...grpc.CallOption) *SnapshotOperationsIterator {
	return &SnapshotOperationsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &compute.ListSnapshotOperationsRequest{
			SnapshotId: snapshotId,
			PageSize:   1000,
		},
	}
}

func (it *SnapshotOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SnapshotOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SnapshotOperationsIterator) Error() error {
	return it.err
}

// Update implements compute.SnapshotServiceClient
func (c *SnapshotServiceClient) Update(ctx context.Context, in *compute.UpdateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewSnapshotServiceClient(conn).Update(ctx, in, opts...)
}
