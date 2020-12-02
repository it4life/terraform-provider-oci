// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// AddPublicIpPoolCapacityDetails The data to add capacity to a public ip pool
type AddPublicIpPoolCapacityDetails struct {

	// The OCID of the Byoip Range Id object to whch the cidr block belongs.
	ByoipRangeId *string `mandatory:"true" json:"byoipRangeId"`

	// The CIDR IP address range to be added to the Public Ip Pool
	// Example: `10.0.1.0/24`
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`
}

func (m AddPublicIpPoolCapacityDetails) String() string {
	return common.PointerString(m)
}