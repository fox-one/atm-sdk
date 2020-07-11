package atm

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	fmt "fmt"
	io "io"
	ioutil "io/ioutil"
	http "net/http"
	url "net/url"
	strconv "strconv"
	strings "strings"

	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	twirp "github.com/twitchtv/twirp"
	ctxsetters "github.com/twitchtv/twirp/ctxsetters"
)

// Imports only used by utility functions:

// =====================
// UserService Interface
// =====================

// UserService handle user requests
type UserService interface {
	// 获取个人信息 GET /api/me
	Me(context.Context, *UserServiceReq_Me) (*User, error)
}

// ===========================
// UserService Protobuf Client
// ===========================

type userServiceProtobufClient struct {
	client HTTPClient
	urls   [1]string
	opts   twirp.ClientOptions
}

// NewUserServiceProtobufClient creates a Protobuf client that implements the UserService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewUserServiceProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + UserServicePathPrefix
	urls := [1]string{
		prefix + "Me",
	}

	return &userServiceProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *userServiceProtobufClient) Me(ctx context.Context, in *UserServiceReq_Me) (*User, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	out := new(User)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// =======================
// UserService JSON Client
// =======================

type userServiceJSONClient struct {
	client HTTPClient
	urls   [1]string
	opts   twirp.ClientOptions
}

// NewUserServiceJSONClient creates a JSON client that implements the UserService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewUserServiceJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + UserServicePathPrefix
	urls := [1]string{
		prefix + "Me",
	}

	return &userServiceJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *userServiceJSONClient) Me(ctx context.Context, in *UserServiceReq_Me) (*User, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	out := new(User)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==========================
// UserService Server Handler
// ==========================

type userServiceServer struct {
	UserService
	hooks *twirp.ServerHooks
}

func NewUserServiceServer(svc UserService, hooks *twirp.ServerHooks) TwirpServer {
	return &userServiceServer{
		UserService: svc,
		hooks:       hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *userServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// UserServicePathPrefix is used for all URL paths on a twirp UserService server.
// Requests are always: POST UserServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const UserServicePathPrefix = "/twirp/fox.atm.service.UserService/"

func (s *userServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/fox.atm.service.UserService/Me":
		s.serveMe(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *userServiceServer) serveMe(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveMeJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveMeProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveMeJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(UserServiceReq_Me)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *User
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Me(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *User and nil error while calling Me. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveMeProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(UserServiceReq_Me)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *User
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Me(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *User and nil error while calling Me. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *userServiceServer) ProtocGenTwirpVersion() string {
	return "v5.10.1"
}

func (s *userServiceServer) PathPrefix() string {
	return UserServicePathPrefix
}

// =========================
// MerchantService Interface
// =========================

// MerchantService handle merchant request
type MerchantService interface {
	// 预创建订单 POST /api/m/books
	CreateBook(context.Context, *MerchantServiceReq_CreateBook) (*Book, error)

	// 获取订单详情 GET /api/m/order/{trace_id}
	ReadOrder(context.Context, *MerchantServiceReq_ReadOrder) (*Order, error)

	// 查询订单列表 GET /api/m/orders
	ListOrders(context.Context, *MerchantServiceReq_ListOrders) (*MerchantServiceResp_ListOrders, error)

	// 查询订单报表 GET /api/m/order-reports
	ListOrderReports(context.Context, *MerchantServiceReq_ListOrderReports) (*MerchantServiceResp_ListOrderReports, error)

	// 撤单 DELETE /api/m/order/{trace_id}
	CancelOrder(context.Context, *MerchantServiceReq_CancelOrder) (*MerchantServiceResp_CancelOrder, error)
}

// ===============================
// MerchantService Protobuf Client
// ===============================

type merchantServiceProtobufClient struct {
	client HTTPClient
	urls   [5]string
	opts   twirp.ClientOptions
}

// NewMerchantServiceProtobufClient creates a Protobuf client that implements the MerchantService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewMerchantServiceProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) MerchantService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + MerchantServicePathPrefix
	urls := [5]string{
		prefix + "CreateBook",
		prefix + "ReadOrder",
		prefix + "ListOrders",
		prefix + "ListOrderReports",
		prefix + "CancelOrder",
	}

	return &merchantServiceProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *merchantServiceProtobufClient) CreateBook(ctx context.Context, in *MerchantServiceReq_CreateBook) (*Book, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateBook")
	out := new(Book)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *merchantServiceProtobufClient) ReadOrder(ctx context.Context, in *MerchantServiceReq_ReadOrder) (*Order, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "ReadOrder")
	out := new(Order)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *merchantServiceProtobufClient) ListOrders(ctx context.Context, in *MerchantServiceReq_ListOrders) (*MerchantServiceResp_ListOrders, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "ListOrders")
	out := new(MerchantServiceResp_ListOrders)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *merchantServiceProtobufClient) ListOrderReports(ctx context.Context, in *MerchantServiceReq_ListOrderReports) (*MerchantServiceResp_ListOrderReports, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "ListOrderReports")
	out := new(MerchantServiceResp_ListOrderReports)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[3], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *merchantServiceProtobufClient) CancelOrder(ctx context.Context, in *MerchantServiceReq_CancelOrder) (*MerchantServiceResp_CancelOrder, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "CancelOrder")
	out := new(MerchantServiceResp_CancelOrder)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[4], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===========================
// MerchantService JSON Client
// ===========================

type merchantServiceJSONClient struct {
	client HTTPClient
	urls   [5]string
	opts   twirp.ClientOptions
}

// NewMerchantServiceJSONClient creates a JSON client that implements the MerchantService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewMerchantServiceJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) MerchantService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + MerchantServicePathPrefix
	urls := [5]string{
		prefix + "CreateBook",
		prefix + "ReadOrder",
		prefix + "ListOrders",
		prefix + "ListOrderReports",
		prefix + "CancelOrder",
	}

	return &merchantServiceJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *merchantServiceJSONClient) CreateBook(ctx context.Context, in *MerchantServiceReq_CreateBook) (*Book, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateBook")
	out := new(Book)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *merchantServiceJSONClient) ReadOrder(ctx context.Context, in *MerchantServiceReq_ReadOrder) (*Order, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "ReadOrder")
	out := new(Order)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *merchantServiceJSONClient) ListOrders(ctx context.Context, in *MerchantServiceReq_ListOrders) (*MerchantServiceResp_ListOrders, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "ListOrders")
	out := new(MerchantServiceResp_ListOrders)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *merchantServiceJSONClient) ListOrderReports(ctx context.Context, in *MerchantServiceReq_ListOrderReports) (*MerchantServiceResp_ListOrderReports, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "ListOrderReports")
	out := new(MerchantServiceResp_ListOrderReports)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[3], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *merchantServiceJSONClient) CancelOrder(ctx context.Context, in *MerchantServiceReq_CancelOrder) (*MerchantServiceResp_CancelOrder, error) {
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithMethodName(ctx, "CancelOrder")
	out := new(MerchantServiceResp_CancelOrder)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[4], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==============================
// MerchantService Server Handler
// ==============================

type merchantServiceServer struct {
	MerchantService
	hooks *twirp.ServerHooks
}

func NewMerchantServiceServer(svc MerchantService, hooks *twirp.ServerHooks) TwirpServer {
	return &merchantServiceServer{
		MerchantService: svc,
		hooks:           hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *merchantServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// MerchantServicePathPrefix is used for all URL paths on a twirp MerchantService server.
// Requests are always: POST MerchantServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const MerchantServicePathPrefix = "/twirp/fox.atm.service.MerchantService/"

func (s *merchantServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "fox.atm.service")
	ctx = ctxsetters.WithServiceName(ctx, "MerchantService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/fox.atm.service.MerchantService/CreateBook":
		s.serveCreateBook(ctx, resp, req)
		return
	case "/twirp/fox.atm.service.MerchantService/ReadOrder":
		s.serveReadOrder(ctx, resp, req)
		return
	case "/twirp/fox.atm.service.MerchantService/ListOrders":
		s.serveListOrders(ctx, resp, req)
		return
	case "/twirp/fox.atm.service.MerchantService/ListOrderReports":
		s.serveListOrderReports(ctx, resp, req)
		return
	case "/twirp/fox.atm.service.MerchantService/CancelOrder":
		s.serveCancelOrder(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *merchantServiceServer) serveCreateBook(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateBookJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateBookProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *merchantServiceServer) serveCreateBookJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateBook")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(MerchantServiceReq_CreateBook)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Book
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.CreateBook(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Book and nil error while calling CreateBook. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveCreateBookProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateBook")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(MerchantServiceReq_CreateBook)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Book
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.CreateBook(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Book and nil error while calling CreateBook. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveReadOrder(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveReadOrderJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveReadOrderProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *merchantServiceServer) serveReadOrderJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ReadOrder")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(MerchantServiceReq_ReadOrder)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Order
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.ReadOrder(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Order and nil error while calling ReadOrder. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveReadOrderProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ReadOrder")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(MerchantServiceReq_ReadOrder)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Order
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.ReadOrder(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Order and nil error while calling ReadOrder. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveListOrders(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListOrdersJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListOrdersProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *merchantServiceServer) serveListOrdersJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListOrders")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(MerchantServiceReq_ListOrders)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *MerchantServiceResp_ListOrders
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.ListOrders(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *MerchantServiceResp_ListOrders and nil error while calling ListOrders. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveListOrdersProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListOrders")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(MerchantServiceReq_ListOrders)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *MerchantServiceResp_ListOrders
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.ListOrders(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *MerchantServiceResp_ListOrders and nil error while calling ListOrders. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveListOrderReports(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListOrderReportsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListOrderReportsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *merchantServiceServer) serveListOrderReportsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListOrderReports")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(MerchantServiceReq_ListOrderReports)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *MerchantServiceResp_ListOrderReports
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.ListOrderReports(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *MerchantServiceResp_ListOrderReports and nil error while calling ListOrderReports. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveListOrderReportsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListOrderReports")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(MerchantServiceReq_ListOrderReports)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *MerchantServiceResp_ListOrderReports
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.ListOrderReports(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *MerchantServiceResp_ListOrderReports and nil error while calling ListOrderReports. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveCancelOrder(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCancelOrderJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCancelOrderProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *merchantServiceServer) serveCancelOrderJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CancelOrder")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(MerchantServiceReq_CancelOrder)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *MerchantServiceResp_CancelOrder
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.CancelOrder(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *MerchantServiceResp_CancelOrder and nil error while calling CancelOrder. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) serveCancelOrderProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CancelOrder")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(MerchantServiceReq_CancelOrder)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *MerchantServiceResp_CancelOrder
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.MerchantService.CancelOrder(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *MerchantServiceResp_CancelOrder and nil error while calling CancelOrder. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *merchantServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 1
}

func (s *merchantServiceServer) ProtocGenTwirpVersion() string {
	return "v5.10.1"
}

func (s *merchantServiceServer) PathPrefix() string {
	return MerchantServicePathPrefix
}

// =====
// Utils
// =====

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type TwirpServer interface {
	http.Handler
	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// github.com/golang/protobuf/protoc-gen-go/descriptor.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)
	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string
	// PathPrefix returns the HTTP URL path prefix for all methods handled by this
	// service. This can be used with an HTTP mux to route twirp requests
	// alongside non-twirp requests on one HTTP listener.
	PathPrefix() string
}

// WriteError writes an HTTP response with a valid Twirp error format (code, msg, meta).
// Useful outside of the Twirp server (e.g. http middleware), but does not trigger hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func WriteError(resp http.ResponseWriter, err error) {
	writeError(context.Background(), resp, err, nil)
}

// writeError writes Twirp errors in the response and triggers hooks.
func writeError(ctx context.Context, resp http.ResponseWriter, err error, hooks *twirp.ServerHooks) {
	// Non-twirp errors are wrapped as Internal (default)
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
	ctx = ctxsetters.WithStatusCode(ctx, statusCode)
	ctx = callError(ctx, hooks, twerr)

	respBody := marshalErrorToJSON(twerr)

	resp.Header().Set("Content-Type", "application/json") // Error responses are always JSON
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBody)))
	resp.WriteHeader(statusCode) // set HTTP status code and send response

	_, writeErr := resp.Write(respBody)
	if writeErr != nil {
		// We have three options here. We could log the error, call the Error
		// hook, or just silently ignore the error.
		//
		// Logging is unacceptable because we don't have a user-controlled
		// logger; writing out to stderr without permission is too rude.
		//
		// Calling the Error hook would confuse users: it would mean the Error
		// hook got called twice for one request, which is likely to lead to
		// duplicated log messages and metrics, no matter how well we document
		// the behavior.
		//
		// Silently ignoring the error is our least-bad option. It's highly
		// likely that the connection is broken and the original 'err' says
		// so anyway.
		_ = writeErr
	}

	callResponseSent(ctx, hooks)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// getCustomHTTPReqHeaders retrieves a copy of any headers that are set in
// a context through the twirp.WithHTTPRequestHeaders function.
// If there are no headers set, or if they have the wrong type, nil is returned.
func getCustomHTTPReqHeaders(ctx context.Context) http.Header {
	header, ok := twirp.HTTPRequestHeaders(ctx)
	if !ok || header == nil {
		return nil
	}
	copied := make(http.Header)
	for k, vv := range header {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}
	return copied
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
		req.Header = customHeader
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Twirp-Version", "v5.10.1")
	return req, nil
}

// JSON serialization for errors
type twerrJSON struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

// marshalErrorToJSON returns JSON from a twirp.Error, that can be used as HTTP error response body.
// If serialization fails, it will use a descriptive Internal error instead.
func marshalErrorToJSON(twerr twirp.Error) []byte {
	// make sure that msg is not too large
	msg := twerr.Msg()
	if len(msg) > 1e6 {
		msg = msg[:1e6]
	}

	tj := twerrJSON{
		Code: string(twerr.Code()),
		Msg:  msg,
		Meta: twerr.MetaMap(),
	}

	buf, err := json.Marshal(&tj)
	if err != nil {
		buf = []byte("{\"type\": \"" + twirp.Internal + "\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}

	return buf
}

// errorFromResponse builds a twirp.Error from a non-200 HTTP response.
// If the response has a valid serialized Twirp error, then it's returned.
// If not, the response status code is used to generate a similar twirp
// error. See twirpErrorFromIntermediary for more info on intermediary errors.
func errorFromResponse(resp *http.Response) twirp.Error {
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	if isHTTPRedirect(statusCode) {
		// Unexpected redirect: it must be an error from an intermediary.
		// Twirp clients don't follow redirects automatically, Twirp only handles
		// POST requests, redirects should only happen on GET and HEAD requests.
		location := resp.Header.Get("Location")
		msg := fmt.Sprintf("unexpected HTTP status code %d %q received, Location=%q", statusCode, statusText, location)
		return twirpErrorFromIntermediary(statusCode, msg, location)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wrapInternal(err, "failed to read server error response body")
	}

	var tj twerrJSON
	dec := json.NewDecoder(bytes.NewReader(respBodyBytes))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&tj); err != nil || tj.Code == "" {
		// Invalid JSON response; it must be an error from an intermediary.
		msg := fmt.Sprintf("Error from intermediary with HTTP status code %d %q", statusCode, statusText)
		return twirpErrorFromIntermediary(statusCode, msg, string(respBodyBytes))
	}

	errorCode := twirp.ErrorCode(tj.Code)
	if !twirp.IsValidErrorCode(errorCode) {
		msg := "invalid type returned from server error response: " + tj.Code
		return twirp.InternalError(msg)
	}

	twerr := twirp.NewError(errorCode, tj.Msg)
	for k, v := range tj.Meta {
		twerr = twerr.WithMeta(k, v)
	}
	return twerr
}

// twirpErrorFromIntermediary maps HTTP errors from non-twirp sources to twirp errors.
// The mapping is similar to gRPC: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md.
// Returned twirp Errors have some additional metadata for inspection.
func twirpErrorFromIntermediary(status int, msg string, bodyOrLocation string) twirp.Error {
	var code twirp.ErrorCode
	if isHTTPRedirect(status) { // 3xx
		code = twirp.Internal
	} else {
		switch status {
		case 400: // Bad Request
			code = twirp.Internal
		case 401: // Unauthorized
			code = twirp.Unauthenticated
		case 403: // Forbidden
			code = twirp.PermissionDenied
		case 404: // Not Found
			code = twirp.BadRoute
		case 429, 502, 503, 504: // Too Many Requests, Bad Gateway, Service Unavailable, Gateway Timeout
			code = twirp.Unavailable
		default: // All other codes
			code = twirp.Unknown
		}
	}

	twerr := twirp.NewError(code, msg)
	twerr = twerr.WithMeta("http_error_from_intermediary", "true") // to easily know if this error was from intermediary
	twerr = twerr.WithMeta("status_code", strconv.Itoa(status))
	if isHTTPRedirect(status) {
		twerr = twerr.WithMeta("location", bodyOrLocation)
	} else {
		twerr = twerr.WithMeta("body", bodyOrLocation)
	}
	return twerr
}

func isHTTPRedirect(status int) bool {
	return status >= 300 && status <= 399
}

// wrapInternal wraps an error with a prefix as an Internal error.
// The original error cause is accessible by github.com/pkg/errors.Cause.
func wrapInternal(err error, prefix string) twirp.Error {
	return twirp.InternalErrorWith(&wrappedError{prefix: prefix, cause: err})
}

type wrappedError struct {
	prefix string
	cause  error
}

func (e *wrappedError) Cause() error  { return e.cause }
func (e *wrappedError) Error() string { return e.prefix + ": " + e.cause.Error() }

// ensurePanicResponses makes sure that rpc methods causing a panic still result in a Twirp Internal
// error response (status 500), and error hooks are properly called with the panic wrapped as an error.
// The panic is re-raised so it can be handled normally with middleware.
func ensurePanicResponses(ctx context.Context, resp http.ResponseWriter, hooks *twirp.ServerHooks) {
	if r := recover(); r != nil {
		// Wrap the panic as an error so it can be passed to error hooks.
		// The original error is accessible from error hooks, but not visible in the response.
		err := errFromPanic(r)
		twerr := &internalWithCause{msg: "Internal service panic", cause: err}
		// Actually write the error
		writeError(ctx, resp, twerr, hooks)
		// If possible, flush the error to the wire.
		f, ok := resp.(http.Flusher)
		if ok {
			f.Flush()
		}

		panic(r)
	}
}

// errFromPanic returns the typed error if the recovered panic is an error, otherwise formats as error.
func errFromPanic(p interface{}) error {
	if err, ok := p.(error); ok {
		return err
	}
	return fmt.Errorf("panic: %v", p)
}

// internalWithCause is a Twirp Internal error wrapping an original error cause, accessible
// by github.com/pkg/errors.Cause, but the original error message is not exposed on Msg().
type internalWithCause struct {
	msg   string
	cause error
}

func (e *internalWithCause) Cause() error                                { return e.cause }
func (e *internalWithCause) Error() string                               { return e.msg + ": " + e.cause.Error() }
func (e *internalWithCause) Code() twirp.ErrorCode                       { return twirp.Internal }
func (e *internalWithCause) Msg() string                                 { return e.msg }
func (e *internalWithCause) Meta(key string) string                      { return "" }
func (e *internalWithCause) MetaMap() map[string]string                  { return nil }
func (e *internalWithCause) WithMeta(key string, val string) twirp.Error { return e }

// malformedRequestError is used when the twirp server cannot unmarshal a request
func malformedRequestError(msg string) twirp.Error {
	return twirp.NewError(twirp.Malformed, msg)
}

// badRouteError is used when the twirp server cannot route a request
func badRouteError(msg string, method, url string) twirp.Error {
	err := twirp.NewError(twirp.BadRoute, msg)
	err = err.WithMeta("twirp_invalid_route", method+" "+url)
	return err
}

// withoutRedirects makes sure that the POST request can not be redirected.
// The standard library will, by default, redirect requests (including POSTs) if it gets a 302 or
// 303 response, and also 301s in go1.8. It redirects by making a second request, changing the
// method to GET and removing the body. This produces very confusing error messages, so instead we
// set a redirect policy that always errors. This stops Go from executing the redirect.
//
// We have to be a little careful in case the user-provided http.Client has its own CheckRedirect
// policy - if so, we'll run through that policy first.
//
// Because this requires modifying the http.Client, we make a new copy of the client and return it.
func withoutRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			// Run the input's redirect if it exists, in case it has side effects, but ignore any error it
			// returns, since we want to use ErrUseLastResponse.
			err := in.CheckRedirect(req, via)
			_ = err // Silly, but this makes sure generated code passes errcheck -blank, which some people use.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

// doProtobufRequest makes a Protobuf request to the remote Twirp service.
func doProtobufRequest(ctx context.Context, client HTTPClient, hooks *twirp.ClientHooks, url string, in, out proto.Message) (err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return wrapInternal(err, "failed to marshal proto request")
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return wrapInternal(err, "could not build request")
	}
	ctx, err = callClientRequestPrepared(ctx, hooks, req)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wrapInternal(err, "failed to read response body")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return wrapInternal(err, "failed to unmarshal proto response")
	}
	return nil
}

// doJSONRequest makes a JSON request to the remote Twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, hooks *twirp.ClientHooks, url string, in, out proto.Message) (err error) {
	reqBody := bytes.NewBuffer(nil)
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(reqBody, in); err != nil {
		return wrapInternal(err, "failed to marshal json request")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/json")
	if err != nil {
		return wrapInternal(err, "could not build request")
	}
	ctx, err = callClientRequestPrepared(ctx, hooks, req)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
		return wrapInternal(err, "failed to unmarshal json response")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}
	return nil
}

// Call twirp.ServerHooks.RequestReceived if the hook is available
func callRequestReceived(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestReceived == nil {
		return ctx, nil
	}
	return h.RequestReceived(ctx)
}

// Call twirp.ServerHooks.RequestRouted if the hook is available
func callRequestRouted(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestRouted == nil {
		return ctx, nil
	}
	return h.RequestRouted(ctx)
}

// Call twirp.ServerHooks.ResponsePrepared if the hook is available
func callResponsePrepared(ctx context.Context, h *twirp.ServerHooks) context.Context {
	if h == nil || h.ResponsePrepared == nil {
		return ctx
	}
	return h.ResponsePrepared(ctx)
}

// Call twirp.ServerHooks.ResponseSent if the hook is available
func callResponseSent(ctx context.Context, h *twirp.ServerHooks) {
	if h == nil || h.ResponseSent == nil {
		return
	}
	h.ResponseSent(ctx)
}

// Call twirp.ServerHooks.Error if the hook is available
func callError(ctx context.Context, h *twirp.ServerHooks, err twirp.Error) context.Context {
	if h == nil || h.Error == nil {
		return ctx
	}
	return h.Error(ctx, err)
}

func callClientResponseReceived(ctx context.Context, h *twirp.ClientHooks) {
	if h == nil || h.ResponseReceived == nil {
		return
	}
	h.ResponseReceived(ctx)
}

func callClientRequestPrepared(ctx context.Context, h *twirp.ClientHooks, req *http.Request) (context.Context, error) {
	if h == nil || h.RequestPrepared == nil {
		return ctx, nil
	}
	return h.RequestPrepared(ctx, req)
}

func callClientError(ctx context.Context, h *twirp.ClientHooks, err twirp.Error) {
	if h == nil || h.Error == nil {
		return
	}
	h.Error(ctx, err)
}

var twirpFileDescriptor0 = []byte{
	// 1261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x92, 0xdb, 0x44,
	0x10, 0x8e, 0x64, 0xd9, 0x96, 0xda, 0x6b, 0xaf, 0x32, 0xf9, 0x41, 0x51, 0x08, 0x59, 0x9c, 0x2a,
	0x30, 0xa9, 0x42, 0x1b, 0x0c, 0x81, 0x4a, 0x51, 0x14, 0xe5, 0x2c, 0xa4, 0x30, 0x95, 0x0d, 0x5b,
	0xf2, 0x72, 0x81, 0x83, 0x6b, 0xd6, 0x1a, 0x7b, 0xc5, 0x5a, 0x3f, 0x68, 0xc6, 0x61, 0x7d, 0xe2,
	0xc0, 0x81, 0x97, 0xe0, 0x05, 0x78, 0x0f, 0x2e, 0x3c, 0x07, 0x6f, 0xc0, 0x13, 0x50, 0xf3, 0x23,
	0x5b, 0x5a, 0xd9, 0xec, 0x06, 0x72, 0x92, 0xa6, 0xfb, 0xeb, 0xe9, 0x9e, 0xe9, 0xaf, 0xbb, 0x07,
	0xda, 0x94, 0x64, 0x2f, 0xc3, 0x09, 0xf1, 0xd2, 0x2c, 0x61, 0x09, 0xda, 0x9d, 0x26, 0xe7, 0x1e,
	0x66, 0x91, 0xa7, 0xc4, 0xee, 0xfd, 0x59, 0x92, 0xcc, 0xe6, 0x64, 0x5f, 0xa8, 0x4f, 0x16, 0xd3,
	0x7d, 0x16, 0x46, 0x84, 0x32, 0x1c, 0xa5, 0xd2, 0xa2, 0xfb, 0x15, 0xc0, 0x11, 0x9e, 0x85, 0x31,
	0x66, 0x61, 0x12, 0xa3, 0xfb, 0xd0, 0x8a, 0xc9, 0x39, 0x1b, 0x4f, 0x16, 0x19, 0x4d, 0x32, 0x47,
	0xdb, 0xd3, 0x7a, 0x96, 0x0f, 0x5c, 0x74, 0x20, 0x24, 0xe8, 0x0e, 0x98, 0xa7, 0x98, 0x8e, 0xb9,
	0xc4, 0xd1, 0xf7, 0xb4, 0x9e, 0xe9, 0x37, 0x4f, 0x31, 0x7d, 0x41, 0xce, 0x59, 0xf7, 0x4f, 0x03,
	0x8c, 0xa7, 0x49, 0x72, 0x86, 0x3a, 0xa0, 0x87, 0x81, 0xb2, 0xd5, 0xc3, 0x00, 0x3d, 0x01, 0x98,
	0x64, 0x04, 0x33, 0x12, 0x8c, 0xb1, 0xb4, 0x6a, 0xf5, 0x5d, 0x4f, 0x06, 0xe6, 0xe5, 0x81, 0x79,
	0xc7, 0x79, 0x60, 0xbe, 0xa5, 0xd0, 0x03, 0x86, 0x3e, 0x80, 0x3a, 0x65, 0x98, 0x11, 0xa7, 0xb6,
	0xa7, 0xf5, 0x3a, 0xfd, 0xbb, 0xde, 0x85, 0xf3, 0x79, 0xdc, 0xa1, 0x37, 0xe2, 0x10, 0x5f, 0x22,
	0xf9, 0x11, 0x22, 0x92, 0x4d, 0x4e, 0x71, 0xcc, 0xc6, 0x61, 0xe0, 0x18, 0xf2, 0x08, 0xb9, 0x68,
	0x18, 0xa0, 0xbb, 0x60, 0x9d, 0x64, 0xc9, 0x19, 0xc9, 0xb8, 0xba, 0x2e, 0xd4, 0xa6, 0x14, 0x0c,
	0x03, 0x7e, 0x3e, 0x96, 0xe1, 0x09, 0xe1, 0xba, 0x86, 0xd0, 0x35, 0xc5, 0x7a, 0x18, 0xa0, 0x37,
	0xa0, 0xb9, 0xa0, 0xd2, 0xaa, 0x29, 0x34, 0x0d, 0xbe, 0x1c, 0x06, 0xe8, 0x26, 0xd4, 0x53, 0xbc,
	0x24, 0x99, 0x63, 0x0a, 0xb1, 0x5c, 0xf0, 0x38, 0x68, 0x8c, 0x53, 0x7a, 0x9a, 0x88, 0x38, 0x2c,
	0x19, 0x47, 0x2e, 0x1a, 0x06, 0xe8, 0x1e, 0x40, 0x46, 0x26, 0x24, 0x4c, 0x85, 0x1e, 0x84, 0xde,
	0x52, 0x92, 0x61, 0x80, 0x10, 0x18, 0x11, 0x89, 0x12, 0xa7, 0x25, 0x14, 0xe2, 0x9f, 0x7b, 0x9a,
	0x2e, 0xe2, 0x80, 0x3a, 0x3b, 0xd2, 0x93, 0x58, 0xf0, 0x8d, 0x52, 0xbc, 0x1c, 0xd3, 0x65, 0x74,
	0x92, 0xcc, 0x9d, 0xb6, 0xdc, 0x28, 0xc5, 0xcb, 0x91, 0x10, 0xf0, 0x40, 0xa6, 0xe1, 0x7c, 0x9e,
	0xeb, 0x3b, 0x32, 0x10, 0x2e, 0x52, 0x80, 0xc7, 0x60, 0x52, 0x96, 0x61, 0x46, 0x66, 0x4b, 0x67,
	0x57, 0xdc, 0xf3, 0x9d, 0xca, 0x3d, 0x8f, 0x14, 0xc0, 0x5f, 0x41, 0xc5, 0xb1, 0xb3, 0x70, 0x42,
	0x1c, 0x5b, 0x1d, 0x9b, 0x2f, 0x90, 0x0b, 0x66, 0x10, 0xd2, 0x49, 0xb2, 0x88, 0x99, 0x73, 0x5d,
	0x5e, 0x6e, 0xbe, 0xee, 0xf6, 0xa0, 0x2e, 0x52, 0x85, 0x5a, 0xd0, 0x3c, 0x22, 0x71, 0x10, 0xc6,
	0x33, 0xfb, 0x1a, 0x32, 0xc1, 0x38, 0xc2, 0x61, 0x60, 0x6b, 0xfc, 0xef, 0x8b, 0x24, 0x26, 0xb6,
	0xde, 0xfd, 0xa3, 0x01, 0xf5, 0x6f, 0xb2, 0x80, 0x64, 0xaf, 0x93, 0x4c, 0x4f, 0x00, 0x16, 0x69,
	0x90, 0x9b, 0xd6, 0x2e, 0x37, 0x55, 0xe8, 0x01, 0x43, 0x9f, 0xc1, 0xce, 0x04, 0xc7, 0x13, 0x32,
	0x9f, 0x4b, 0x63, 0xe3, 0x52, 0xe3, 0xd6, 0x0a, 0x3f, 0x60, 0x45, 0xea, 0xd4, 0x4b, 0xd4, 0xb9,
	0x40, 0xd6, 0x46, 0x85, 0xac, 0xfd, 0xbc, 0x00, 0x9a, 0x22, 0x31, 0x6f, 0x56, 0x12, 0x23, 0x6e,
	0xa9, 0x5c, 0x01, 0x65, 0x3e, 0x98, 0x97, 0xf0, 0xc1, 0xaa, 0xf0, 0xe1, 0x36, 0x34, 0x94, 0x4e,
	0x92, 0x52, 0xad, 0xd0, 0x7b, 0x60, 0xd0, 0x30, 0x20, 0x82, 0x91, 0x9d, 0xfe, 0xad, 0x2a, 0x47,
	0xc2, 0x80, 0xf8, 0x02, 0x52, 0xa2, 0xd4, 0xce, 0x7f, 0xa0, 0x54, 0x7b, 0x1b, 0xa5, 0x3a, 0x65,
	0x4a, 0xad, 0x2b, 0x62, 0xb7, 0x58, 0x11, 0x6f, 0xc3, 0x0e, 0x3f, 0x0f, 0x09, 0xc6, 0x52, 0x29,
	0x19, 0xda, 0x92, 0xb2, 0x67, 0x02, 0xf2, 0x00, 0xda, 0x0a, 0x82, 0xa3, 0x02, 0x59, 0x95, 0xdd,
	0x40, 0xc8, 0x90, 0x07, 0x37, 0xc8, 0x39, 0xcb, 0xf0, 0xb8, 0x0c, 0x45, 0x02, 0x7a, 0x5d, 0xa8,
	0x9e, 0x15, 0xf1, 0xf7, 0x00, 0xa6, 0x84, 0xe4, 0xb0, 0x1b, 0xf2, 0xe6, 0xa7, 0x84, 0x28, 0xf5,
	0x03, 0x68, 0xe3, 0x97, 0x24, 0xc3, 0x33, 0x32, 0x96, 0xc7, 0xbc, 0x29, 0x7d, 0x2a, 0xe1, 0x11,
	0x97, 0x75, 0xbf, 0x2e, 0x14, 0xc9, 0x71, 0x86, 0x55, 0x91, 0x00, 0x34, 0xa4, 0x27, 0x5b, 0x43,
	0x6d, 0xb0, 0x0e, 0x72, 0x72, 0xd9, 0x3a, 0xda, 0x01, 0xd3, 0x27, 0x3f, 0x90, 0x09, 0x23, 0x81,
	0x5d, 0x13, 0x56, 0x61, 0x44, 0x92, 0x05, 0xb3, 0x8d, 0xee, 0xdf, 0x3a, 0xb4, 0x04, 0x41, 0x7c,
	0x92, 0x26, 0x19, 0x7b, 0x9d, 0xc5, 0x84, 0xc0, 0x08, 0xf2, 0xc6, 0x6c, 0xf9, 0xe2, 0xbf, 0x48,
	0x73, 0xe3, 0xdf, 0x68, 0x5e, 0xaf, 0xd0, 0x7c, 0x4d, 0xb9, 0x46, 0x89, 0x72, 0x48, 0x51, 0x4e,
	0x36, 0x5c, 0xc9, 0xad, 0x4a, 0xe6, 0xcc, 0x0d, 0x99, 0x7b, 0x17, 0x76, 0x93, 0x13, 0x86, 0xc3,
	0x78, 0x0d, 0x93, 0x44, 0xef, 0xe4, 0xe2, 0x8d, 0x29, 0x83, 0x8b, 0x29, 0xbb, 0x0b, 0x96, 0x50,
	0x53, 0x4a, 0x98, 0x6a, 0xc5, 0x26, 0xd7, 0xf2, 0x35, 0x27, 0x9f, 0x64, 0x25, 0xa7, 0x78, 0xdd,
	0x97, 0x8b, 0xee, 0xf7, 0x60, 0x7c, 0x4b, 0x37, 0x74, 0x2e, 0x04, 0x46, 0x8c, 0x23, 0x22, 0xae,
	0xd9, 0xf2, 0xc5, 0x3f, 0x97, 0x65, 0xc9, 0x7c, 0x75, 0x8b, 0xfc, 0xbf, 0x3c, 0x9f, 0x8c, 0xf2,
	0x7c, 0xea, 0xde, 0x86, 0x0e, 0xdf, 0x7c, 0x24, 0x6b, 0xc8, 0x27, 0x3f, 0xba, 0x06, 0xe8, 0x87,
	0xa4, 0xfb, 0x8b, 0x01, 0xe8, 0x50, 0xdd, 0x67, 0x41, 0x39, 0x06, 0x38, 0x10, 0x29, 0x13, 0x83,
	0xb9, 0x38, 0xdc, 0xb4, 0xad, 0xc3, 0x4d, 0x2f, 0xa5, 0xae, 0x3c, 0xa5, 0x6a, 0x17, 0xa6, 0x94,
	0xfb, 0x0e, 0x58, 0x3e, 0xc1, 0x81, 0xec, 0xd5, 0xdb, 0xf7, 0x77, 0xff, 0xd2, 0x00, 0x9e, 0x87,
	0x94, 0x09, 0x20, 0x2d, 0xe4, 0x5b, 0xdb, 0x98, 0x6f, 0xbd, 0x90, 0x6f, 0xb7, 0xd0, 0x4b, 0xa4,
	0xff, 0x52, 0xc3, 0x90, 0xed, 0x51, 0xde, 0x93, 0x6a, 0x80, 0x5b, 0xdb, 0xed, 0x23, 0xa8, 0x27,
	0x3c, 0x00, 0xc1, 0xb2, 0x4e, 0xdf, 0xad, 0xf6, 0xa4, 0x24, 0x93, 0x21, 0xfa, 0x12, 0xc8, 0x03,
	0x55, 0x6f, 0x21, 0x35, 0xf3, 0xe5, 0x8a, 0x3b, 0x9e, 0x87, 0x51, 0x28, 0xc9, 0x57, 0xf3, 0xe5,
	0xc2, 0xed, 0x41, 0x4b, 0x56, 0xe6, 0xa5, 0xf7, 0x71, 0x0c, 0xf6, 0xea, 0x3a, 0x64, 0x71, 0xd2,
	0x55, 0x49, 0x69, 0x85, 0x92, 0x5a, 0xfb, 0xd7, 0x37, 0xfb, 0xaf, 0x15, 0xfc, 0x77, 0x7f, 0xd7,
	0xe1, 0x46, 0x85, 0x05, 0x34, 0x75, 0xdb, 0xa5, 0xb8, 0xdc, 0x65, 0x29, 0x17, 0x1e, 0x34, 0xc4,
	0x59, 0xa9, 0xa3, 0xed, 0xd5, 0x7a, 0xad, 0xfe, 0xed, 0xcd, 0x33, 0xc6, 0x57, 0x28, 0xf4, 0x29,
	0x1f, 0x2f, 0xf9, 0x8b, 0x51, 0x35, 0x8d, 0xea, 0xc3, 0x6c, 0xfd, 0xa8, 0xf4, 0x0b, 0x70, 0xf7,
	0x57, 0x6d, 0xc3, 0xc1, 0x3f, 0x86, 0x66, 0x26, 0x7f, 0x55, 0x08, 0x5b, 0xc6, 0x9c, 0xc4, 0xfb,
	0x39, 0xf8, 0x7f, 0x45, 0xf2, 0xd0, 0x01, 0x83, 0x0f, 0x2c, 0xd4, 0x84, 0xda, 0x80, 0x9e, 0xd9,
	0xd7, 0xf8, 0xcf, 0x53, 0xfe, 0x0c, 0x79, 0xf8, 0x09, 0x98, 0xf9, 0x6c, 0xe2, 0x7d, 0xf7, 0x10,
	0x67, 0x67, 0x84, 0xd9, 0xd7, 0x90, 0x05, 0xf5, 0xe7, 0xfc, 0x9a, 0x6d, 0x4d, 0xb4, 0xe3, 0x64,
	0x3e, 0x4f, 0x7e, 0xb2, 0x75, 0xfe, 0x6a, 0x79, 0x36, 0x27, 0xe7, 0x76, 0xed, 0xe1, 0x5b, 0x60,
	0xad, 0x08, 0x24, 0x1e, 0x33, 0x5f, 0x8e, 0x0e, 0xe4, 0xc6, 0x83, 0xd1, 0x81, 0xad, 0xf5, 0x5f,
	0x40, 0xab, 0x50, 0xbc, 0xe8, 0x73, 0x5e, 0xb9, 0xa8, 0x5b, 0x09, 0xb8, 0x5c, 0xe0, 0xde, 0x21,
	0x71, 0x6f, 0x6d, 0xc4, 0xf4, 0x7f, 0x33, 0x60, 0xf7, 0x42, 0xba, 0xd1, 0xa8, 0x54, 0xf1, 0x5e,
	0xc5, 0xb0, 0xda, 0x24, 0xbc, 0x35, 0x7e, 0x83, 0x23, 0xb1, 0x8d, 0x5f, 0xac, 0xf2, 0xf7, 0xaf,
	0xb2, 0xe7, 0x0a, 0xee, 0x6e, 0xa1, 0x13, 0x8a, 0xca, 0x24, 0xbc, 0xca, 0xa6, 0x6b, 0xbc, 0xbb,
	0x7f, 0x39, 0x9e, 0xa6, 0x05, 0x03, 0xf4, 0xf3, 0x06, 0xde, 0x7d, 0xf4, 0x4a, 0x4e, 0x95, 0x95,
	0xfb, 0xf8, 0xd5, 0x5c, 0xe7, 0xce, 0xd2, 0x72, 0x6f, 0xd8, 0xbf, 0x52, 0x66, 0x0a, 0x45, 0xfb,
	0xe8, 0x4a, 0x6e, 0x0b, 0x16, 0x4f, 0x9b, 0xdf, 0xd5, 0xe5, 0x0c, 0x6f, 0x88, 0xcf, 0x87, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xba, 0x1a, 0xea, 0x2d, 0x0e, 0x00, 0x00,
}
