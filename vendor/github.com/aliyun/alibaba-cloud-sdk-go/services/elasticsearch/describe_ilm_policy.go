package elasticsearch

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

// DescribeILMPolicy invokes the elasticsearch.DescribeILMPolicy API synchronously
func (client *Client) DescribeILMPolicy(request *DescribeILMPolicyRequest) (response *DescribeILMPolicyResponse, err error) {
	response = CreateDescribeILMPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeILMPolicyWithChan invokes the elasticsearch.DescribeILMPolicy API asynchronously
func (client *Client) DescribeILMPolicyWithChan(request *DescribeILMPolicyRequest) (<-chan *DescribeILMPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeILMPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeILMPolicy(request)
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

// DescribeILMPolicyWithCallback invokes the elasticsearch.DescribeILMPolicy API asynchronously
func (client *Client) DescribeILMPolicyWithCallback(request *DescribeILMPolicyRequest, callback func(response *DescribeILMPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeILMPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeILMPolicy(request)
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

// DescribeILMPolicyRequest is the request struct for api DescribeILMPolicy
type DescribeILMPolicyRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	PolicyName string `position:"Path" name:"PolicyName"`
}

// DescribeILMPolicyResponse is the response struct for api DescribeILMPolicy
type DescribeILMPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeILMPolicyRequest creates a request to invoke DescribeILMPolicy API
func CreateDescribeILMPolicyRequest() (request *DescribeILMPolicyRequest) {
	request = &DescribeILMPolicyRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DescribeILMPolicy", "/openapi/instances/[InstanceId]/ilm-policies/[PolicyName]", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeILMPolicyResponse creates a response to parse from DescribeILMPolicy response
func CreateDescribeILMPolicyResponse() (response *DescribeILMPolicyResponse) {
	response = &DescribeILMPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}