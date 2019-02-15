package job

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// PipelineClient is the creates an Azure Data Lake Analytics job client.
type PipelineClient struct {
	BaseClient
}

// NewPipelineClient creates an instance of the PipelineClient client.
func NewPipelineClient() PipelineClient {
	return PipelineClient{New()}
}

// Get gets the Pipeline information for the specified pipeline ID.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// pipelineIdentity - pipeline ID.
// startDateTime - the start date for when to get the pipeline and aggregate its data. The startDateTime and
// endDateTime can be no more than 30 days apart.
// endDateTime - the end date for when to get the pipeline and aggregate its data. The startDateTime and
// endDateTime can be no more than 30 days apart.
func (client PipelineClient) Get(ctx context.Context, accountName string, pipelineIdentity uuid.UUID, startDateTime *date.Time, endDateTime *date.Time) (result PipelineInformation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, accountName, pipelineIdentity, startDateTime, endDateTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.PipelineClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.PipelineClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.PipelineClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PipelineClient) GetPreparer(ctx context.Context, accountName string, pipelineIdentity uuid.UUID, startDateTime *date.Time, endDateTime *date.Time) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"pipelineIdentity": autorest.Encode("path", pipelineIdentity),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startDateTime != nil {
		queryParameters["startDateTime"] = autorest.Encode("query", *startDateTime)
	}
	if endDateTime != nil {
		queryParameters["endDateTime"] = autorest.Encode("query", *endDateTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/pipelines/{pipelineIdentity}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PipelineClient) GetResponder(resp *http.Response) (result PipelineInformation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all pipelines.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// startDateTime - the start date for when to get the list of pipelines. The startDateTime and endDateTime can
// be no more than 30 days apart.
// endDateTime - the end date for when to get the list of pipelines. The startDateTime and endDateTime can be
// no more than 30 days apart.
func (client PipelineClient) List(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (result PipelineInformationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineClient.List")
		defer func() {
			sc := -1
			if result.pilr.Response.Response != nil {
				sc = result.pilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, accountName, startDateTime, endDateTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.PipelineClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.PipelineClient", "List", resp, "Failure sending request")
		return
	}

	result.pilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.PipelineClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PipelineClient) ListPreparer(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startDateTime != nil {
		queryParameters["startDateTime"] = autorest.Encode("query", *startDateTime)
	}
	if endDateTime != nil {
		queryParameters["endDateTime"] = autorest.Encode("query", *endDateTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPath("/pipelines"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PipelineClient) ListResponder(resp *http.Response) (result PipelineInformationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PipelineClient) listNextResults(ctx context.Context, lastResults PipelineInformationListResult) (result PipelineInformationListResult, err error) {
	req, err := lastResults.pipelineInformationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "job.PipelineClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "job.PipelineClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.PipelineClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PipelineClient) ListComplete(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (result PipelineInformationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, accountName, startDateTime, endDateTime)
	return
}
