package transaction

import (
	"net/http"
	"example/online_shop/go/lib"
	//"encoding/json"

	//"github.com/gorilla/mux"
	//"gorm.io/gorm"
)

type Transaction struct {
	lib.ApiBase
}

// Action Index
func (t Transaction) ApiIdx() lib.ApiHandler {
	return lib.ApiHandlerFunc(func (ar *lib.ApiRequest) *lib.ApiResponse {
		// method := r.Method
		// params := ar.Params
		query := ar.Query

		/*
		ar.Wri.Header().Set("Content-Type", "text/plain")
		ar.Wri.WriteHeader(http.StatusOK)
		ar.Wri.Write([]byte(params["id"] + "yeah"))
		*/

		return &lib.ApiResponse{
			Headers: lib.ApiResHeaders{
				lib.ApiResHeader{ "Content-Type", "text/plain" },
			},
			ResponseStatus: http.StatusOK,
			ResponseBody: []byte(query.Get("test") + "your bigboss"),
		}
	})
}

// Action Show
func (t Transaction) ApiShow() lib.ApiHandler {
	return lib.ApiHandlerFunc(func (ar *lib.ApiRequest) *lib.ApiResponse {
		return nil
	})
}

// Action Create
func (t Transaction) ApiCreate() lib.ApiHandler {
	return lib.ApiHandlerFunc(func (ar *lib.ApiRequest) *lib.ApiResponse {
		return nil
	})
}

// Action Update
func (t Transaction) ApiUpdate() lib.ApiHandler  {
	return lib.ApiHandlerFunc(func (ar *lib.ApiRequest) *lib.ApiResponse {
		return nil
	})
}

// Action Destroy
func (t Transaction) ApiDestroy() lib.ApiHandler  {
	return lib.ApiHandlerFunc(func (ar *lib.ApiRequest) *lib.ApiResponse {
		return nil
	})
}

/* 特殊なルーティングを設定したいときの書き方!!
func testHandler() lib.ApiHandler {
	return lib.ApiHandlerFunc(func (ar *lib.ApiRequest) {})
}

// Define additional routes
func (t Transaction) ApiOtherRouter() lib.ApiAddRoutes {
	return lib.ApiAddRoutes {
		lib.AddRoute("/path", testHandler(), "GET"),
		// router.HandleFunc("your_path_here", your_method).Methods("HTTP_METHOD")
	}
}
*/

func New() *Transaction {
	return &Transaction{
		ApiBase: lib.ApiBase{
			ResourceName: "transactions",
		},
	}
}
