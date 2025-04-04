// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package publish_test

import (
	"context"

	publish "cloud.google.com/go/streetview/publish/apiv1"
	publishpb "cloud.google.com/go/streetview/publish/apiv1/publishpb"
	"google.golang.org/api/iterator"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func ExampleNewStreetViewPublishClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewStreetViewPublishRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleStreetViewPublishClient_BatchDeletePhotos() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.BatchDeletePhotosRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#BatchDeletePhotosRequest.
	}
	resp, err := c.BatchDeletePhotos(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_BatchGetPhotos() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.BatchGetPhotosRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#BatchGetPhotosRequest.
	}
	resp, err := c.BatchGetPhotos(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_BatchUpdatePhotos() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.BatchUpdatePhotosRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#BatchUpdatePhotosRequest.
	}
	resp, err := c.BatchUpdatePhotos(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_CreatePhoto() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.CreatePhotoRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#CreatePhotoRequest.
	}
	resp, err := c.CreatePhoto(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_CreatePhotoSequence() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.CreatePhotoSequenceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#CreatePhotoSequenceRequest.
	}
	op, err := c.CreatePhotoSequence(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_DeletePhoto() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.DeletePhotoRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#DeletePhotoRequest.
	}
	err = c.DeletePhoto(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleStreetViewPublishClient_DeletePhotoSequence() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.DeletePhotoSequenceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#DeletePhotoSequenceRequest.
	}
	err = c.DeletePhotoSequence(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleStreetViewPublishClient_GetPhoto() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.GetPhotoRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#GetPhotoRequest.
	}
	resp, err := c.GetPhoto(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_GetPhotoSequence() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.GetPhotoSequenceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#GetPhotoSequenceRequest.
	}
	op, err := c.GetPhotoSequence(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_ListPhotoSequences() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.ListPhotoSequencesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#ListPhotoSequencesRequest.
	}
	it := c.ListPhotoSequences(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*publishpb.ListPhotoSequencesResponse)
	}
}

func ExampleStreetViewPublishClient_ListPhotos() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.ListPhotosRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#ListPhotosRequest.
	}
	it := c.ListPhotos(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*publishpb.ListPhotosResponse)
	}
}

func ExampleStreetViewPublishClient_StartPhotoSequenceUpload() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/protobuf/types/known/emptypb#Empty.
	}
	resp, err := c.StartPhotoSequenceUpload(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_StartUpload() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/protobuf/types/known/emptypb#Empty.
	}
	resp, err := c.StartUpload(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStreetViewPublishClient_UpdatePhoto() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := publish.NewStreetViewPublishClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &publishpb.UpdatePhotoRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/streetview/publish/apiv1/publishpb#UpdatePhotoRequest.
	}
	resp, err := c.UpdatePhoto(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
