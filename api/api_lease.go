package api

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/infrmods/xbus/utils"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

func parseLeaseId(s string) (clientv3.LeaseID, error) {
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return clientv3.LeaseID(n), nil
	} else {
		return 0, echo.NewHTTPError(http.StatusNotFound)
	}
}

func (server *APIServer) KeepAliveLease(c echo.Context) error {
	leaseId, err := parseLeaseId(c.P(0))
	if err != nil {
		return err
	}
	if _, err := server.etcdClient.KeepAliveOnce(context.Background(), leaseId); err == nil {
		return JsonOk(c)
	} else {
		return JsonError(c, utils.CleanErr(err, "keepalive fail", "keepalive(%d) fail: %v", leaseId, err))
	}
}

func (server *APIServer) RevokeLease(c echo.Context) error {
	leaseId, err := parseLeaseId(c.P(0))
	if err != nil {
		return err
	}
	if _, err := server.etcdClient.Revoke(context.Background(), leaseId); err == nil {
		return JsonOk(c)
	} else {
		return JsonError(c, utils.CleanErr(err, "revoke fail", "revoke(%d) fail: %v", leaseId, err))
	}
}
