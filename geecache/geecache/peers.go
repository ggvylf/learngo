package geecache

import pb "github.com/ggvylf/learngo/geecache/geecache/geecachepb"

//根据key选择对应的PeerGetter
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

//从对应的group中查找缓存值
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
