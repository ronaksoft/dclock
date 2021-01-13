package service

import (
	fmt "fmt"
	rony "github.com/ronaksoft/rony"
	edge "github.com/ronaksoft/rony/edge"
	edgec "github.com/ronaksoft/rony/edgec"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_HookSetRequest int64 = 2791338713

type poolHookSetRequest struct {
	pool sync.Pool
}

func (p *poolHookSetRequest) Get() *HookSetRequest {
	x, ok := p.pool.Get().(*HookSetRequest)
	if !ok {
		return &HookSetRequest{}
	}
	return x
}

func (p *poolHookSetRequest) Put(x *HookSetRequest) {
	x.UniqueID = ""
	x.Timestamp = 0
	x.HookUrl = ""
	x.HookJsonData = ""
	p.pool.Put(x)
}

var PoolHookSetRequest = poolHookSetRequest{}

const C_HookSetResponse int64 = 2706970787

type poolHookSetResponse struct {
	pool sync.Pool
}

func (p *poolHookSetResponse) Get() *HookSetResponse {
	x, ok := p.pool.Get().(*HookSetResponse)
	if !ok {
		return &HookSetResponse{}
	}
	return x
}

func (p *poolHookSetResponse) Put(x *HookSetResponse) {
	x.Successful = false
	p.pool.Put(x)
}

var PoolHookSetResponse = poolHookSetResponse{}

const C_HookDeleteRequest int64 = 3968236869

type poolHookDeleteRequest struct {
	pool sync.Pool
}

func (p *poolHookDeleteRequest) Get() *HookDeleteRequest {
	x, ok := p.pool.Get().(*HookDeleteRequest)
	if !ok {
		return &HookDeleteRequest{}
	}
	return x
}

func (p *poolHookDeleteRequest) Put(x *HookDeleteRequest) {
	x.UniqueID = ""
	p.pool.Put(x)
}

var PoolHookDeleteRequest = poolHookDeleteRequest{}

const C_HookDeleteResponse int64 = 1487544771

type poolHookDeleteResponse struct {
	pool sync.Pool
}

func (p *poolHookDeleteResponse) Get() *HookDeleteResponse {
	x, ok := p.pool.Get().(*HookDeleteResponse)
	if !ok {
		return &HookDeleteResponse{}
	}
	return x
}

func (p *poolHookDeleteResponse) Put(x *HookDeleteResponse) {
	x.Successful = false
	p.pool.Put(x)
}

var PoolHookDeleteResponse = poolHookDeleteResponse{}

func init() {
	registry.RegisterConstructor(2791338713, "HookSetRequest")
	registry.RegisterConstructor(2706970787, "HookSetResponse")
	registry.RegisterConstructor(3968236869, "HookDeleteRequest")
	registry.RegisterConstructor(1487544771, "HookDeleteResponse")
	registry.RegisterConstructor(3312939871, "HookSet")
	registry.RegisterConstructor(3778745165, "HookDelete")
}

func (x *HookSetRequest) DeepCopy(z *HookSetRequest) {
	z.UniqueID = x.UniqueID
	z.Timestamp = x.Timestamp
	z.HookUrl = x.HookUrl
	z.HookJsonData = x.HookJsonData
}

func (x *HookSetResponse) DeepCopy(z *HookSetResponse) {
	z.Successful = x.Successful
}

func (x *HookDeleteRequest) DeepCopy(z *HookDeleteRequest) {
	z.UniqueID = x.UniqueID
}

func (x *HookDeleteResponse) DeepCopy(z *HookDeleteResponse) {
	z.Successful = x.Successful
}

func (x *HookSetRequest) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookSetRequest, x)
}

func (x *HookSetResponse) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookSetResponse, x)
}

func (x *HookDeleteRequest) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookDeleteRequest, x)
}

func (x *HookDeleteResponse) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookDeleteResponse, x)
}

func (x *HookSetRequest) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookSetResponse) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookDeleteRequest) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookDeleteResponse) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookSetRequest) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *HookSetResponse) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *HookDeleteRequest) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *HookDeleteResponse) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

const C_HookSet int64 = 3312939871
const C_HookDelete int64 = 3778745165

type IClock interface {
	HookSet(ctx *edge.RequestCtx, req *HookSetRequest, res *HookSetResponse)
	HookDelete(ctx *edge.RequestCtx, req *HookDeleteRequest, res *HookDeleteResponse)
}

type clockWrapper struct {
	h IClock
}

func RegisterClock(h IClock, e *edge.Server) {
	w := clockWrapper{
		h: h,
	}
	w.Register(e)
}

func (sw *clockWrapper) Register(e *edge.Server) {
	e.SetHandlers(C_HookSet, true, sw.HookSetWrapper)
	e.SetHandlers(C_HookDelete, true, sw.HookDeleteWrapper)
}

func (sw *clockWrapper) HookSetWrapper(ctx *edge.RequestCtx, in *rony.MessageEnvelope) {
	req := PoolHookSetRequest.Get()
	defer PoolHookSetRequest.Put(req)
	res := PoolHookSetResponse.Get()
	defer PoolHookSetResponse.Put(res)
	err := proto.UnmarshalOptions{Merge: true}.Unmarshal(in.Message, req)
	if err != nil {
		ctx.PushError(rony.ErrCodeInvalid, rony.ErrItemRequest)
		return
	}

	sw.h.HookSet(ctx, req, res)
	if !ctx.Stopped() {
		ctx.PushMessage(C_HookSetResponse, res)
	}
}

func (sw *clockWrapper) HookDeleteWrapper(ctx *edge.RequestCtx, in *rony.MessageEnvelope) {
	req := PoolHookDeleteRequest.Get()
	defer PoolHookDeleteRequest.Put(req)
	res := PoolHookDeleteResponse.Get()
	defer PoolHookDeleteResponse.Put(res)
	err := proto.UnmarshalOptions{Merge: true}.Unmarshal(in.Message, req)
	if err != nil {
		ctx.PushError(rony.ErrCodeInvalid, rony.ErrItemRequest)
		return
	}

	sw.h.HookDelete(ctx, req, res)
	if !ctx.Stopped() {
		ctx.PushMessage(C_HookDeleteResponse, res)
	}
}

func ExecuteRemoteHookSet(ctx *edge.RequestCtx, replicaSet uint64, req *HookSetRequest, res *HookSetResponse, kvs ...*rony.KeyValue) error {
	out := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(out)
	in := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(in)
	out.Fill(ctx.ReqID(), C_HookSet, req, kvs...)
	err := ctx.ExecuteRemote(replicaSet, true, out, in)
	if err != nil {
		return err
	}

	switch in.GetConstructor() {
	case C_HookSetResponse:
		_ = res.Unmarshal(in.GetMessage())
		return nil
	case rony.C_Error:
		x := &rony.Error{}
		_ = x.Unmarshal(in.GetMessage())
		return x
	default:
		return edge.ErrUnexpectedTunnelResponse
	}
}

func ExecuteRemoteHookDelete(ctx *edge.RequestCtx, replicaSet uint64, req *HookDeleteRequest, res *HookDeleteResponse, kvs ...*rony.KeyValue) error {
	out := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(out)
	in := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(in)
	out.Fill(ctx.ReqID(), C_HookDelete, req, kvs...)
	err := ctx.ExecuteRemote(replicaSet, true, out, in)
	if err != nil {
		return err
	}

	switch in.GetConstructor() {
	case C_HookDeleteResponse:
		_ = res.Unmarshal(in.GetMessage())
		return nil
	case rony.C_Error:
		x := &rony.Error{}
		_ = x.Unmarshal(in.GetMessage())
		return x
	default:
		return edge.ErrUnexpectedTunnelResponse
	}
}

type ClockClient struct {
	c edgec.Client
}

func NewClockClient(ec edgec.Client) *ClockClient {
	return &ClockClient{
		c: ec,
	}
}

func (c *ClockClient) HookSet(req *HookSetRequest, kvs ...*rony.KeyValue) (*HookSetResponse, error) {
	out := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(out)
	in := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(in)
	out.Fill(c.c.GetRequestID(), C_HookSet, req, kvs...)
	err := c.c.Send(out, in, true)
	if err != nil {
		return nil, err
	}
	switch in.GetConstructor() {
	case C_HookSetResponse:
		x := &HookSetResponse{}
		_ = proto.Unmarshal(in.Message, x)
		return x, nil
	case rony.C_Error:
		x := &rony.Error{}
		_ = proto.Unmarshal(in.Message, x)
		return nil, fmt.Errorf("%s:%s", x.GetCode(), x.GetItems())
	default:
		return nil, fmt.Errorf("unknown message: %d", in.GetConstructor())
	}
}

func (c *ClockClient) HookDelete(req *HookDeleteRequest, kvs ...*rony.KeyValue) (*HookDeleteResponse, error) {
	out := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(out)
	in := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(in)
	out.Fill(c.c.GetRequestID(), C_HookDelete, req, kvs...)
	err := c.c.Send(out, in, true)
	if err != nil {
		return nil, err
	}
	switch in.GetConstructor() {
	case C_HookDeleteResponse:
		x := &HookDeleteResponse{}
		_ = proto.Unmarshal(in.Message, x)
		return x, nil
	case rony.C_Error:
		x := &rony.Error{}
		_ = proto.Unmarshal(in.Message, x)
		return nil, fmt.Errorf("%s:%s", x.GetCode(), x.GetItems())
	default:
		return nil, fmt.Errorf("unknown message: %d", in.GetConstructor())
	}
}
