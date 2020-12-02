// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"net/http"
)

// ListVaultsRequest wrapper for the ListVaults operation
type ListVaultsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `TIMECREATED` is descending. The default order for `DISPLAYNAME`
	// is ascending.
	SortBy ListVaultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListVaultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVaultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVaultsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVaultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListVaultsResponse wrapper for the ListVaults operation
type ListVaultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []VaultSummary instances
	Items []VaultSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVaultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVaultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVaultsSortByEnum Enum with underlying type: string
type ListVaultsSortByEnum string

// Set of constants representing the allowable values for ListVaultsSortByEnum
const (
	ListVaultsSortByTimecreated ListVaultsSortByEnum = "TIMECREATED"
	ListVaultsSortByDisplayname ListVaultsSortByEnum = "DISPLAYNAME"
)

var mappingListVaultsSortBy = map[string]ListVaultsSortByEnum{
	"TIMECREATED": ListVaultsSortByTimecreated,
	"DISPLAYNAME": ListVaultsSortByDisplayname,
}

// GetListVaultsSortByEnumValues Enumerates the set of values for ListVaultsSortByEnum
func GetListVaultsSortByEnumValues() []ListVaultsSortByEnum {
	values := make([]ListVaultsSortByEnum, 0)
	for _, v := range mappingListVaultsSortBy {
		values = append(values, v)
	}
	return values
}

// ListVaultsSortOrderEnum Enum with underlying type: string
type ListVaultsSortOrderEnum string

// Set of constants representing the allowable values for ListVaultsSortOrderEnum
const (
	ListVaultsSortOrderAsc  ListVaultsSortOrderEnum = "ASC"
	ListVaultsSortOrderDesc ListVaultsSortOrderEnum = "DESC"
)

var mappingListVaultsSortOrder = map[string]ListVaultsSortOrderEnum{
	"ASC":  ListVaultsSortOrderAsc,
	"DESC": ListVaultsSortOrderDesc,
}

// GetListVaultsSortOrderEnumValues Enumerates the set of values for ListVaultsSortOrderEnum
func GetListVaultsSortOrderEnumValues() []ListVaultsSortOrderEnum {
	values := make([]ListVaultsSortOrderEnum, 0)
	for _, v := range mappingListVaultsSortOrder {
		values = append(values, v)
	}
	return values
}