package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListClusterTemplates invokes the emr.ListClusterTemplates API synchronously
func (client *Client) ListClusterTemplates(request *ListClusterTemplatesRequest) (response *ListClusterTemplatesResponse, err error) {
	response = CreateListClusterTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterTemplatesWithChan invokes the emr.ListClusterTemplates API asynchronously
func (client *Client) ListClusterTemplatesWithChan(request *ListClusterTemplatesRequest) (<-chan *ListClusterTemplatesResponse, <-chan error) {
	responseChan := make(chan *ListClusterTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterTemplates(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListClusterTemplatesWithCallback invokes the emr.ListClusterTemplates API asynchronously
func (client *Client) ListClusterTemplatesWithCallback(request *ListClusterTemplatesRequest, callback func(response *ListClusterTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterTemplatesResponse
		var err error
		defer close(result)
		response, err = client.ListClusterTemplates(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListClusterTemplatesRequest is the request struct for api ListClusterTemplates
type ListClusterTemplatesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ProductType     string           `position:"Query" name:"ProductType"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	BizId           string           `position:"Query" name:"BizId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// ListClusterTemplatesResponse is the response struct for api ListClusterTemplates
type ListClusterTemplatesResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	TotalCount       int              `json:"TotalCount" xml:"TotalCount"`
	PageNumber       int              `json:"PageNumber" xml:"PageNumber"`
	PageSize         int              `json:"PageSize" xml:"PageSize"`
	TemplateInfoList TemplateInfoList `json:"TemplateInfoList" xml:"TemplateInfoList"`
}

// CreateListClusterTemplatesRequest creates a request to invoke ListClusterTemplates API
func CreateListClusterTemplatesRequest() (request *ListClusterTemplatesRequest) {
	request = &ListClusterTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterTemplates", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterTemplatesResponse creates a response to parse from ListClusterTemplates response
func CreateListClusterTemplatesResponse() (response *ListClusterTemplatesResponse) {
	response = &ListClusterTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}