package synchronizationjob

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateSynchronizationJobsCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ValidateSynchronizationJobsCredentialsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultValidateSynchronizationJobsCredentialsOperationOptions() ValidateSynchronizationJobsCredentialsOperationOptions {
	return ValidateSynchronizationJobsCredentialsOperationOptions{}
}

func (o ValidateSynchronizationJobsCredentialsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ValidateSynchronizationJobsCredentialsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ValidateSynchronizationJobsCredentialsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ValidateSynchronizationJobsCredentials - Invoke action validateCredentials
func (c SynchronizationJobClient) ValidateSynchronizationJobsCredentials(ctx context.Context, id stable.ServicePrincipalId, input ValidateSynchronizationJobsCredentialsRequest, options ValidateSynchronizationJobsCredentialsOperationOptions) (result ValidateSynchronizationJobsCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/synchronization/jobs/validateCredentials", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
