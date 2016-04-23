package service

import (
	"encoding/json"
	"github.com/coreos/etcd/storage/storagepb"
	"github.com/golang/glog"
	"github.com/infrmods/xbus/comm"
)

func (xbus *XBus) makeEndpoints(kvs []*storagepb.KeyValue) ([]comm.ServiceEndpoint, error) {
	endpoints := make([]comm.ServiceEndpoint, 0, len(kvs))
	for _, kv := range kvs {
		var endpoint comm.ServiceEndpoint
		if err := json.Unmarshal(kv.Value, &endpoint); err != nil {
			glog.Errorf("unmarshal endpoint fail(%#v): %v", string(kv.Value), err)
			return nil, comm.NewError(comm.EcodeDamagedEndpointValue, "")
		}
		endpoints = append(endpoints, endpoint)
	}
	return endpoints, nil
}
