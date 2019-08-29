// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: dm.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	dm.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathDMSendMsg = "/live.approom.v1.DM/SendMsg"
var PathDMGetHistory = "/live.approom.v1.DM/GetHistory"

// ============
// DM Interface
// ============

type DMBMServer interface {
	// `method:"POST"`
	SendMsg(ctx context.Context, req *SendDMReq) (resp *SendMsgResp, err error)

	// `method:"POST"`
	GetHistory(ctx context.Context, req *HistoryReq) (resp *HistoryResp, err error)
}

var v1DMSvc DMBMServer

func dMSendMsg(c *bm.Context) {
	p := new(SendDMReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DMSvc.SendMsg(c, p)
	c.JSON(resp, err)
}

func dMGetHistory(c *bm.Context) {
	p := new(HistoryReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DMSvc.GetHistory(c, p)
	c.JSON(resp, err)
}

// RegisterV1DMService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1DMService(e *bm.Engine, svc DMBMServer, midMap map[string]bm.HandlerFunc) {
	v1DMSvc = svc
	e.POST("/xlive/app-room/v1/dM/SendMsg", dMSendMsg)
	e.POST("/xlive/app-room/v1/dM/GetHistory", dMGetHistory)
}

// RegisterDMBMServer Register the blademaster route
func RegisterDMBMServer(e *bm.Engine, server DMBMServer) {
	v1DMSvc = server
	e.POST("/live.approom.v1.DM/SendMsg", dMSendMsg)
	e.POST("/live.approom.v1.DM/GetHistory", dMGetHistory)
}
