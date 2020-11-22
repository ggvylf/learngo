package geecache

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/ggvylf/learngo/geecache/geecache/consistenthash"
	pb "github.com/ggvylf/learngo/geecache/geecache/geecachepb"

	"github.com/golang/protobuf/proto"
)

//确保指定的类型实现了接口，如果会没实现会报错
var _PeerGetter = (*httpGetter)(nil)
var _PeerPicker = (*HTTPPool)(nil)

//制定默认值
const (
	defaultBasePath = "/_geecache/"
	defualtReplicas = 50
)

type HTTPPool struct {
	//自己的地址，主机名或ip+端口
	self string

	//节点通讯地址的前缀
	basePath string

	//互斥锁
	mu sync.Mutex

	//一致性hash算法的Map，用来根据具体的key选择节点
	peers *consistenthash.Map

	//远程节点和对应的httpGetter的映射关系，每个远程点对应一个httpGetter
	httpGetters map[string]*httpGetter
}

//给每个节点实例化一个HTTPPool
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

//实现ServerHTTP方法
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//前缀basepath是否匹配，不匹配直接panic
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

	//从groups中获取group
	group := GetGroup(groupName)

	//groupname不在groups里
	if group == nil {
		http.Error(w, "no such group:"+groupName, http.StatusNotFound)
		return
	}

	//从group中获取key的value
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//使用proto来处理response的body
	body, err := proto.Marshal(&pb.Response{Value: view.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//返回response
	w.Header().Set("Context-type", "application/octet-stream")
	w.Write(body)
}

//http客户端
type httpGetter struct {
	//访问的远程节点地址
	baseURL string
}

//从远程的url地址获取返回值，并转换为[]bytes
func (h *httpGetter) Get(in *pb.Request, out *pb.Response) error {
	//url=baseURL+grupname+/+key
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		//安全替换group和key
		url.QueryEscape(in.GetGroup()),
		url.QueryEscape(in.GetKey()),
	)

	//请求远程地址
	res, err := http.Get(u)
	if err != nil {
		return err

	}

	defer res.Body.Close()

	//状态码不是200
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v\n", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)

	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}

	return nil
}

//实现PeerPicker接口，给每个节点创建一个httpGetter
func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	//实例化一个一致性hash算法
	p.peers = consistenthash.New(defualtReplicas, nil)
	//添加传入的节点
	p.peers.Add(peers...)

	p.httpGetters = make(map[string]*httpGetter, len(peers))

	//给每个节点创建一个http客户端httpGetter
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

//PickPeer方法包装了Get方法，根据具体的key，选择节点，并返回节点对应的http客户端
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}
