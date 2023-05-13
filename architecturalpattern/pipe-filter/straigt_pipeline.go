package main

// pipe结构体 包含pipeline的名字和对应的Filter的数组
type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

// 实例化pipline
func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

func (sp *StraightPipeline) Porcess(data Request) (Response, error) {
	var ret interface{}
	var err error

	// 依次调用filter实例处理
	for _, filter := range *sp.Filters {
		ret, err = filter.Porcess(data)
		if err != nil {
			return ret, err
		}
		// 上一个FIlter处理结果作为下一个Filter的data
		data = ret
	}

	// 执行完全部filter，返回结果
	return ret, err
}
