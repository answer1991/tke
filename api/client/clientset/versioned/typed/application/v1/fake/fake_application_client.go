/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1 "tkestack.io/tke/api/client/clientset/versioned/typed/application/v1"
)

type FakeApplicationV1 struct {
	*testing.Fake
}

func (c *FakeApplicationV1) Apps(namespace string) v1.AppInterface {
	return &FakeApps{c, namespace}
}

func (c *FakeApplicationV1) AppHistories(namespace string) v1.AppHistoryInterface {
	return &FakeAppHistories{c, namespace}
}

func (c *FakeApplicationV1) AppResources(namespace string) v1.AppResourceInterface {
	return &FakeAppResources{c, namespace}
}

func (c *FakeApplicationV1) ConfigMaps() v1.ConfigMapInterface {
	return &FakeConfigMaps{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeApplicationV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
