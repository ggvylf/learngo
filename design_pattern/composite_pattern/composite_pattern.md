组合模式
# 引用
https://mp.weixin.qq.com/s/z5WBdkf4fM__oDan5gC3oA


## 什么是组合模式

组合模式（Composite Pattern）又叫作部分-整体（Part-Whole）模式，它的宗旨是通过将单个对象（叶子节点）和组合对象（树枝节点）用相同的接口进行表示，使得客户对单个对象和组合对象的使用具有一致性，属于结构型设计模式。

## 应用场景

组合模式的使用要求业务场景中的实体必须能够表示成树形结构才行，由组合模式将一组对象组织成树形结构，客户端（代码的使用者）可以将单个对象和组合对象都看做树中的节点，以统一处理逻辑，并且利用树形结构的特点，将对树、子树的处理转化成叶节点的递归处理，依次简化代码实现。

通过上边的描述我们可以马上想到文件系统、公司组织架构这些有层级结构的事物的操作会更适合应用组合模式。

## 组合模式的结构

![图片](data:image/svg+xml,%3C%3Fxml version='1.0' encoding='UTF-8'%3F%3E%3Csvg width='1px' height='1px' viewBox='0 0 1 1' version='1.1' xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink'%3E%3Ctitle%3E%3C/title%3E%3Cg stroke='none' stroke-width='1' fill='none' fill-rule='evenodd' fill-opacity='0'%3E%3Cg transform='translate(-249.000000, -126.000000)' fill='%23FFFFFF'%3E%3Crect x='249' y='126' width='1' height='1'%3E%3C/rect%3E%3C/g%3E%3C/g%3E%3C/svg%3E)

组合模式由以下几个角色构成：

1. 组件 （Component）：组件是一个接口，描述了树中单个对象和组合对象都要实现的的操作。
    
2. 叶节点 （Leaf） ：即单个对象节点，是树的基本结构， 它不包含子节点，因此也就无法将工作指派给下去，叶节点最终会完成大部分的实际工作。
    
3. 组合对象 （Composite）”——是包含叶节点或其他组合对象等子项目的符合对象。组合对象不知道其子项目所属的具体类， 它只通过通用的组件接口与其子项目交互。
    
4. 客户端 （Client）：通过组件接口与所有项目交互。因此， 客户端能以相同方式与树状结构中的简单或复杂对象进行交互。
    

## 组合模式代码实现

下面用一个公司组织架构的例子来演示下用代码怎么实现组合模式。

我们都知道大公司的组织架构会很复杂，往往是由集团总公司-->分公司，每个层级的公司还有不同的部门，比如说总公司有财务部，分公司也会有。分公司偏传统一点，在互联网大厂有可能会按BG、BU这样分，不过在展示层级结构上意思都一样。

咱们来看下这个例子，使用的是Go语言的代码来实现组合模式。首先我们定义一个组织的行为接口，这个接口大到总公司小到一个部门都得实现：

`// 表示组织机构的接口  
type Organization interface {  
    display()  
    duty()  
}  
`

这里为了简单演示，接口里就提供两个方法，一个是打印出自己的组织结构的方法`display()`另外一个是展示组织职责的方法`duty()`。接下来定义和实现组合对象的行为：

`// 组合对象--上级部门  
"本文使用的完整可运行源码  
去公众号「网管叨bi叨」发送【设计模式】即可领取"  
type CompositeOrganization struct {  
    orgName string  
    depth   int  
    list    []Organization  
}  
  
func NewCompositeOrganization(name string, depth int) *CompositeOrganization {  
    return &CompositeOrganization{name, depth, []Organization{}}  
}  
  
func (c *CompositeOrganization) add(org Organization) {  
    if c == nil {  
        return  
    }  
    c.list = append(c.list, org)  
}  
  
func (c *CompositeOrganization) remove(org Organization) {  
    if c == nil {  
        return  
    }  
    for i, val := range c.list {  
        if val == org {  
            c.list = append(c.list[:i], c.list[i+1:]...)  
            return  
        }  
    }  
    return  
}  
  
func (c *CompositeOrganization) display() {  
    if c == nil {  
        return  
    }  
    fmt.Println(strings.Repeat("-", c.depth * 2), " ", c.orgName)  
    for _, val := range c.list {  
        val.display()  
    }  
}  
  
func (c *CompositeOrganization) duty() {  
    if c == nil {  
        return  
    }  
  
    for _, val := range c.list {  
        val.duty()  
    }  
}  
`

组合对象用来表示有下属部门的组织，在代码里可以看到，它持有一个`[]Organization`类型的列表，这里存放的是它的下属组织。组合对象的`display`、`duty`这两个方法的实现完全就是把工作委托给他们的下属组织来做的，这也是组合模式的特点。

下面我们再来看两个职能部门人力资源和财务部门的类型实现。

`// Leaf对象--人力资源部门  
"本文使用的完整可运行源码  
去公众号「网管叨bi叨」发送【设计模式】即可领取"  
type HRDOrg struct {  
    orgName string  
    depth   int  
}  
  
func (o *HRDOrg) display() {  
    if o == nil {  
        return  
    }  
    fmt.Println(strings.Repeat("-", o.depth * 2), " ", o.orgName)  
}  
  
func (o *HRDOrg) duty() {  
    if o == nil {  
        return  
    }  
    fmt.Println(o.orgName, "员工招聘培训管理")  
}  
  
// Leaf对象--财务部门  
type FinanceOrg struct {  
    orgName string  
    depth   int  
}  
  
func (f *FinanceOrg) display() {  
    if f == nil {  
        return  
    }  
    fmt.Println(strings.Repeat("-", f.depth * 2), " ", f.orgName)  
}  
  
func (f *FinanceOrg) duty() {  
    if f == nil {  
        return  
    }  
    fmt.Println(f.orgName, "员工招聘培训管理")  
}  
`

只要我们在客户端中组合好组织架构的结构，不管有几层组织，客户端对整个组织的调用是不会改变的。

`func main() {  
    root := NewCompositeOrganization("北京总公司", 1)  
    root.add(&HRDOrg{orgName: "总公司人力资源部", depth: 2})  
    root.add(&FinanceOrg{orgName: "总公司财务部", depth: 2})  
  
    compSh := NewCompositeOrganization("上海分公司", 2)  
    compSh.add(&HRDOrg{orgName: "上海分公司人力资源部", depth: 3})  
    compSh.add(&FinanceOrg{orgName: "上海分公司财务部", depth: 3})  
    root.add(compSh)  
  
    compGd := NewCompositeOrganization("广东分公司", 2)  
    compGd.add(&HRDOrg{orgName: "广东分公司人力资源部", depth: 3})  
    compGd.add(&FinanceOrg{orgName: "南京办事处财务部", depth: 3})  
    root.add(compGd)  
  
    fmt.Println("公司组织架构：")  
    root.display()  
  
    fmt.Println("各组织的职责：")  
    root.duty()  
}  
`

> 本文的完整源码，已经同步收录到我整理的电子教程里啦，可向我的公众号「网管叨bi叨」发送关键字【设计模式】领取。

![图片](https://mmbiz.qpic.cn/mmbiz_png/z4pQ0O5h0f4sIbSOaQzc4h3daCJAazUUGOUv57J43skHF8Gq2kf8SMHdpZHY7IeYeiamNOwibjTggchWEbMX57pw/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

公众号「网管叨bi叨」发送关键字【设计模式】领取。

组合模式和上一节我们学的[装饰器模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497282&idx=1&sn=0de76856e8649967bd3979cb122383fd&scene=21#wechat_redirect)在结构上挺像的，下面我们来说说他们的区别。

## 组合和装饰器的区别

组合模式和装饰器模式在结构上很像，拥有非常相似的类结构（相似到组合模式的类图就是我Copy装饰器模式改了下方法名字......）。但是两者在使用意图上是有区别的。

**组合模式**：为叶子对象和组合对象提供了统一的接口，叶子对象分担组合对象要做的工作。其实组合对象就是派了下活儿，等下面的干完后，它再给上层调用者返（汇）回（报），类似于公司里的那些组合\*。

**装饰器模式**：装饰器属于大哥带小弟的类型，核心的活儿是小弟干的（小弟就是被装饰的对象）但是各位大哥会帮你做好干活儿之外的事儿，比如公司你在公司里的Mentor、项目经理、领导们干的事儿就是给在给你做增强，你可以把他们理解成是你的装饰器😉。

说点题外话，如果你的Mentor、领导没有给你做增强，那当初他们给你定级P7是高于你面试的水平的。是希望进来后你能够拼一把，快速成长起来。P7这个层级，不是把事情做好就可以的。你需要有体系化思考的能力，它的价值点在哪里，你是否做出了壁垒形成了核心竞争力，是否沉淀了一套可复用的物理资料和方法论？...... （字儿太多了，完整版请自行搜索）‍‍‍‍‍‍‍‍‍‍‍‍‍‍‍‍‍‍‍‍‍‍

## 总结

组合模式的优点主要有以下两点

1. 实现类似树形结构，可以清楚地定义各层次的复杂对象，表示对象的全部或部分层次。
    
2. 简化了客户端代码，让客户端忽略了层次的差异，方便对整个层次结构进行控制。
    

实际上，组合模式与其说是一种设计模式，倒不如说是对业务场景的一种数据结构和算法的抽象，场景中的数据可以表示成树这种结构，业务需求的逻辑可以通过对树的递归遍历算法实现。

好了，如果喜欢今天的文章还请多多转发和关注，设计模式这个系列春节前咱们就不再更新啦，大家好好过年。

过年我应该就回家里宅着，打打游戏、拼拼乐高。看天气预报过年要下雪… 到时候有机会发点照片上来，让你们感受下什么叫做"我在南方的艳阳里大雪纷飞"