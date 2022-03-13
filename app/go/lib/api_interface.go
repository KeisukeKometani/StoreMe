package lib

import (
	"net/http"
	"net/url"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io"
)

type ApiRouteConfig struct {
	index   bool
	show    bool
	create  bool
	update  bool
	destroy bool
}

func NewApiRouteConfig() ApiRouteConfig {
	u := ApiRouteConfig{}
	u.index = true
	u.show = true
	u.create = true
	u.update = true
	u.destroy = true
	return u
}

type ApiHandler interface {
	ApiServe(*ApiRequest) *ApiResponse
}
type ApiHandlerFunc func(*ApiRequest) *ApiResponse
func (f ApiHandlerFunc) ApiServe(ar *ApiRequest) *ApiResponse {
	return f(ar)
}

type ApiAddRoutes []func(*mux.Router)

type ApiInterface interface {
	ApiDBInit()
	ApiDefActs()     ApiRouteConfig

	ApiIdx()         ApiHandler
	ApiShow()        ApiHandler
	ApiCreate()      ApiHandler
	ApiUpdate()      ApiHandler
	ApiDestroy()     ApiHandler

	ApiIdxPath()     string
	ApiShowPath()    string
	ApiCUDPath()     string
	ApiOtherRouter() ApiAddRoutes
}

func AddRoute(path string, handler ApiHandler, methods ...string) func (*mux.Router){
	return func (r *mux.Router){
		r.HandleFunc(path, handlerDispatcher(handler)).Methods(methods...)
	}
}

func addRouteNoop() func(*mux.Router){
	return func (r *mux.Router){
	}
}

// Define Routes
// require implement ApiInterface all methods.
func DefineRoutes(r *mux.Router, intf ApiInterface){
	intf.ApiDBInit()
	conf := intf.ApiDefActs()

	if conf.index == true {
		r.HandleFunc(intf.ApiIdxPath(), handlerDispatcher(intf.ApiIdx())).Methods("GET")
	}

	if conf.show == true {
		r.HandleFunc(intf.ApiShowPath(), handlerDispatcher(intf.ApiShow())).Methods("GET")
	}

	if conf.create == true {
		r.HandleFunc(intf.ApiCUDPath(), handlerDispatcher(intf.ApiCreate())).Methods("POST")
	}

	if conf.update == true {
		r.HandleFunc(intf.ApiCUDPath(), handlerDispatcher(intf.ApiUpdate())).Methods("PUT")
	}

	if conf.destroy == true {
		r.HandleFunc(intf.ApiCUDPath(), handlerDispatcher(intf.ApiDestroy())).Methods("DELETE")
	}
	
	// その他のカスタムルート設定を呼び出す
	for _,v := range intf.ApiOtherRouter() {
		v(r)
	}
}

type ApiRequest struct {
	Params   map[string]string
	Query    url.Values
	JsonBody interface{}
	ByteBody io.Reader
	Req      *http.Request
	Wri      http.ResponseWriter
}

type ApiResHeader struct {Name string; Value string}
type ApiResHeaders []ApiResHeader
type ApiResponse struct {
	Headers        ApiResHeaders
	ResponseStatus int
	ResponseBody   []byte
}

// handlerDispatcher は、ApiHandlerを起動するための関数を返す関数
// パラメータなどを事前にパースして渡します
func handlerDispatcher(handler ApiHandler) func(http.ResponseWriter, *http.Request){
	return func (w http.ResponseWriter, r *http.Request) {
		param := &ApiRequest{}
		param.Query = r.URL.Query()
		param.Params = mux.Vars(r)
		param.ByteBody = r.Body
		// param.jsonBody
		param.Req = r
		param.Wri = w

		res := handler.ApiServe(param)
		responseBuild(res, w, r)
	}
}

func responseBuild(res *ApiResponse, w http.ResponseWriter, r *http.Request) {
		if res == nil {
			responseBody := "The routing for this URL is defined,\n" +
											"but there is no response :(\n\n" +
											"Path:"   + r.URL.Path     + "\n" +
			                "Query:"  + r.URL.RawQuery + "\n" +
											"Method:" + r.Method       + "\n" +
											"Addr:"   + r.RemoteAddr   + "\n"
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(responseBody))
		}else{
			// ヘッダをセット
			if res.Headers != nil {
				for _,v := range res.Headers {
					w.Header().Set(v.Name, v.Value)
				}
			}

			// レスポンスステータスをセット
			if res.ResponseStatus != 0 {
				w.WriteHeader(res.ResponseStatus)
			}

			// レスポンス本文をセット
			if res.ResponseBody != nil {
				w.Write(res.ResponseBody)
			}
		}
}

//
// Api Base
// これを移譲してルートを実装すれば楽なはず...?
//
type ApiBase struct {
	DB *gorm.DB
	ResourceName string
}

// ルーティング前にデータベースコネクションを初期化
func (t ApiBase) ApiDBInit(){
	t.DB = GetDBInstance()
}

// デフォルト定義したアクションの有効無効の設定値を返す
func (t ApiBase) ApiDefActs() ApiRouteConfig {
	return NewApiRouteConfig()
}

// Action Index
func (t ApiBase) ApiIdx() ApiHandler {
	return ApiHandlerFunc(func (ar *ApiRequest) *ApiResponse { return nil })
}

// Action Show
func (t ApiBase) ApiShow() ApiHandler {
	return ApiHandlerFunc(func (ar *ApiRequest) *ApiResponse { return nil })
}

// Action Create
func (t ApiBase) ApiCreate() ApiHandler {
	return ApiHandlerFunc(func (ar *ApiRequest) *ApiResponse { return nil })
}

// Action Update
func (t ApiBase) ApiUpdate() ApiHandler {
	return ApiHandlerFunc(func (ar *ApiRequest) *ApiResponse { return nil })
}

// Action Destroy
func (t ApiBase) ApiDestroy() ApiHandler {
	return ApiHandlerFunc(func (ar *ApiRequest) *ApiResponse{ return nil })
}

// Define additional routes
func (t ApiBase) ApiOtherRouter() ApiAddRoutes{
	return ApiAddRoutes{ addRouteNoop(),
		// addRoute("path here", handler, "GET"),
	}
}

func (t ApiBase) ApiIdxPath() string {
	return "/" + t.ResourceName
}

func (t ApiBase) ApiShowPath() string {
	return "/" + t.ResourceName + "/{id}"
}

func (t ApiBase) ApiCUDPath() string {
	return "/" + t.ResourceName + "/{id}"
}

func NewApiBase(resourceName string, db *gorm.DB) *ApiBase{
	return &ApiBase{ResourceName: resourceName, DB: db}
}
