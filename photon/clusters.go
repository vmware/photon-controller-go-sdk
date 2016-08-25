// Copyright (c) 2016 VMware, Inc. All Rights Reserved.
//
// This product is licensed to you under the Apache License, Version 2.0 (the "License").
// You may not use this product except in compliance with the License.
//
// This product may include a number of subcomponents with separate copyright notices and
// license terms. Your use of these subcomponents is subject to the terms and conditions
// of the subcomponent's license, as noted in the LICENSE file.

package photon

import (
	"bytes"
	"encoding/json"
)

// Contains functionality for clusters API.
type ClustersAPI struct {
	client *Client
}

var clusterUrl string = "/clusters/"
const EXTENDED_PROPERTY_DNS string = "dns"
const EXTENDED_PROPERTY_GATEWAY string = "gateway"
const EXTENDED_PROPERTY_NETMASK string = "netmask"
const EXTENDED_PROPERTY_MASTER_IP string = "master_ip"
const EXTENDED_PROPERTY_CONTAINER_NETWORK string = "container_network"
const EXTENDED_PROPERTY_ZOOKEEPER_IP1 string = "zookeeper_ip1"
const EXTENDED_PROPERTY_ZOOKEEPER_IP2 string = "zookeeper_ip2"
const EXTENDED_PROPERTY_ZOOKEEPER_IP3 string = "zookeeper_ip3"
const EXTENDED_PROPERTY_ETCD_IP1 string = "etcd_ip1"
const EXTENDED_PROPERTY_ETCD_IP2 string = "etcd_ip2"
const EXTENDED_PROPERTY_ETCD_IP3 string = "etcd_ip3"
const EXTENDED_PROPERTY_SSH_KEY string = "ssh_key"

// Deletes a cluster with specified ID.
func (api *ClustersAPI) Delete(id string) (task *Task, err error) {
	res, err := api.client.restClient.Delete(api.client.Endpoint+clusterUrl+id, api.client.options.TokenOptions.AccessToken)
	if err != nil {
		return
	}
	defer res.Body.Close()
	task, err = getTask(getError(res))
	return
}

// Gets a cluster with the specified ID.
func (api *ClustersAPI) Get(id string) (cluster *Cluster, err error) {
	res, err := api.client.restClient.Get(api.client.Endpoint+clusterUrl+id, api.client.options.TokenOptions.AccessToken)
	if err != nil {
		return
	}
	defer res.Body.Close()
	res, err = getError(res)
	if err != nil {
		return
	}
	var result Cluster
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, nil
}

// Gets vms for clusters with the specified ID
func (api *ClustersAPI) GetVMs(id string) (result *VMs, err error) {
	uri := api.client.Endpoint + clusterUrl + id + "/vms"
	res, err := api.client.restClient.GetList(api.client.Endpoint, uri, api.client.options.TokenOptions.AccessToken)
	if err != nil {
		return
	}

	result = &VMs{}
	err = json.Unmarshal(res, result)
	return
}

// Resize a cluster to specified count
func (api *ClustersAPI) Resize(id string, resize *ClusterResizeOperation) (task *Task, err error) {
	body, err := json.Marshal(resize)
	if err != nil {
		return
	}
	res, err := api.client.restClient.Post(
		api.client.Endpoint+clusterUrl+id+"/resize",
		"application/json",
		bytes.NewReader(body),
		api.client.options.TokenOptions.AccessToken)
	if err != nil {
		return
	}
	defer res.Body.Close()
	task, err = getTask(getError(res))
	return
}
