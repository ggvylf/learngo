package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/_geecache/"

type HTTPPool struct {
	self     string
	basePath string
}

func NewHTTPPool(self string) *HTTPPool {

	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

//自定义日志格式
func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//前缀basepath是否匹配
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}

	p.Log("%s %s", r.Method, r.URL.Path)

	//   /basepath/groupname/key  判断groupname和key是否存在
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	//获取group
	group := GetGroup(groupName)

	//groupname不在groups里
	if group == nil {
		http.Error(w, "no such group:"+groupName, http.StatusNotFound)
		return
	}


	//从group中获取get
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//返回response
	w.Header().Set("Context-type", "application/octet-stream")
	w.Write(view.ByteSlice())

}
