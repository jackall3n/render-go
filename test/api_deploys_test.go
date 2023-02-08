/*
Render Public API

Testing DeploysApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package render

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_render_DeploysApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DeploysApiService CreateDeploy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.DeploysApi.CreateDeploy(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DeploysApiService GetDeploy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string
        var deployId string

        resp, httpRes, err := apiClient.DeploysApi.GetDeploy(context.Background(), serviceId, deployId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DeploysApiService GetDeploys", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.DeploysApi.GetDeploys(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}