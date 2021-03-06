/*
Copyright 2020 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package snapshotmgr

import (
	v1api "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/apis/veleroplugin/v1"
	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/builder"
	plugin_clientset "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"testing"
	"time"
)

func TestDownload_Creation(t *testing.T) {
	path := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		t.Fatal("Got error " + err.Error())
	}
	pluginClient, err := plugin_clientset.NewForConfig(config)
	if err != nil {
		t.Fatal("Got error " + err.Error())
	}

	download := builder.ForDownload("velero", "download-1").RestoreTimestamp(time.Now()).SnapshotID("ssid-1").Phase(v1api.DownloadPhaseNew).Result()

	pluginClient.VeleropluginV1().Downloads("velero").Create(download)
}