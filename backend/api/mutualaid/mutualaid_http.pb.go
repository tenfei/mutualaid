// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.2

package mutualaid

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type UserServiceHTTPServer interface {
	ActiveUser(context.Context, *ActiveUserReq) (*ActiveUserResp, error)
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	JSAPISign(context.Context, *JSAPISignReq) (*JSAPISignResp, error)
	WxLogin(context.Context, *WxLoginReq) (*WxLoginResp, error)
	WxOAuth2(context.Context, *WxOAuth2Req) (*WxOAuth2Resp, error)
	WxPhoneNumber(context.Context, *WxPhoneNumberReq) (*WxPhoneNumberResp, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/wxoauth2", _UserService_WxOAuth20_HTTP_Handler(srv))
	r.POST("/wxlogin", _UserService_WxLogin0_HTTP_Handler(srv))
	r.POST("/activeuser", _UserService_ActiveUser0_HTTP_Handler(srv))
	r.GET("/wx_phone_number", _UserService_WxPhoneNumber0_HTTP_Handler(srv))
	r.GET("/user", _UserService_GetUser0_HTTP_Handler(srv))
	r.POST("/js_api_sign", _UserService_JSAPISign0_HTTP_Handler(srv))
}

func _UserService_WxOAuth20_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxOAuth2Req
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserService/WxOAuth2")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxOAuth2(ctx, req.(*WxOAuth2Req))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxOAuth2Resp)
		return ctx.Result(200, reply)
	}
}

func _UserService_WxLogin0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxLoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserService/WxLogin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxLogin(ctx, req.(*WxLoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxLoginResp)
		return ctx.Result(200, reply)
	}
}

func _UserService_ActiveUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ActiveUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserService/ActiveUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ActiveUser(ctx, req.(*ActiveUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ActiveUserResp)
		return ctx.Result(200, reply)
	}
}

func _UserService_WxPhoneNumber0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxPhoneNumberReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserService/WxPhoneNumber")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxPhoneNumber(ctx, req.(*WxPhoneNumberReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxPhoneNumberResp)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserService/GetUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserResp)
		return ctx.Result(200, reply)
	}
}

func _UserService_JSAPISign0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in JSAPISignReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserService/JSAPISign")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.JSAPISign(ctx, req.(*JSAPISignReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*JSAPISignResp)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	ActiveUser(ctx context.Context, req *ActiveUserReq, opts ...http.CallOption) (rsp *ActiveUserResp, err error)
	GetUser(ctx context.Context, req *GetUserReq, opts ...http.CallOption) (rsp *GetUserResp, err error)
	JSAPISign(ctx context.Context, req *JSAPISignReq, opts ...http.CallOption) (rsp *JSAPISignResp, err error)
	WxLogin(ctx context.Context, req *WxLoginReq, opts ...http.CallOption) (rsp *WxLoginResp, err error)
	WxOAuth2(ctx context.Context, req *WxOAuth2Req, opts ...http.CallOption) (rsp *WxOAuth2Resp, err error)
	WxPhoneNumber(ctx context.Context, req *WxPhoneNumberReq, opts ...http.CallOption) (rsp *WxPhoneNumberResp, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) ActiveUser(ctx context.Context, in *ActiveUserReq, opts ...http.CallOption) (*ActiveUserResp, error) {
	var out ActiveUserResp
	pattern := "/activeuser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.UserService/ActiveUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) GetUser(ctx context.Context, in *GetUserReq, opts ...http.CallOption) (*GetUserResp, error) {
	var out GetUserResp
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.UserService/GetUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) JSAPISign(ctx context.Context, in *JSAPISignReq, opts ...http.CallOption) (*JSAPISignResp, error) {
	var out JSAPISignResp
	pattern := "/js_api_sign"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.UserService/JSAPISign"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) WxLogin(ctx context.Context, in *WxLoginReq, opts ...http.CallOption) (*WxLoginResp, error) {
	var out WxLoginResp
	pattern := "/wxlogin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.UserService/WxLogin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) WxOAuth2(ctx context.Context, in *WxOAuth2Req, opts ...http.CallOption) (*WxOAuth2Resp, error) {
	var out WxOAuth2Resp
	pattern := "/wxoauth2"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.UserService/WxOAuth2"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) WxPhoneNumber(ctx context.Context, in *WxPhoneNumberReq, opts ...http.CallOption) (*WxPhoneNumberResp, error) {
	var out WxPhoneNumberResp
	pattern := "/wx_phone_number"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.UserService/WxPhoneNumber"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type MutualAidQueryHTTPServer interface {
	Discovery(context.Context, *DiscoveryReq) (*DiscoveryResp, error)
}

func RegisterMutualAidQueryHTTPServer(s *http.Server, srv MutualAidQueryHTTPServer) {
	r := s.Route("/")
	r.GET("/discovery", _MutualAidQuery_Discovery0_HTTP_Handler(srv))
}

func _MutualAidQuery_Discovery0_HTTP_Handler(srv MutualAidQueryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DiscoveryReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.MutualAidQuery/Discovery")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Discovery(ctx, req.(*DiscoveryReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DiscoveryResp)
		return ctx.Result(200, reply)
	}
}

type MutualAidQueryHTTPClient interface {
	Discovery(ctx context.Context, req *DiscoveryReq, opts ...http.CallOption) (rsp *DiscoveryResp, err error)
}

type MutualAidQueryHTTPClientImpl struct {
	cc *http.Client
}

func NewMutualAidQueryHTTPClient(client *http.Client) MutualAidQueryHTTPClient {
	return &MutualAidQueryHTTPClientImpl{client}
}

func (c *MutualAidQueryHTTPClientImpl) Discovery(ctx context.Context, in *DiscoveryReq, opts ...http.CallOption) (*DiscoveryResp, error) {
	var out DiscoveryResp
	pattern := "/discovery"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.MutualAidQuery/Discovery"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type ExamineAidHTTPServer interface {
	BlockUser(context.Context, *BlockUserReq) (*BlockUserResp, error)
	ExamineAid(context.Context, *ExamineAidReq) (*ExamineAidResp, error)
	ExamineLogin(context.Context, *ExamineLoginReq) (*ExamineLoginResp, error)
	GetBlockUserList(context.Context, *GetBlockUserListReq) (*GetBlockUserListResp, error)
	GetExamineList(context.Context, *GetExamineListReq) (*ExamineListResp, error)
	PassUser(context.Context, *PassUserReq) (*PassUserResp, error)
}

func RegisterExamineAidHTTPServer(s *http.Server, srv ExamineAidHTTPServer) {
	r := s.Route("/")
	r.POST("/examine_login", _ExamineAid_ExamineLogin0_HTTP_Handler(srv))
	r.GET("/examine/aid", _ExamineAid_GetExamineList0_HTTP_Handler(srv))
	r.POST("/examine/aid/{id}", _ExamineAid_ExamineAid0_HTTP_Handler(srv))
	r.DELETE("/examine/user/{id}", _ExamineAid_BlockUser0_HTTP_Handler(srv))
	r.GET("/examine/blockuser", _ExamineAid_GetBlockUserList0_HTTP_Handler(srv))
	r.POST("/examine/user/{id}", _ExamineAid_PassUser0_HTTP_Handler(srv))
}

func _ExamineAid_ExamineLogin0_HTTP_Handler(srv ExamineAidHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExamineLoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.ExamineAid/ExamineLogin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExamineLogin(ctx, req.(*ExamineLoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExamineLoginResp)
		return ctx.Result(200, reply)
	}
}

func _ExamineAid_GetExamineList0_HTTP_Handler(srv ExamineAidHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetExamineListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.ExamineAid/GetExamineList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetExamineList(ctx, req.(*GetExamineListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExamineListResp)
		return ctx.Result(200, reply)
	}
}

func _ExamineAid_ExamineAid0_HTTP_Handler(srv ExamineAidHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExamineAidReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.ExamineAid/ExamineAid")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExamineAid(ctx, req.(*ExamineAidReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExamineAidResp)
		return ctx.Result(200, reply)
	}
}

func _ExamineAid_BlockUser0_HTTP_Handler(srv ExamineAidHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BlockUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.ExamineAid/BlockUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BlockUser(ctx, req.(*BlockUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BlockUserResp)
		return ctx.Result(200, reply)
	}
}

func _ExamineAid_GetBlockUserList0_HTTP_Handler(srv ExamineAidHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBlockUserListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.ExamineAid/GetBlockUserList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBlockUserList(ctx, req.(*GetBlockUserListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBlockUserListResp)
		return ctx.Result(200, reply)
	}
}

func _ExamineAid_PassUser0_HTTP_Handler(srv ExamineAidHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PassUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.ExamineAid/PassUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PassUser(ctx, req.(*PassUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PassUserResp)
		return ctx.Result(200, reply)
	}
}

type ExamineAidHTTPClient interface {
	BlockUser(ctx context.Context, req *BlockUserReq, opts ...http.CallOption) (rsp *BlockUserResp, err error)
	ExamineAid(ctx context.Context, req *ExamineAidReq, opts ...http.CallOption) (rsp *ExamineAidResp, err error)
	ExamineLogin(ctx context.Context, req *ExamineLoginReq, opts ...http.CallOption) (rsp *ExamineLoginResp, err error)
	GetBlockUserList(ctx context.Context, req *GetBlockUserListReq, opts ...http.CallOption) (rsp *GetBlockUserListResp, err error)
	GetExamineList(ctx context.Context, req *GetExamineListReq, opts ...http.CallOption) (rsp *ExamineListResp, err error)
	PassUser(ctx context.Context, req *PassUserReq, opts ...http.CallOption) (rsp *PassUserResp, err error)
}

type ExamineAidHTTPClientImpl struct {
	cc *http.Client
}

func NewExamineAidHTTPClient(client *http.Client) ExamineAidHTTPClient {
	return &ExamineAidHTTPClientImpl{client}
}

func (c *ExamineAidHTTPClientImpl) BlockUser(ctx context.Context, in *BlockUserReq, opts ...http.CallOption) (*BlockUserResp, error) {
	var out BlockUserResp
	pattern := "/examine/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.ExamineAid/BlockUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExamineAidHTTPClientImpl) ExamineAid(ctx context.Context, in *ExamineAidReq, opts ...http.CallOption) (*ExamineAidResp, error) {
	var out ExamineAidResp
	pattern := "/examine/aid/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.ExamineAid/ExamineAid"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExamineAidHTTPClientImpl) ExamineLogin(ctx context.Context, in *ExamineLoginReq, opts ...http.CallOption) (*ExamineLoginResp, error) {
	var out ExamineLoginResp
	pattern := "/examine_login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.ExamineAid/ExamineLogin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExamineAidHTTPClientImpl) GetBlockUserList(ctx context.Context, in *GetBlockUserListReq, opts ...http.CallOption) (*GetBlockUserListResp, error) {
	var out GetBlockUserListResp
	pattern := "/examine/blockuser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.ExamineAid/GetBlockUserList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExamineAidHTTPClientImpl) GetExamineList(ctx context.Context, in *GetExamineListReq, opts ...http.CallOption) (*ExamineListResp, error) {
	var out ExamineListResp
	pattern := "/examine/aid"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.ExamineAid/GetExamineList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExamineAidHTTPClientImpl) PassUser(ctx context.Context, in *PassUserReq, opts ...http.CallOption) (*PassUserResp, error) {
	var out PassUserResp
	pattern := "/examine/user/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.ExamineAid/PassUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type UserAidQueryHTTPServer interface {
	GetAidDetail(context.Context, *GetAidDetailReq) (*GetAidDetailResp, error)
	ListAidNeeds(context.Context, *ListAidNeedsReq) (*ListAidNeedsResp, error)
	ListAidOffered(context.Context, *ListAidOfferedReq) (*ListAidOfferedResp, error)
}

func RegisterUserAidQueryHTTPServer(s *http.Server, srv UserAidQueryHTTPServer) {
	r := s.Route("/")
	r.GET("/user/aid/offered", _UserAidQuery_ListAidOffered0_HTTP_Handler(srv))
	r.GET("/user/aid/needs", _UserAidQuery_ListAidNeeds0_HTTP_Handler(srv))
	r.GET("/user/aid/{id}", _UserAidQuery_GetAidDetail0_HTTP_Handler(srv))
}

func _UserAidQuery_ListAidOffered0_HTTP_Handler(srv UserAidQueryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAidOfferedReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserAidQuery/ListAidOffered")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAidOffered(ctx, req.(*ListAidOfferedReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAidOfferedResp)
		return ctx.Result(200, reply)
	}
}

func _UserAidQuery_ListAidNeeds0_HTTP_Handler(srv UserAidQueryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAidNeedsReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserAidQuery/ListAidNeeds")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAidNeeds(ctx, req.(*ListAidNeedsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAidNeedsResp)
		return ctx.Result(200, reply)
	}
}

func _UserAidQuery_GetAidDetail0_HTTP_Handler(srv UserAidQueryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAidDetailReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserAidQuery/GetAidDetail")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAidDetail(ctx, req.(*GetAidDetailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAidDetailResp)
		return ctx.Result(200, reply)
	}
}

type UserAidQueryHTTPClient interface {
	GetAidDetail(ctx context.Context, req *GetAidDetailReq, opts ...http.CallOption) (rsp *GetAidDetailResp, err error)
	ListAidNeeds(ctx context.Context, req *ListAidNeedsReq, opts ...http.CallOption) (rsp *ListAidNeedsResp, err error)
	ListAidOffered(ctx context.Context, req *ListAidOfferedReq, opts ...http.CallOption) (rsp *ListAidOfferedResp, err error)
}

type UserAidQueryHTTPClientImpl struct {
	cc *http.Client
}

func NewUserAidQueryHTTPClient(client *http.Client) UserAidQueryHTTPClient {
	return &UserAidQueryHTTPClientImpl{client}
}

func (c *UserAidQueryHTTPClientImpl) GetAidDetail(ctx context.Context, in *GetAidDetailReq, opts ...http.CallOption) (*GetAidDetailResp, error) {
	var out GetAidDetailResp
	pattern := "/user/aid/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.UserAidQuery/GetAidDetail"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserAidQueryHTTPClientImpl) ListAidNeeds(ctx context.Context, in *ListAidNeedsReq, opts ...http.CallOption) (*ListAidNeedsResp, error) {
	var out ListAidNeedsResp
	pattern := "/user/aid/needs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.UserAidQuery/ListAidNeeds"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserAidQueryHTTPClientImpl) ListAidOffered(ctx context.Context, in *ListAidOfferedReq, opts ...http.CallOption) (*ListAidOfferedResp, error) {
	var out ListAidOfferedResp
	pattern := "/user/aid/offered"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.UserAidQuery/ListAidOffered"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type UserAidManagerHTTPServer interface {
	CancelAid(context.Context, *CancelAidReq) (*CancelAidResp, error)
	CreateAid(context.Context, *CreateAidReq) (*CreateAidResp, error)
	CreateAidMessage(context.Context, *CreateAidMessageReq) (*CreateAidMessageResp, error)
	FinishAid(context.Context, *FinishAidReq) (*FinishAidResp, error)
}

func RegisterUserAidManagerHTTPServer(s *http.Server, srv UserAidManagerHTTPServer) {
	r := s.Route("/")
	r.POST("/user/aid", _UserAidManager_CreateAid0_HTTP_Handler(srv))
	r.DELETE("/user/aid/{id}", _UserAidManager_CancelAid0_HTTP_Handler(srv))
	r.PUT("/user/aid/{id}", _UserAidManager_FinishAid0_HTTP_Handler(srv))
	r.POST("/user/aid/message", _UserAidManager_CreateAidMessage0_HTTP_Handler(srv))
}

func _UserAidManager_CreateAid0_HTTP_Handler(srv UserAidManagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAidReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserAidManager/CreateAid")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAid(ctx, req.(*CreateAidReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAidResp)
		return ctx.Result(200, reply)
	}
}

func _UserAidManager_CancelAid0_HTTP_Handler(srv UserAidManagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CancelAidReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserAidManager/CancelAid")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CancelAid(ctx, req.(*CancelAidReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CancelAidResp)
		return ctx.Result(200, reply)
	}
}

func _UserAidManager_FinishAid0_HTTP_Handler(srv UserAidManagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FinishAidReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserAidManager/FinishAid")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FinishAid(ctx, req.(*FinishAidReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FinishAidResp)
		return ctx.Result(200, reply)
	}
}

func _UserAidManager_CreateAidMessage0_HTTP_Handler(srv UserAidManagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAidMessageReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.mutualaid.UserAidManager/CreateAidMessage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAidMessage(ctx, req.(*CreateAidMessageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAidMessageResp)
		return ctx.Result(200, reply)
	}
}

type UserAidManagerHTTPClient interface {
	CancelAid(ctx context.Context, req *CancelAidReq, opts ...http.CallOption) (rsp *CancelAidResp, err error)
	CreateAid(ctx context.Context, req *CreateAidReq, opts ...http.CallOption) (rsp *CreateAidResp, err error)
	CreateAidMessage(ctx context.Context, req *CreateAidMessageReq, opts ...http.CallOption) (rsp *CreateAidMessageResp, err error)
	FinishAid(ctx context.Context, req *FinishAidReq, opts ...http.CallOption) (rsp *FinishAidResp, err error)
}

type UserAidManagerHTTPClientImpl struct {
	cc *http.Client
}

func NewUserAidManagerHTTPClient(client *http.Client) UserAidManagerHTTPClient {
	return &UserAidManagerHTTPClientImpl{client}
}

func (c *UserAidManagerHTTPClientImpl) CancelAid(ctx context.Context, in *CancelAidReq, opts ...http.CallOption) (*CancelAidResp, error) {
	var out CancelAidResp
	pattern := "/user/aid/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.mutualaid.UserAidManager/CancelAid"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserAidManagerHTTPClientImpl) CreateAid(ctx context.Context, in *CreateAidReq, opts ...http.CallOption) (*CreateAidResp, error) {
	var out CreateAidResp
	pattern := "/user/aid"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.UserAidManager/CreateAid"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserAidManagerHTTPClientImpl) CreateAidMessage(ctx context.Context, in *CreateAidMessageReq, opts ...http.CallOption) (*CreateAidMessageResp, error) {
	var out CreateAidMessageResp
	pattern := "/user/aid/message"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.UserAidManager/CreateAidMessage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserAidManagerHTTPClientImpl) FinishAid(ctx context.Context, in *FinishAidReq, opts ...http.CallOption) (*FinishAidResp, error) {
	var out FinishAidResp
	pattern := "/user/aid/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.mutualaid.UserAidManager/FinishAid"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}