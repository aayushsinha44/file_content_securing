package backend
import(
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"Add Owner",
		"POST",
		"/AddOwner",
		endpoints.AddOwner,
	},
	Route{
		"Get Owner List",
		"GET",
		"/GetOwnerList",
		endpoints.GetOwnerList,
	},
	Route{
		"Get File Name",
		"POST",
		"/GetFileName",
		endpoints.GetFileName,
	},
	Route{
		"Create New File",
		"POST",
		"/CreateNewFile",
		endpoints.CreateNewFile,
	},
	Route{
		"Upload file on IPFS",
		"POST",
		"/UploadFile",
		endpoints.UploadFile
	}
}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}