package service

import (
	fmt "fmt"
	rony "github.com/ronaksoft/rony"
	config "github.com/ronaksoft/rony/config"
	edge "github.com/ronaksoft/rony/edge"
	edgec "github.com/ronaksoft/rony/edgec"
	registry "github.com/ronaksoft/rony/registry"
	cobra "github.com/spf13/cobra"
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
	x.UniqueID = x.UniqueID[:0]
	x.Timestamp = 0
	x.HookUrl = x.HookUrl[:0]
	x.HookJsonData = x.HookJsonData[:0]
	p.pool.Put(x)
}

var PoolHookSetRequest = poolHookSetRequest{}

func (x *HookSetRequest) DeepCopy(z *HookSetRequest) {
	z.UniqueID = append(z.UniqueID[:0], x.UniqueID...)
	z.Timestamp = x.Timestamp
	z.HookUrl = append(z.HookUrl[:0], x.HookUrl...)
	z.HookJsonData = append(z.HookJsonData[:0], x.HookJsonData...)
}

func (x *HookSetRequest) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookSetRequest) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *HookSetRequest) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookSetRequest, x)
}

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

func (x *HookSetResponse) DeepCopy(z *HookSetResponse) {
	z.Successful = x.Successful
}

func (x *HookSetResponse) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookSetResponse) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *HookSetResponse) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookSetResponse, x)
}

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
	x.UniqueID = x.UniqueID[:0]
	p.pool.Put(x)
}

var PoolHookDeleteRequest = poolHookDeleteRequest{}

func (x *HookDeleteRequest) DeepCopy(z *HookDeleteRequest) {
	z.UniqueID = append(z.UniqueID[:0], x.UniqueID...)
}

func (x *HookDeleteRequest) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookDeleteRequest) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *HookDeleteRequest) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookDeleteRequest, x)
}

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

func (x *HookDeleteResponse) DeepCopy(z *HookDeleteResponse) {
	z.Successful = x.Successful
}

func (x *HookDeleteResponse) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookDeleteResponse) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *HookDeleteResponse) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookDeleteResponse, x)
}

const C_HookSet int64 = 3312939871
const C_HookDelete int64 = 3778745165

func init() {
	registry.RegisterConstructor(2791338713, "HookSetRequest")
	registry.RegisterConstructor(2706970787, "HookSetResponse")
	registry.RegisterConstructor(3968236869, "HookDeleteRequest")
	registry.RegisterConstructor(1487544771, "HookDeleteResponse")
	registry.RegisterConstructor(3312939871, "HookSet")
	registry.RegisterConstructor(3778745165, "HookDelete")
}

type IClock interface {
	HookSet(ctx *edge.RequestCtx, req *HookSetRequest, res *HookSetResponse)
	HookDelete(ctx *edge.RequestCtx, req *HookDeleteRequest, res *HookDeleteResponse)
}

type clockWrapper struct {
	h IClock
}

func (sw *clockWrapper) hookSetWrapper(ctx *edge.RequestCtx, in *rony.MessageEnvelope) {
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

func (sw *clockWrapper) hookDeleteWrapper(ctx *edge.RequestCtx, in *rony.MessageEnvelope) {
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

func (sw *clockWrapper) Register(e *edge.Server) {
	e.SetHandlers(C_HookSet, true, sw.hookSetWrapper)
	e.SetHandlers(C_HookDelete, true, sw.hookDeleteWrapper)
}

func RegisterClock(h IClock, e *edge.Server) {
	w := clockWrapper{
		h: h,
	}
	w.Register(e)
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

func prepareClockCommand(cmd *cobra.Command) (*ClockClient, error) {
	// Bind the current flags to registered flags in config package
	err := config.BindCmdFlags(cmd)
	if err != nil {
		return nil, err
	}

	httpC := edgec.NewHttp(edgec.HttpConfig{
		Name:         "Rony Client",
		SeedHostPort: fmt.Sprintf("%s:%d", config.GetString("host"), config.GetInt("port")),
	})

	err = httpC.Start()
	if err != nil {
		return nil, err
	}
	return NewClockClient(httpC), nil
}

var genHookSetCmd = func(h IClockCli) *cobra.Command {
	cmd := &cobra.Command{
		Use: "hook-set",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli, err := prepareClockCommand(cmd)
			if err != nil {
				return err
			}
			return h.HookSet(cli, cmd, args)
		},
	}
	config.SetFlags(cmd,
		config.StringFlag("uniqueID", "", ""),
		config.Int64Flag("timestamp", 0, ""),
		config.StringFlag("hookUrl", "", ""),
		config.StringFlag("hookJsonData", "", ""),
	)
	return cmd
}

var genHookDeleteCmd = func(h IClockCli) *cobra.Command {
	cmd := &cobra.Command{
		Use: "hook-delete",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli, err := prepareClockCommand(cmd)
			if err != nil {
				return err
			}
			return h.HookDelete(cli, cmd, args)
		},
	}
	config.SetFlags(cmd,
		config.StringFlag("uniqueID", "", ""),
	)
	return cmd
}

type IClockCli interface {
	HookSet(cli *ClockClient, cmd *cobra.Command, args []string) error
	HookDelete(cli *ClockClient, cmd *cobra.Command, args []string) error
}

func RegisterClockCli(h IClockCli, rootCmd *cobra.Command) {
	config.SetPersistentFlags(rootCmd,
		config.StringFlag("host", "127.0.0.1", "the seed host's address"),
		config.StringFlag("port", "80", "the seed host's port"),
	)
	rootCmd.AddCommand(
		genHookSetCmd(h), genHookDeleteCmd(h),
	)
}
