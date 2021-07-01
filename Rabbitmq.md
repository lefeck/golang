# docker 安装 rabbitmq

> rabbitMQ是一款开源的高性能消息中间件，最近项目要使用，于是使用docker搭建，快速方便

## 一、获取和拉取镜像

```
# docker search rabbitMq
# docker pull docker.io/rabbitmq:3.8-management
```

我们选择了STARS数最多的官方镜像，此处需要注意，默认rabbitmq镜像是不带web端管理插件的，所以指定了镜像tag为3.8-management，表示下载包含web管理插件版本镜像，其它Tag版本可以访问[DockerHub](https://hub.docker.com/_/rabbitmq)查询

## 二、创建rabbitMq容器

1.使用`docker images`获取查看rabbitMQ镜像ID，我的是`4b23cfb64730`

2.执行`docker run --name rabbitmq -d -p 15672:15672 -p 5672:5672 4b23cfb64730`命令创建rabbitMq容器，关于其中的参数含义如下：

- --name指定了容器名称
- -d 指定容器以后台守护进程方式运行
- -p指定容器内部端口号与宿主机之间的映射，rabbitMq默认要使用15672为其web端界面访问时端口，5672为数据通信端口

命令执行完毕后，docker会使用ID为 `4b23cfb64730`的镜像创建容器，创建完成后返回容器ID为`3ae75edc48e2416292db6bcae7b1054091cb....(太长省略)`

执行`docker ps`可以查看正在运行的容器，我们能看到rabbitMq已经运行



3.查看容器日志 使用`docker logs -f 容器ID`命令可以查看容器日志，我们执行`docker logs -f 3ae`命令查看rabbitMq在启动过程中日志，3ae是容器ID的简写——容器ID太长，使用时其写前几位即可

从日志可以看出，rabbitMq默认创建了guest用户，并且赋予administrator角色权限，同时服务监听5672端口TCP连接和15672端口的HTTP连接，至此说明安装成功。



## 三、访问rabbitMq

#### 1.访问web界面

在浏览器 输入你的`主机Ip:15672`回车即可访问rabbitMq的Web端管理界面，默认用户名和密码都是`guest`。



#### 2.新添加一个账户

默认的`guest` 账户有访问限制，默认只能通过本地网络(如 localhost) 访问，远程网络访问受限，所以在使用时我们一般另外添加用户，例如我们添加一个root用户：

①执行`docker exec -i -t 3ae bin/bash`进入到rabbitMq容器内部

```
[root@localhost docker]# docker exec -i -t 3a bin/bash
root@3ae75edc48e2:/# 
```

②执行`rabbitmqctl add_user root 123456` 添加用户，用户名为root,密码为123456

```
root@3ae75edc48e2:/# rabbitmqctl add_user root 123456 
Adding user "root" ...
```

③执行`abbitmqctl set_permissions -p / root ".*" ".*" ".*"` 赋予root用户所有权限

```
root@3ae75edc48e2:/# rabbitmqctl set_permissions -p / root ".*" ".*" ".*"
Setting permissions for user "root" in vhost "/" ...
```

④执行`rabbitmqctl set_user_tags root administrator`赋予root用户administrator角色

```
root@3ae75edc48e2:/# rabbitmqctl set_user_tags root administrator
Setting tags for user "root" to [adminstrator] ...
```

⑤执行`rabbitmqctl list_users`查看所有用户即可看到root用户已经添加成功

```
root@3ae75edc48e2:/# rabbitmqctl list_users
Listing users ...
user	tags
guest	[administrator]
root	[administrator]
```

执行`exit`命令，从容器内部退出即可。这时我们使用root账户登录web界面也是可以的。到此，rabbitMq的安装就结束了，接下里就实际代码开发。





# RabbitMQ和AMQP详解

https://cloud.tencent.com/developer/article/1513408



[消息队列](https://cloud.tencent.com/product/cmq?from=10680)（Message Queue）提供一个异步通信机制，消息的发送者不必苦苦等待着消息被处理完成，转而继续自己的工作。消息中间件负责处理网络通信，如果网络连接不可用，消息被暂存于队列当中，当网络畅通的时候再用。消息队列在企业中应用很广泛，可选择的有ActiveMQ、RabbitMQ，Kafka，阿里巴巴自主开发RocketMQ等。本文讨论 RabbitMQ 。

## 1.介绍

欲了解 RabbitMQ 先要了解 MQ。 RabbitMQ 是 MQ 的一种实现。

### 1.1 MQ 介绍

>  MQ（Message Queue）消息队列，是基础数据结构中“先进先出”的一种数据结构。一般用来解决应用解耦，异步消息，流量削锋等问题，实现高性能，高可用，可伸缩和最终一致性架构。 

它由这些组成：

- 生产者：生产者产生消息，并把消息放入队列。
- 队列：把要传输的数据（消息）放在队列中，用队列机制来实现消息传递。生产者产生消息并把消息放入队列，然后由消费者去处理。
- 消费者：消费者可以到指定队列拉取消息，或者订阅相应的队列，由MQ服务端给其推送消息。

作用：

- 解耦：对应用/模块进行分离，引入了消息代理作为中间消息服务
- 异步：主业务执行中将消息放入MQ后不等待，从业务异步执行返回结果。
- 削峰：高并发情况下，业务异步处理，提供高峰期业务处理能力，避免系统瘫痪

RabbitMQ 是 MQ 的一种实现，下面介绍下 RabBMQ。

### 1.2 RabbitMQ 介绍

>  RabBMQ是一个广泛部署的开源消息代理。 

![gcp932876c](/Users/jinhuaiwang/Desktop/rabbitmq/picture/gcp932876c.png)

RabbitMQ 流式管道

特点：

- 异步消息传递
- 易于部署
- 支持集群，用于高可用性和吞吐量。支持分布式部署。
- 开发者友好，支持各种流行的开发语言。比如Java,Ruby,GO。
- 方便的 管理与监控 工具

## 2. AMQP（高级消息队列协议）概述

​		RabbitMQ 是一个实现了 AMQP协议 的工具软件，所以 AMQP 中的概念和准则也适用于 RabbitMQ。下面重点介绍AMQP，它能帮助我们深刻的理解。

>  AMQP（高级消息队列协议）是一个网络协议。它支持符合要求的客户端应用 和消息中间件代理之间进行通信。 

### 为什么会有 AMQP？

​		软件系统中存在不同厂商的不兼容产品的问题，异构系统的集成是非常昂贵和复杂的。早期的消息传递解决方案也非常昂贵，往往专门用于大公司负担得起。

​		AMQP 设计目标：

- 为消息中间件创建开放标准
- 启用各种技术和平台之间的互操作性

AMQP 不仅使这些不同的系统能够相互通信，而且能够实现不同的产品。

### 消息中间件（经纪人）及其角色

“消息代理” 收到 “消息生产者”并将它们 “路由” 到 “消费者”。

（发布它们的应用程序）  -->  消息代理 ---> （处理它们的应用程序）

## 2. AMQP 模型简介

### 2.1 工作过程

它工作过程如下图：

- 消息（Message）被发布者（Publisher）发送给交换机（Exchange）
- 交换机（Exchange）可以理解成邮局，交换机将收到的消息根据路由规则分发给绑定的队列（Queue）
- 最后，AMQP代理会将消息投递给订阅了此队列的消费者（Consumer），或者消费者按照需求自行获取。

![1ahinf06lj](/Users/jinhuaiwang/Desktop/rabbitmq/picture/1ahinf06lj.png)

一些其他情况： 

​		**消息属性：** 发布消息时可以给消息指定各种消息属性（message meta-data）。有些属性会被消息代理（brokers）使用，有些只能被接收消息的应用所使用。

​		**回执（acknowledgement）：** 网络是不可靠的，有可能在处理消息的时候失败。AMQP 包含了一个消息确认的概念：当一个消息成功到达消费者后（consumer），消费者会通知一下消息代理（broker），这个可以是自动的也可以由开发者执行。

​		当“消息确认”被启用的时候，消息代理不会完全将消息从队列中删除，直到它收到来自消费者的确认回执（acknowledgement）。

​		**无法到达** 当一个消息无法被成功路由时，消息或许会被返回给发布者并被丢弃。或者，如果消息代理执行了延期操作，消息会被放入一个所谓的死信队列中。此时，消息发布者可以选择某些参数来处理这些特殊情况。

### 2.2  AMQP 内部模型

![iw2pkb5all](/Users/jinhuaiwang/Desktop/rabbitmq/picture/iw2pkb5all.png)

## 3. 交换机（Exchange)

​		交换机是要掌握的重点，这一章节重点来讲。

>  交换机 用来传输消息的，交换机拿到一个消息之后将它路由给一个队列。 

​		它的传输策略是由交换机类型和被称作绑定（bindings）的规则所决定的。

四种交换机：

| Name（交换机类型）            | 默认名称                                |
| :---------------------------- | :-------------------------------------- |
| Direct exchange（直连交换机） | (空字符串) ， amq.direct                |
| Fanout exchange（扇型交换机） | amq.fanout                              |
| Topic exchange（主题交换机）  | amq.topic                               |
| Headers exchange（头交换机）  | amq.match (and amq.headers in RabbitMQ) |

​		**交换机状态** 交换机可以有两个状态：持久（durable）、暂存（transient）。持久化的交换机会在消息代理（broker）重启后依旧存在，而暂存的交换机则不会。并不是所有的应用场景都需要持久化的交换机。



下面分别说明四种交换机类型

### 3.1 直连型交换机（ Direct Exchange）

​		消息可以携带一个属性 “路由键（routing key）”，以辅助标识被路由的方式。直连型交换机（direct exchange）根据消息携带的路由键将消息投递给对应队列的。

它如何工作：

- 将一个队列绑定到某个交换机上，同时赋予该绑定（Binding）一个路由键（routing key）
- 当一个携带着路由键为 “key1” 的消息被发送给直连交换机时，交换机会把它路由给 “Binding名称等于 key1”  的队列。

![nfyrjify24](/Users/jinhuaiwang/Desktop/rabbitmq/picture/nfyrjify24.png)

​																		直连型交换机图例

**总结：** Binding 的 Routing Key 要和 消息的 Routing Key 完全匹配

### 3.2 扇型交换机 （ Fanout Exchange）

扇型交换机将消息路由给绑定到它身上的所有队列，而不理会绑定的路由键。

如果N个队列绑定到某个扇型交换机上，当有消息发送给此扇型交换机时，交换机会将消息的拷贝分别发送给这所有的N个队列。扇型用来交换机处理消息的广播路由（broadcast routing）。

案例：

- MMO游戏可以使用它来处理排行榜更新等全局事件
- 体育新闻网站可以用它来实时地将比分更新分发给多端
- 在群聊的时候，它被用来分发消息给参与群聊的用户。

![af8ztbbrma](/Users/jinhuaiwang/Desktop/rabbitmq/picture/af8ztbbrma.png)

​																	扇型交换机图例

**总结** 不管 消息的Routing Key，广播给这个交换机下的所有绑定队列。

### 3.3 主题交换机（ Topic Exchanges）

主题交换机通过对消息的`路由键`和  “绑定的主题名称” 进行模式匹配，将消息路由给匹配成功的队列。

它的工作方式：

- 为绑定的 Routing Key 指定一个 “主题”。模式匹配用用 *, # 等字符进行模糊匹配。比如 usa.# 表示 以  usa.开头的多个消息 到这里来。
- 交换机将按消息的 Routing Key 的值的不同路由到  匹配的主题队列。

主题交换机经常用来实现各种分发/订阅模式及其变种。主题交换机通常用来实现消息的多播路由（multicast routing）。

![ezoelqmind](/Users/jinhuaiwang/Desktop/rabbitmq/picture/ezoelqmind.png)

使用案例：

- 由多个人完成的后台任务，每个人负责处理某些特定的任务
- 股票价格更新涉及到分类或者标签的新闻更新（

**总结：** 绑定 的 Routing Key 和 消息的 Routing Key 进行字符串的模糊匹配。

### 3.4 头交换机 (Headers exchange)

头交换机使用多个消息属性来代替路由键建立路由规则。通过判断消息头的值能否与指定的绑定相匹配来确立路由规则。

在实际中并不常用。

## 4. 相关实体概念

### 4.1 队列（ Queue ）

队列 存储着即将被应用消费掉的消息。

**名称** 可以为队列指定一个名称。

**队列持久化**

- 持久化队列（Durable queues）会被存储在磁盘上，当消息代理（broker）重启的时候，它可以被重新恢复。
- 没有被持久化的队列称作暂存队列（Transient queues）

### 4.2 绑定（Binding）

绑定是交换机（exchange）将消息（message）路由给队列（queue）所需遵循的规则。

如果要指示交换机“E”将消息路由给队列“Q”，那么“Q”就需要与“E”进行绑定。绑定操作需要定义一个可选的路由键（routing key）属性给某些类型的交换机。

路由键的意义在于从发送给交换机的众多消息中选择出某些消息，将其路由给绑定的队列。

### 4.3  消费者 （ Consumer ）

消费者即使用消息的客户。

**消费者标识** 每个消费者（订阅者）都有一个叫做消费者标签的标识符。它可以被用来退订消息。

一个队列可以注册多个消费者，也可以注册一个独享的消费者（当独享消费者存在时，其他消费者即被排除在外）。

### 4.4 消息确认 (acknowledgement)

什么时候删除消息才是正确的？有两种情况

- 自动确认模式：当消息代理（broker）将消息发送给应用后立即删除。
- 显式确认模式：待应用（application）发送一个确认回执（acknowledgement）后再删除消息。

在显式模式下，由消费者来选择什么时候发送确认回执（acknowledgement）。

- 应用可以在收到消息后立即发送
- 或将未处理的消息存储后发送
- 或等到消息被处理完毕后再发送确认回执。

如果一个消费者在尚未发送确认回执的情况下挂掉了，那代理会将消息重新投递给另一个消费者。如果当时没有可用的消费者了，消息代理会死等下一个注册到此队列的消费者，然后再次尝试投递。

**拒绝消息**

当一个消费者接收到某条消息后，处理过程有可能成功，有可能失败。

应用可以向消息代理表明，本条消息由于“拒绝消息（Rejecting Messages）”的原因处理失败了（或者未能在此时完成）。当拒绝某条消息时，应用可以告诉消息代理如何处理这条消息——销毁它或者重新放入队列。

### 4.5 消息 ( Message )

消息的组成：

- 消息属性
- 消息主体（有效载荷）

消息属性（Attributes）常见的有：

- Content type（内容类型）
- Content encoding（内容编码）
- Routing key（路由键）
- Delivery mode (persistent or not) 投递模式（持久化 或 非持久化）
- Message priority（消息优先权）
- Message publishing timestamp（消息发布的时间戳）
- Expiration period（消息有效期）
- Publisher application id（发布应用的ID）

消息体：

- 消息体即消息实际携带的数据，消息代理不会检查或者修改有效载荷。
- 消息可以只包含属性而不携带有效载荷。
- 它通常会使用类似JSON这种序列化的格式数据。
- 常常约定使用"content-type" 和 "content-encoding" 这两个字段分辨消息。

### 4.5 连接 (Connection)

AMQP 连接通常是长连接。AMQP是一个使用TCP提供可靠投递的应用层协议。AMQP使用认证机制并且提供TLS（SSL）保护。

当一个应用不再需要连接到AMQP代理的时候，需要优雅的释放掉AMQP连接，而不是直接将TCP连接关闭。

### 4.6 通道 （channels）

AMQP 提供了通道（channels）来处理多连接，可以把通道理解成共享一个TCP连接的多个轻量化连接。

这可以应对有些应用需要建立多个连接的情形，开启多个TCP连接会消耗掉过多的系统资源。

在多线程/进程的应用中，为每个线程/进程开启一个通道（channel）是很常见的，并且这些通道不能被线程/进程共享。

**通道号** 通道之间是完全隔离的，因此每个AMQP方法都需要携带一个通道号，这样客户端就可以指定此方法是为哪个通道准备的。

### 4.7 虚拟主机 (vhost)

为了在一个单独的代理上实现多个隔离的环境（用户、用户组、交换机、队列 等），AMQP提供了一个虚拟主机（virtual hosts - vhosts）的概念。

这跟Web servers虚拟主机概念非常相似，这为AMQP实体提供了完全隔离的环境。当连接被建立的时候，AMQP客户端来指定使用哪个虚拟主机。



# Golang连接rabbitmq



## 1.简单模式

RabbitMQ是一个消息中间件，你可以想象它是一个邮局。当你把信件放到邮箱里时，能够确信邮递员会正确地递送你的信件。RabbitMq就是一个邮箱、一个邮局和一个邮递员。

- 发送消息的程序是生产者
- 队列就代表一个邮箱。虽然消息会流经RbbitMQ和你的应用程序，但消息只能被存储在队列里。队列存储空间只受服务器内存和磁盘限制，它本质上是一个大的消息缓冲区。多个生产者可以向同一个队列发送消息，多个消费者也可以从同一个队列接收消息.
- 消费者等待从队列接收消息

![简单模式](https://segmentfault.com/img/remote/1460000025126512)

### 生产者发送消息

```go
producer_task.go

package main

import (
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "math/rand"
    "os"
    "strings"
    "time"
)

const (
    //AMQP URI
    uri = "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange name
    exchangeName = ""
    //Durable AMQP queue name
    queueName = "test-idoall-queues-task"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main() {
    bodyMsg := bodyFrom(os.Args)
    //调用发布消息函数
    publish(uri, exchangeName, queueName, bodyMsg)
    log.Printf("published %dB OK", len(bodyMsg))
}

func bodyFrom(args []string) string {
    var s string
    if (len(args) < 2) || os.Args[1] == "" {
        s = "hello idoall.org"
    } else {
        s = strings.Join(args[1:], " ")
    }
    return s
}

//发布者的方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
//@body, 主体内容
func publish(amqpURI string, exchange string, queue string, body string) {
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()

    log.Printf("got queue, declaring %q", queue)

    //创建一个queue
    q, err := channel.QueueDeclare(
        queueName, // name
        false,     // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    failOnError(err, "Failed to declare a queue")

    log.Printf("declared queue, publishing %dB body (%q)", len(body), body)

    // Producer只能发送到exchange，它是不能直接发送到queue的。
    // 现在我们使用默认的exchange（名字是空字符）。这个默认的exchange允许我们发送给指定的queue。
    // routing_key就是指定的queue名字。
    tick := time.NewTicker(time.Millisecond * time.Duration(rand.Intn(1000)))
    for {
      //写入通道
        <-tick.C
      // 构建一个生产者，将消息放入队列
        err = channel.Publish(
            exchange, // exchange
            q.Name,   // routing key
            false,    // mandatory
            false,    // immediate
          	// 构建一个消息
            amqp.Publishing{
                Headers:         amqp.Table{},
                ContentType:     "text/plain",
                ContentEncoding: "",
                Body:            []byte(body), 
            })
    }
    failOnError(err, "Failed to publish a message")
}
```



### 消费者接收队列

```go
consumer_task.go

package main

import (
    "bytes"
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "time"
)

const (
    //AMQP URI
    uri = "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange nam
    exchangeName = ""
    //Durable AMQP queue name
    queueName = "test-idoall-queues-task"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main() {
    //调用消息接收者
    consumer(uri, exchangeName, queueName)
}

//接收者方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
func consumer(amqpURI string, exchange string, queue string) {
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel,可以理解为多路复用的一个tcp长连接
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()

    log.Printf("got queue, declaring %q", queue)

    //声明一个queue
    q, err := channel.QueueDeclare(
        queueName, // name
        false,     // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    failOnError(err, "Failed to declare a queue")

    log.Printf("Queue bound to Exchange, starting Consume")
    //订阅消息
    msgs, err := channel.Consume(
        q.Name, // queue
        "",     // consumer
        false,  // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    //创建一个channel
    forever := make(chan bool)

    //调用gorountine
    go func() {
      //不断的读取消息
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)
            //*
            dot_count := bytes.Count(d.Body, []byte("."))
            t := time.Duration(dot_count)
            time.Sleep(t * time.Second)
            //*/
            log.Printf("Done")
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

    //没有写入数据，一直等待读，阻塞当前线程，目的是让线程不退出
    <-forever
}
```



## 2.工作模式

![image](https://segmentfault.com/img/bVbRFmM)
		工作队列(即任务队列)背后的**主要思想是避免立即执行资源密集型任务**，并且必须等待它完成。相反，我们将任务安排在稍后完成。

​		我们将任务封装为消息并将其发送到队列。后台运行的工作进程将获取任务并最终执行任务。当运行多个消费者时，任务将在它们之间分发。

​		使用任务队列的一个优点是能够**轻松地并行工作**。如果我们正在积压工作任务，我们可以添加更多工作进程，这样就可以轻松扩展。



### 消息确认

​		一个消费者接收消息后,在消息没有完全处理完时就挂掉了,那么这时会发生什么呢?

​		就上面的代码来说,**rabbitmq把消息发送给消费者后,会立即删除消息,那么消费者挂掉后,它没来得及处理的消息就会丢失**,

​		**为了确保消息不会丢失，rabbitmq支持消息确认(回执)**。当一个消息被消费者接收到并且执行完成后，**消费者会发送一个ack** (acknowledgment) 给rabbitmq服务器, 告诉他我已经执行完成了，你可以把这条消息删除了。

​		如果一个**消费者没有返回消息确认就挂掉了**（信道关闭，连接关闭或者TCP链接丢失），rabbitmq就会明白，这个消息没有被处理完成，**rebbitmq就会把这条消息重新放入队列**，如果在这时**有其他的消费者在线，那么rabbitmq就会迅速的把这条消息传递给其他的消费者**，这样就确保了没有消息会丢失。

​		这里不存在消息超时, **rabbitmq只在消费者挂掉时重新分派消息**, 即使消费者花非常久的时间来处理消息也可以

​		**手动消息确认默认是开启的**，前面的例子我们通过autoAck=ture把它关闭了。我们现在要把它设置为**autoAck=false，然后工作进程处理完意向任务时,发送一个消息确认(回执)**。接下来我们就来实现一下:

#### 生产者发送消息

```go
producer_acknowledgments.go

package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "github.com/streadway/amqp"
)

/**
 * use
 * go run producer_acknowledgments.go First message. && go run producer_acknowledgments.go Second message.. && go run producer_acknowledgments.go Third message... && go run producer_acknowledgments.go Fourth message.... && go run producer_acknowledgments.go Fifth message.....
 */
const (
    //AMQP URI
    uri          =  "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange name
    exchangeName =  ""
    //Durable AMQP queue name
    queueName    = "test-idoall-queues-acknowledgments"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main(){
    bodyMsg := bodyFrom(os.Args)
    //调用发布消息函数
    publish(uri, exchangeName, queueName, bodyMsg)
    log.Printf("published %dB OK", len(bodyMsg))
}

func bodyFrom(args []string) string {
    var s string
    if (len(args) < 2) || os.Args[1] == "" {
        s = "hello idoall.org"
    } else {
        s = strings.Join(args[1:], " ")
    }
    return s
}

//发布者的方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
//@body, 主体内容
func publish(amqpURI string, exchange string, queue string, body string){
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()

    log.Printf("got queue, declaring %q", queue)

    //创建一个queue
    q, err := channel.QueueDeclare(
        queueName, // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    log.Printf("declared queue, publishing %dB body (%q)", len(body), body)

    // Producer只能发送到exchange，它是不能直接发送到queue的。
    // 现在我们使用默认的exchange（名字是空字符）。这个默认的exchange允许我们发送给指定的queue。
    // routing_key就是指定的queue名字。
    err = channel.Publish(
        exchange,     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            Headers:         amqp.Table{},
            ContentType: "text/plain",
            ContentEncoding: "",
            Body:        []byte(body),
        })
    failOnError(err, "Failed to publish a message")
}
```

#### 消费者接收消息

```go
consumer_acknowledgments.go

package main

import (
    "fmt"
    "log"
    "bytes"
    "time"
    "github.com/streadway/amqp"
)

const (
    //AMQP URI
    uri           =  "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange nam
    exchangeName  = ""
    //Durable AMQP queue name
    queueName     = "test-idoall-queues-acknowledgments"
)
/**
 *
 */
//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main(){
    //调用消息接收者
    consumer(uri, exchangeName, queueName)
}

//接收者方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
func consumer(amqpURI string, exchange string, queue string){
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()

    log.Printf("got queue, declaring %q", queue)

    //创建一个queue
    q, err := channel.QueueDeclare(
        queueName, // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    log.Printf("Queue bound to Exchange, starting Consume")
    //订阅消息
    msgs, err := channel.Consume(
        q.Name, // queue
        "",     // consumer
        false,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    //创建一个channel
    forever := make(chan bool)

    //调用gorountine
    go func() {
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)
            dot_count := bytes.Count(d.Body, []byte("."))
            t := time.Duration(dot_count)
            time.Sleep(t * time.Second)
            log.Printf("Done")
            d.Ack(false)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

    //没有写入数据，一直等待读，阻塞当前线程，目的是让线程不退出
    <-forever
}
```



### 消息持久化

​		当rabbitmq关闭时, 我们队列中的消息仍然会丢失, 除非明确要求它不要丢失数据

​		要求rabbitmq不丢失数据要做如下两点: 把队列和消息都设置为可持久化(durable)

​		队列设置为可持久化, 可以在定义队列时指定参数durable为true

**由于我们之前已经定义过队列"hello world"是不可持久化的, 对已存在的队列, rabbitmq不允许对其定义不同的参数, 否则会出错,有两种方式进行修改:**

* 删除重建队列

* 另起一个名字 所以这里我们定义一个不同名字的队列"task_queue"

#### 生产者发送消息

```go
producer_durability.go
package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "github.com/streadway/amqp"
)


const (
    //AMQP URI
    uri          =  "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange name
    exchangeName =  ""
    //Durable AMQP queue name
    queueName    = "test-idoall-queues-durability"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main(){
    bodyMsg := bodyFrom(os.Args)
    //调用发布消息函数
    publish(uri, exchangeName, queueName, bodyMsg)
    log.Printf("published %dB OK", len(bodyMsg))
}

func bodyFrom(args []string) string {
    var s string
    if (len(args) < 2) || os.Args[1] == "" {
        s = "hello idoall.org"
    } else {
        s = strings.Join(args[1:], " ")
    }
    return s
}

//发布者的方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
//@body, 主体内容
func publish(amqpURI string, exchange string, queue string, body string){
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()

    log.Printf("got queue, declaring %q", queue)

    //创建一个queue
    q, err := channel.QueueDeclare(
        queueName, // name
        true,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    log.Printf("declared queue, publishing %dB body (%q)", len(body), body)

    // Producer只能发送到exchange，它是不能直接发送到queue的。
    // 现在我们使用默认的exchange（名字是空字符）。这个默认的exchange允许我们发送给指定的queue。
    // routing_key就是指定的queue名字。
    err = channel.Publish(
        exchange,     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            Headers:         amqp.Table{},
            //消息持久化
            DeliveryMode: amqp.Persistent,
            ContentType: "text/plain",
            ContentEncoding: "",
            Body:        []byte(body),
        })
    failOnError(err, "Failed to publish a message")
}
```

#### 消费者接收队列

```go
consumer_durability.go
package main

import (
    "fmt"
    "log"
    "bytes"
    "time"
    "github.com/streadway/amqp"
)

const (
    //AMQP URI
    uri           =  "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange nam
    exchangeName  = ""
    //Durable AMQP queue name
    queueName     = "test-idoall-queues-durability"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main(){
    //调用消息接收者
    consumer(uri, exchangeName, queueName)
}

//接收者方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
func consumer(amqpURI string, exchange string, queue string){
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()

    log.Printf("got queue, declaring %q", queue)

    //创建一个queue
    q, err := channel.QueueDeclare(
        queueName, // name
        true,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    log.Printf("Queue bound to Exchange, starting Consume")
    //订阅消息
    msgs, err := channel.Consume(
        q.Name, // queue
        "",     // consumer
        false,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    //创建一个channel
    forever := make(chan bool)

    //调用gorountine
    go func() {
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)
            dot_count := bytes.Count(d.Body, []byte("."))
            t := time.Duration(dot_count)
            time.Sleep(t * time.Second)
            log.Printf("Done")
            d.Ack(false)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

    //没有写入数据，一直等待读，阻塞当前线程，目的是让线程不退出
    <-forever
}
```



### 公平分发

​		rabbitmq会一次把多个消息分发给消费者, 这样可能造成有的消费者非常繁忙, 而其它消费者空闲. 而rabbitmq对此一无所知, 仍然会均匀的分发消息

​		我们可以使用` channel.Qos(1)` 方法, 这告诉rabbitmq一次只向消费者发送一条消息, 在返回确认回执前, 不要向消费者发送新消息. 而是把消息发给下一个空闲的消费者



#### 生产者发送消息

```go
producer_fair_dispatch.go
package main

import (
   "fmt"
   "log"
   "os"
   "strings"
   "github.com/streadway/amqp"
)

const (
   //AMQP URI
   uri          =  "amqp://guest:guest@localhost:5672/"
   //Durable AMQP exchange name
   exchangeName =  ""
   //Durable AMQP queue name
   queueName    = "test-idoall-queues-fair_dispatch"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
   if err != nil {
      log.Fatalf("%s: %s", msg, err)
      panic(fmt.Sprintf("%s: %s", msg, err))
   }
}

func main(){
   bodyMsg := bodyFrom(os.Args)
   //调用发布消息函数
   publish(uri, exchangeName, queueName, bodyMsg)
   log.Printf("published %dB OK", len(bodyMsg))
}

func bodyFrom(args []string) string {
   var s string
   if (len(args) < 2) || os.Args[1] == "" {
      s = "hello idoall.org"
   } else {
      s = strings.Join(args[1:], " ")
   }
   return s
}

//发布者的方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
//@body, 主体内容
func publish(amqpURI string, exchange string, queue string, body string){
   //建立连接
   log.Printf("dialing %q", amqpURI)
   connection, err := amqp.Dial(amqpURI)
   failOnError(err, "Failed to connect to RabbitMQ")
   defer connection.Close()

   //创建一个Channel
   log.Printf("got Connection, getting Channel")
   channel, err := connection.Channel()
   failOnError(err, "Failed to open a channel")
   defer channel.Close()

   log.Printf("got queue, declaring %q", queue)

   //创建一个queue
   q, err := channel.QueueDeclare(
      queueName, // name
      true,   // durable
      false,   // delete when unused
      false,   // exclusive
      false,   // no-wait
      nil,     // arguments
   )
   failOnError(err, "Failed to declare a queue")

   log.Printf("declared queue, publishing %dB body (%q)", len(body), body)

   // Producer只能发送到exchange，它是不能直接发送到queue的。
   // 现在我们使用默认的exchange（名字是空字符）。这个默认的exchange允许我们发送给指定的queue。
   // routing_key就是指定的queue名字。
   err = channel.Publish(
      exchange,     // exchange
      q.Name, // routing key
      false,  // mandatory
      false,  // immediate
      amqp.Publishing {
         Headers:         amqp.Table{},
         DeliveryMode: amqp.Persistent,
         ContentType: "text/plain",
         ContentEncoding: "",
         Body:        []byte(body),
      })
   failOnError(err, "Failed to publish a message")
}
 
```



#### 消费者接受消息

```go
consumer_fair_dispatch.go
package main

import (
    "fmt"
    "log"
    "bytes"
    "time"
    "github.com/streadway/amqp"
)

const (
    //AMQP URI
    uri           =  "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange nam
    exchangeName  = ""
    //Durable AMQP queue name
    queueName     = "test-idoall-queues-fair_dispatch"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main(){
    //调用消息接收者
    consumer(uri, exchangeName, queueName)
}

//接收者方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
func consumer(amqpURI string, exchange string, queue string){
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()

    log.Printf("got queue, declaring %q", queue)

    //创建一个queue
    q, err := channel.QueueDeclare(
        queueName, // name
        true,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    //每次只取一条消息
    err = channel.Qos(
        1,     // prefetch count
        0,     // prefetch size
        false, // global
    )
    failOnError(err, "Failed to set QoS")

    log.Printf("Queue bound to Exchange, starting Consume")
    //订阅消息
    msgs, err := channel.Consume(
        q.Name, // queue
        "",     // consumer
        false,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    //创建一个channel
    forever := make(chan bool)

    //调用gorountine
    go func() {
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)
            dot_count := bytes.Count(d.Body, []byte("."))
            t := time.Duration(dot_count)
            time.Sleep(t * time.Second)
            log.Printf("Done")
            //确认消息
            d.Ack(false)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

    //没有写入数据，一直等待读，阻塞当前线程，目的是让线程不退出
    <-forever
}
```



## 3.发布订阅模式

即向多个消费者传递同一条信息
![image](https://segmentfault.com/img/bVbRI0t)

### Exchanges 交换机

RabbitMQ消息传递模型的核心思想是，生产者永远不会将任何消息直接发送到队列。

相反，生产者只能向交换机(Exchange)发送消息。交换机是一个非常简单的东西。一边接收来自生产者的消息，另一边将消息推送到队列。交换器必须确切地知道如何处理它接收到的消息。它应该被添加到一个特定的队列中吗?它应该添加到多个队列中吗?或者它应该被丢弃。这些规则由exchange的类型定义。

有几种可用的交换类型:**direct、topic、header和fanout**。
创建exchange交换机logs: `c.exchangeDeclare("logs", "fanout");`
或`c.exchangeDeclare("logs", BuiltinExchangeType.FANOUT);`

fanout交换机非常简单。它只是将接收到的所有消息广播给它所知道的所有队列。



### 绑定 Bindings

![image](https://segmentfault.com/img/bVbRI10)

创建了一个exchange交换机和一个队列。现在我们需要告诉exchange向指定队列发送消息。exchange和队列之间的关系称为绑定。



### 生产者发送消息

```go
producer_exchange_logs.go
package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "github.com/streadway/amqp"
)


const (
    //AMQP URI
    uri          =  "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange name
    exchangeName =  "test-idoall-exchange-logs"
    //Exchange type - direct|fanout|topic|x-custom
    exchangeType = "fanout"
    //AMQP routing key
    routingKey   = ""
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main(){
    bodyMsg := bodyFrom(os.Args)
    //调用发布消息函数
    publish(uri, exchangeName, exchangeType, routingKey, bodyMsg)
    log.Printf("published %dB OK", len(bodyMsg))
}


func bodyFrom(args []string) string {
    var s string
    if (len(args) < 2) || os.Args[1] == "" {
        s = "hello idoall.org"
    } else {
        s = strings.Join(args[1:], " ")
    }
    return s
}

//发布者的方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@exchangeType, exchangeType的类型direct|fanout|topic
//@routingKey, routingKey的名称
//@body, 主体内容
func publish(amqpURI string, exchange string, exchangeType string, routingKey string, body string){
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()


    //创建一个queue
    log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
    err = channel.ExchangeDeclare(
        exchange,     // name
        exchangeType, //exchangeType的类型direct|fanout|topic
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // noWait
        nil,          // arguments
    )
    failOnError(err, "Failed to declare a queue")

    // 发布消息
    log.Printf("declared queue, publishing %dB body (%q)", len(body), body)
    err = channel.Publish(
        exchange,     // exchange
        routingKey, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            Headers:         amqp.Table{},
            ContentType: "text/plain",
            ContentEncoding: "",
            Body:        []byte(body),
        })
    failOnError(err, "Failed to publish a message")
}
```



### 消费者接收消息

```go
consumer_exchange_logs.go
package main

import (
    "fmt"
    "log"
    "github.com/streadway/amqp"
)

const (
    //AMQP URI
    uri           =  "amqp://guest:guest@localhost:5672/"
    //Durable AMQP exchange name
    exchangeName =  "test-idoall-exchange-logs"
    //Exchange type - direct|fanout|topic|x-custom
    exchangeType = "fanout"
    //AMQP binding key
    bindingKey   = ""
    //Durable AMQP queue name
    queueName     = ""
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main(){
    //调用消息接收者
    consumer(uri, exchangeName, exchangeType, queueName, bindingKey)
}

//接收者方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@exchangeType, exchangeType的类型direct|fanout|topic
//@queue, queue的名称
//@key , 绑定的key名称
func consumer(amqpURI string, exchange string, exchangeType string, queue string, key string){
    //建立连接
    log.Printf("dialing %q", amqpURI)
    connection, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer connection.Close()

    //创建一个Channel
    log.Printf("got Connection, getting Channel")
    channel, err := connection.Channel()
    failOnError(err, "Failed to open a channel")
    defer channel.Close()

    //创建一个exchange
    log.Printf("got Channel, declaring Exchange (%q)", exchange)
    err = channel.ExchangeDeclare(
        exchange,     // name of the exchange
        exchangeType, // type
        true,         // durable
        false,        // delete when complete
        false,        // internal
        false,        // noWait
        nil,          // arguments
    );
    failOnError(err, "Exchange Declare:")

    //创建一个queue
    q, err := channel.QueueDeclare(
        queueName, // name
        false,   // durable
        false,   // delete when unused
        true,   // exclusive 当Consumer关闭连接时，这个queue要被deleted
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    //绑定到exchange
    err = channel.QueueBind(
        q.Name, // name of the queue
        key,        // bindingKey
        exchange,   // sourceExchange
        false,      // noWait
        nil,        // arguments
    );
    failOnError(err, "Failed to bind a queue")

    log.Printf("Queue bound to Exchange, starting Consume")
    //订阅消息
    msgs, err := channel.Consume(
        q.Name, // queue
        "",     // consumer
        false,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    //创建一个channel
    forever := make(chan bool)

    //调用gorountine
    go func() {
        for d := range msgs {
            log.Printf(" [x] %s", d.Body)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

    //没有写入数据，一直等待读，阻塞当前线程，目的是让线程不退出
    <-forever
}
```



![img](https://img2018.cnblogs.com/blog/203395/201905/203395-20190504204227505-536500619.png)

 



## 4.路由模式

路由模式与订阅模式不同之处在于,我们将向其添加一个特性—我们将只订阅所有消息中的一部分.本文中已添加err/info/warning等报错提示来示范.
![image](https://segmentfault.com/img/bVbRI2I)

### 绑定

绑定是交换机和队列之间的关系。这可以简单地理解为:队列对来自此交换的消息感兴趣。

绑定可以使用额外的routingKey参数。为了避免与basic_publish参数混淆，我们将其称为bindingKey。这是我们如何创建一个键绑定:

```
ch.queueBind(queueName, EXCHANGE_NAME, "black");
```

bindingKey的含义取决于交换机类型。我们前面使用的fanout交换机完全忽略它。

### 直连交换机

上一节中的日志系统向所有消费者广播所有消息。我们希望扩展它，允许根据消息的严重性过滤消息。

前面我们使用的是fanout交换机，这并没有给我们太多的灵活性——它只能进行简单的广播。

我们将用直连交换机(Direct exchange)代替。它背后的路由算法很简单——消息传递到bindingKey与routingKey完全匹配的队列。

### 多重绑定 Multiple bindings

使用相同的bindingKey绑定多个队列是完全允许的。可以使用binding key "black"将X与Q1和Q2绑定。在这种情况下，直连交换机的行为类似于fanout，并将消息广播给所有匹配的队列。一条路由键为black的消息将同时发送到Q1和Q2。



### 生产者发送消息

```go
package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func failOnError(err error, msg string)  {
	if err !=nil {
		log.Fatalf("%s: %s", msg,err)
	}
}

const (
	uri = "amqp://guest:guest@192.168.10.168:5672"
	exchangeName = "logs_direct"
	exchangeType = "direct"
	routingKey = "info"
)

func bodyFrom(args []string)string  {
	var s string
	if (len(args)<2) || os.Args[1] == ""{
		s = "hello idoall.aorg"
	}else {
		s = strings.Join(args[1:]," ")
	}
	return s
}

func main()  {
	bodyMsg := bodyFrom(os.Args)
	SendMsg(uri,exchangeName,exchangeType,routingKey,bodyMsg)
}

func SendMsg(uri string, exchange string, exchangeType string, routingKey string, body string )  {
	conn, err := amqp.Dial(uri)
	failOnError(err,"failed to connect to rabbitmq")
	defer conn.Close()
	log.Printf("got queue, declaring %q", exchange)

	c, err := conn.Channel()
	failOnError(err,"failed to open a chanel")
	defer c.Close()

	log.Printf("got queue, declaring %q", routingKey)
	err = c.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err,"failed to declare a queue")
	log.Printf("declared queue, publishing %dB body(%q)",len(body))

	tick := time.NewTicker(time.Millisecond * time.Duration(rand.Intn(1000)))
	for  {
		//写数据到channel
		<-tick.C
		err = c.Publish(
			exchange,
			routingKey,
			false,
			false,
			amqp.Publishing{
				Headers: amqp.Table{},
				ContentType: "text/plain",
				ContentEncoding: "",
				Body: []byte(body),
			})
	}
	failOnError(err,"fail to publish a message")
}
```



### 消费者接收消息

```go
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)


const (
	//AMQP URI
	uri = "amqp://guest:guest@192.168.10.168:5672/"
	//Durable AMQP exchange nam
	//Durable AMQP exchange name
	exchangeName =  "logs_direct"
	//Exchange type - direct|fanout|topic|x-custom
	exchangeType = "direct"
	//AMQP binding key
	bindingKey   = "info"
	//Durable AMQP queue name
	queueName     = ""
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	//调用消息接收者
	consumer(uri, exchangeName, exchangeType,queueName,bindingKey)
}

//接收者方法
//
//@amqpURI, amqp的地址
//@exchange, exchange的名称
//@queue, queue的名称
func consumer(amqpURI string, exchangename string, exchangetype string,  queue string, key string) {
	//建立连接
	log.Printf("dialing %q", amqpURI)
	connection, err := amqp.Dial(amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	//创建一个Channel
	log.Printf("got Connection, getting Channel")
	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	log.Printf("got queue, declaring %q", exchangeType)

	//创建交换机exchange
	err = channel.ExchangeDeclare(
    exchangeName, //name:交换机名称
    exchangeType,  //exchange type: 路由模式下我们需要将类型设置为direct类型
    true,     // durable:进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
    false,     // autodelete: 是否为自动删除
    false,     // exclusive: true表示这个exchange不可以被客户端用来推送消息，仅仅是用来进行exchange和exchange之间的绑定
    false,     // no-wait:队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf("Queue bound to Exchange, starting Consume")
	//创建队列
	msgs, err := channel.QueueDeclare(
    queueName, // name: 队列名称
		false,     // durable: 进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
    false,  // autodelete: 是否为自动删除
    true,  // exclusive: 具有排他性
		false,  // no-local
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	//bind队列到exchange
	err = channel.QueueBind(
    msgs.Name, //name: 队列的名字,通过key去找绑定好的队列
    bindingKey,  //key: 在路由模式下，这里的key
    exchangeName, //exchange: 所绑定的交换器
		false,  //nowait
		nil,
		)
	failOnError(err, "Failed to bind a queue")
  //消费消息
	msg,err := channel.Consume(
    msgs.Name, 		//queuename: 队列名称
    "", //consumer: 用来区分多个消费者
    true, //autoack: 是否自动应答 意思就是收到一个消息已经被消费者消费完了是否主动告诉rabbitmq服务器我已经消费完了你可以去删除这个消息啦 默认是true
    false, //exclusive:是否具有排他性
    false, //nolocal:如果设置为true表示不能将同一个connection中发送的消息传递给同个connectio中的消费者
    false, //nowait: 队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		nil,
	)
  failOnError(err, "Failed to get consume")

	//创建通道
	forever := make(chan bool)
	//创建goroutine去接收消息
	go func() {
		for d := range msg {
			log.Printf("Received a message: [x] %s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<- forever //读取通道
}
```



## 5.topic模式

1. *号#号代表通配符
2. *号代表多个单词，#号代表一个单词
3. 路由功能添加模糊匹配
4. 消息产生者产生消息，把消息交给交换机
5. 交换机根据 key 的规则模糊匹配到对应的队列，由队列的监听消费者接收消息消费

一个消息被多个消费者获取。消息的目标 queue 可用 bindingkey以通配符（#：一个或者多个词，*: 一个词）的方式指定



[![Rabbitmq工作模式之topic模式](https://cdn.learnku.com/uploads/images/202006/22/58888/V9nPjbPZoG.png!large)](https://cdn.learnku.com/uploads/images/202006/22/58888/V9nPjbPZoG.png!large)



​		其实就是根据传入的 key 进行模糊匹配，匹配到 key 之后就去找交换机上跟这个 key 绑定的队列去读取进行消费，当然匹配到可能多个 key 那么可能就会有多个队列被消费！
直接上代码吧：

公共代码:

```go
rbtmqcs.go
package rbtmqcs


import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url 格式 amqp://账号:密码@rabbitmq服务器地址:端口号/Virtual Host
//格式在golang语言当中是固定不变的
const MQURL = "amqp://guest:guest@192.168.10.168:5672/"
//const MQURL = "amqp://huj:123456@192.168.1.219:5672/rbtmq"
type RabbitMQ struct {
	conn *amqp.Connection //需要引入amqp包 https://learnku.com/articles/44185教会你如何引用amqp包
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//链接信息
	Mqurl string
}

//创建RabbitMQ结构体实例
func NewRabbitMQ(queuename string,exchange string,key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName:queuename,Exchange:exchange,Key:key,Mqurl:MQURL}
	var err error
	//创建rabbitmq连接
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)  //通过amqp.Dial()方法去链接rabbitmq服务端
	rabbitmq.failOnErr(err,"创建连接错误!")  //调用我们自定义的failOnErr()方法去处理异常错误信息
	rabbitmq.channel,err = rabbitmq.conn.Channel() //链接上rabbitmq之后通过rabbitmq.conn.Channel()去设置channel信道
	rabbitmq.failOnErr(err,"获取channel失败!")
	return rabbitmq
}

//断开channel和connection
//为什么要断开channel和connection 因为如果不断开他会始终使用和占用我们的channel和connection 断开是为了避免资源浪费
func (r *RabbitMQ) Destory() {
	r.channel.Close()    //关闭信道资源
	r.conn.Close()       //关闭链接资源
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error,message string) {
	if err != nil {
		log.Fatalf("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}

//简单模式step：1.创建简单模式下的rabbitmq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//simple模式下交换机为空因为会默认使用rabbitmq默认的default交换机而不是真的没有 bindkey绑定建key也是为空的
	//特别注意：simple模式是最简单的rabbitmq的一种模式 他只需要传递queue队列名称过去即可  像exchange交换机会默认使用default交换机  绑定建key的会不必要传
	return NewRabbitMQ(queueName,"","")
}

//简单模式step:2.简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列,如果队列不存在，则会自动创建，如果队列存在则跳过创建直接使用  这样的好处保障队列存在，消息能发送到队列当中
	_,err := r.channel.QueueDeclare (
		r.QueueName,
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		false,
		//是否为自动删除  意思是最后一个消费者断开链接以后是否将消息从队列当中删除  默认设置为false不自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞 发送消息以后是否要等待消费者的响应 消费了下一个才进来 就跟golang里面的无缓冲channle一个道理 默认为非阻塞即可设置为false
		false,
		//其他的属性，没有则直接诶传入空即可 nil
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//2.发送消息到队列当中
	r.channel.Publish(
		//交换机 simple模式下默认为空 我们在上边已经赋值为空了  虽然为空 但其实也是在用的rabbitmq当中的default交换机运行
		r.Exchange,
		//队列的名称
		r.QueueName,
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		//要发送的消息
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(message),
		})
}

//简单模式step:3.简单模式下消费者代码
func (r *RabbitMQ) ConsumeSimple() {
	//申请队列和生产消息当中是一样一样滴 直接复制即可
	//1.申请队列,如果队列不存在，则会自动创建，如果队列存在则跳过创建直接使用  这样的好处保障队列存在，消息能发送到队列当中
	_,err := r.channel.QueueDeclare(
		//队列名称
		r.QueueName,
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		false,
		//是否为自动删除  意思是最后一个消费者断开链接以后是否将消息从队列当中删除  默认设置为false不自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞 发送消息以后是否要等待消费者的响应 消费了下一个才进来 就跟golang里面的无缓冲channle一个道理 默认为非阻塞即可设置为false
		false,
		//其他的属性，没有则直接诶传入空即可 nil
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//2.接收消息
	//建立了链接 就跟socket一样 一直在监听 从未被终止  这也就保证了下边的子协程当中程序的无线循环的成立
	msgs,err := r.channel.Consume(
		//队列名称
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答 意思就是收到一个消息已经被消费者消费完了是否主动告诉rabbitmq服务器我已经消费完了你可以去删除这个消息啦 默认是true
		true,
		//是否具有排他性
		false,
		//如果设置为true表示不能将同一个connection中发送的消息传递给同个connectio中的消费者
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,)
	if err != nil {
		fmt.Println(err)
	}

	//3.消费消息
	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		//子协程不会结束  因为msgs有监听 会不断的有值进来 就算没值也在监听  就跟socket服务一样一直在监听从未被中断！
		for d:=range msgs {
			//实现我们要处理的逻辑函数
			log.Printf("Recieved a message : %s",d.Body)
			//fmt.Println(d.Body)
		}
	}()
	log.Printf("[*] Waiting for message,To exit press CTRL+C")
	//最后我们来阻塞一下 这样主程序就不会死掉
	<-forever
}

//订阅模式step1:创建rabbitmq实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	//创建RabbitMQ实例
	return NewRabbitMQ("",exchangeName,"")
}

//订阅模式step2:生产者
func (r *RabbitMQ) PublishPub(message string) {
	//1.尝试创建交换机exchange 如果交换机存在就不用管他，如果不存在则会创建交换机
	err := r.channel.ExchangeDeclare(
		//交换机名称
		r.Exchange,
		//广播类型  订阅模式下我们需要将类型设置为广播类型
		"fanout",
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		true,
		//是否为自动删除  这里解释的会更加清楚：https://blog.csdn.net/weixin_30646315/article/details/96224842?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase
		false,
		//true表示这个exchange不可以被客户端用来推送消息，仅仅是用来进行exchange和exchange之间的绑定
		false,
		//直接false即可 也不知道干啥滴
		false,
		nil,
	)

	r.failOnErr(err,"Failed to declare an excha " + "nge")

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			Body:            []byte(message),//发送的内容一定要转换成字节的形式
		})
}

//订阅模式step3:消费者
func (r *RabbitMQ) RecieveSub() {
	//这一步和 订阅模式step2:生产者 里面的代码是一样一样滴
	//1.尝试创建交换机exchange 如果交换机存在就不用管他，如果不存在则会创建交换机
	err := r.channel.ExchangeDeclare(
		//交换机名称
		r.Exchange,
		//广播类型  订阅模式下我们需要将类型设置为广播类型
		"fanout",
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		true,
		//是否为自动删除  这里解释的会更加清楚：https://blog.csdn.net/weixin_30646315/article/details/96224842?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase
		false,
		//true表示这个exchange不可以被客户端用来推送消息，仅仅是用来进行exchange和exchange之间的绑定
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	r.failOnErr(err,"Failed to declare an excha " + "nge")
	//2.试探性创建队列，这里注意队列名称不要写
	q,err := r.channel.QueueDeclare(
		//随机生产队列名称 这个地方一定要留空
		"",
		false,
		false,
		//具有排他性   排他性的理解 这篇文章还是比较好的：https://www.jianshu.com/p/94d6d5d98c3d
		true,
		false,
		nil,
	)
	//3.绑定队列到exchange中去
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.Exchange,
		false,
		nil,
	)
	//4.消费代码
	//4.1接收队列消息
	message,err := r.channel.Consume(
		//队列名称
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答 意思就是收到一个消息已经被消费者消费完了是否主动告诉rabbitmq服务器我已经消费完了你可以去删除这个消息啦 默认是true
		true,
		//是否具有排他性
		false,
		//如果设置为true表示不能将同一个connection中发送的消息传递给同个connectio中的消费者
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//4.2真正开始消费消息
	forever := make(chan bool)
	go func() {
		for d:=range message {
			log.Printf("Received a message: %s",d.Body)
		}
	}()
	fmt.Println("退出请按 ctrl+c")
	<-forever
}

//路由模式step1:创建RabbitMQ实例
func NewRabbitMQRouting(exchangeName string,routingKey string) *RabbitMQ {
	return NewRabbitMQ("",exchangeName,routingKey)
}

//路由模式step2:发送消息
func (r *RabbitMQ) PublishRouting(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		//交换机名称
		r.Exchange,
		//类型  路由模式下我们需要将类型设置为direct这个和在订阅模式下是不一样的
		"direct",
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		true,
		//是否为自动删除  这里解释的会更加清楚：https://blog.csdn.net/weixin_30646315/article/details/96224842?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase
		false,
		//true表示这个exchange不可以被客户端用来推送消息，仅仅是用来进行exchange和exchange之间的绑定
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	r.failOnErr(err,"Failed to declare an excha " + "nge")
	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		//除了设置交换机这也要设置绑定的key值
		r.Key,
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			Body:            []byte(message),//发送的内容一定要转换成字节的形式
		})
}

//路由模式step3：消费者
func (r *RabbitMQ) RecieveRouting() {
	//1.尝试创建交换机exchange 如果交换机存在就不用管他，如果不存在则会创建交换机
	err := r.channel.ExchangeDeclare(
		//交换机名称
		r.Exchange,
		//类型  路由模式下我们需要将类型设置为direct类型
		"direct",
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		true,
		//是否为自动删除  这里解释的会更加清楚：https://blog.csdn.net/weixin_30646315/article/details/96224842?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase
		false,
		//true表示这个exchange不可以被客户端用来推送消息，仅仅是用来进行exchange和exchange之间的绑定
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	r.failOnErr(err,"Failed to declare an excha " + "nge")
	//2.试探性创建队列，这里注意队列名称不要写哦
	q,err := r.channel.QueueDeclare(
		//随机生产队列名称 这个地方一定要留空
		"",
		false,
		false,
		//具有排他性   排他性的理解 这篇文章还是比较好的：https://www.jianshu.com/p/94d6d5d98c3d
		true,
		false,
		nil,
	)
	r.failOnErr(err,"Failed to declare a queue")
	//3.绑定队列到exchange中去
	err = r.channel.QueueBind(
		q.Name,//队列的名称  通过key去找绑定好的队列
		//在路由模式下，这里的key要填写
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	//4.消费代码
	//4.1接收队列消息
	message,err := r.channel.Consume(
		//队列名称
		q.Name,
		//用来区分多个消费者
		"",
		//是否自动应答 意思就是收到一个消息已经被消费者消费完了是否主动告诉rabbitmq服务器我已经消费完了你可以去删除这个消息啦 默认是true
		true,
		//是否具有排他性
		false,
		//如果设置为true表示不能将同一个connection中发送的消息传递给同个connectio中的消费者
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//4.2真正开始消费消息
	forever := make(chan bool)
	go func() {
		for d:=range message {
			log.Printf("Received a message: %s",d.Body)
		}
	}()
	fmt.Println("退出请按 ctrl+c")
	<-forever
}

//topic主题模式step1:创建RabbitMQ实例
func NewRabbitMQTopic(exchange string,routingkey string) *RabbitMQ {
	//创建RabbitMQ实例
	return NewRabbitMQ("",exchange,routingkey)
}

//topic主题模式step2:发送消息
func (r *RabbitMQ) PublishTopic(message string){
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		//交换机名称
		r.Exchange,
		//类型 topic主题模式下我们需要将类型设置为topic
		"topic",
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		true,
		//是否为自动删除  这里解释的会更加清楚：https://blog.csdn.net/weixin_30646315/article/details/96224842?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase
		false,
		//true表示这个exchange不可以被客户端用来推送消息，仅仅是用来进行exchange和exchange之间的绑定
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	r.failOnErr(err,"Failed to declare an excha " + "nge")
	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		//除了设置交换机这也要设置绑定的key值
		r.Key,
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			Body:            []byte(message),//发送的内容一定要转换成字节的形式
		})
}

//topic主题模式step2:消费者
//要注意key 规则
//其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
//匹配 huxiaobai.* 表示匹配 huxiaobai.hello 但是huxiaobai.one.two 需要用huxiaobai.# 才能匹配到
func (r *RabbitMQ) RecieveTopic(){
	//1.尝试创建交换机exchange 如果交换机存在就不用管他，如果不存在则会创建交换机
	err := r.channel.ExchangeDeclare(
		//交换机名称
		r.Exchange,
		//类型 topic主题模式下我们需要将类型设置为topic
		"topic",
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		true,
		//是否为自动删除  这里解释的会更加清楚：https://blog.csdn.net/weixin_30646315/article/details/96224842?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase
		false,
		//true表示这个exchange不可以被客户端用来推送消息，仅仅是用来进行exchange和exchange之间的绑定
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	r.failOnErr(err,"Failed to declare an excha " + "nge")
	//2.试探性创建队列，这里注意队列名称不要写哦
	q,err := r.channel.QueueDeclare(
		//随机生产队列名称 这个地方一定要留空
		"",
		false,
		false,
		//具有排他性   排他性的理解 这篇文章还是比较好的：https://www.jianshu.com/p/94d6d5d98c3d
		true,
		false,
		nil,
	)
	r.failOnErr(err,"Failed to declare a queue")
	//3.绑定队列到exchange中去
	err = r.channel.QueueBind(
		q.Name,//队列的名称  通过key去找绑定好的队列
		//在路由模式下，这里的key要填写
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	//4.消费代码
	//4.1接收队列消息
	message,err := r.channel.Consume(
		//队列名称
		q.Name,
		//用来区分多个消费者
		"",
		//是否自动应答 意思就是收到一个消息已经被消费者消费完了是否主动告诉rabbitmq服务器我已经消费完了你可以去删除这个消息啦 默认是true
		true,
		//是否具有排他性
		false,
		//如果设置为true表示不能将同一个connection中发送的消息传递给同个connectio中的消费者
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//4.2真正开始消费消息
	forever := make(chan bool)
	go func() {
		for d:=range message {
			log.Printf("Received a message: %s",d.Body)
		}
	}()
	fmt.Println("退出请按 ctrl+c")
	<-forever
}
```

topicOne.go 生产者：

```go
//topic主题模式生产者
package main

import (
   "fmt"
 "rbtmq/rbtmqcs" "strconv" "time")

func main(){
   rabbitmqOne := rbtmqcs.NewRabbitMQTopic("hxbExc","huxiaobai.one")
   rabbitmqTwo := rbtmqcs.NewRabbitMQTopic("hxbExc","huxiaobai.two.cs")
   for i:=0;i<=10;i++{
      rabbitmqOne.PublishTopic("hello huxiaobai one" + strconv.Itoa(i))
      rabbitmqTwo.PublishTopic("hello huxiaobai two" + strconv.Itoa(i))
      time.Sleep(1 * time.Second)
      fmt.Println(i)
   }
}
```



topicTwo.go 消费者一

```go
package main


import "rabbitmq/rbtmqcs"

func main(){
	//#号表示匹配多个单词 也就是读取hxbExc交换机里面所有队列的消息
	rabbitmq := rbtmqcs.NewRabbitMQTopic("hxbExc","#")
	rabbitmq.RecieveTopic()
}
```

topicThree 消费者二

```go
package main

import "rabbitmq/rbtmqcs"

func main(){
	//这里只是匹配到了huxiaobai.后边只能是一个单词的key 通过这个key去找绑定到交换机上的相应的队列
	rabbitmq := rbtmqcs.NewRabbitMQTopic("hxbExc","huxiaobai.*.cs")
	rabbitmq.RecieveTopic()
}
```



通过下图你会发现 #是属于匹配全部的哦 把交换机上所有绑定好的队列都消费了 另外一个只是匹配到了一个队列 那么就会消费这一个队列：

[![Rabbitmq工作模式之topic模式](https://cdn.learnku.com/uploads/images/202006/22/58888/Wqrdqou133.png!large)](https://cdn.learnku.com/uploads/images/202006/22/58888/Wqrdqou133.png!large)



[![Rabbitmq工作模式之topic模式](https://cdn.learnku.com/uploads/images/202006/22/58888/JWMqEwvtCQ.png!large)](https://cdn.learnku.com/uploads/images/202006/22/58888/JWMqEwvtCQ.png!large)

# 远程调用RPC

```go
 rpc_server.go 
package main

import (
        "fmt"
        "log"
        "strconv"

        "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
        if err != nil {
                log.Fatalf("%s: %s", msg, err)
        }
}

func fib(n int) int {
        if n == 0 {
                return 0
        } else if n == 1 {
                return 1
        } else {
                return fib(n-1) + fib(n-2)
        }
}

func main() {
        conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
        failOnError(err, "Failed to connect to RabbitMQ")
        defer conn.Close()

        ch, err := conn.Channel()
        failOnError(err, "Failed to open a channel")
        defer ch.Close()

        q, err := ch.QueueDeclare(
                "rpc_queue", // name
                false,       // durable
                false,       // delete when usused
                false,       // exclusive
                false,       // no-wait
                nil,         // arguments
        )
        failOnError(err, "Failed to declare a queue")

        err = ch.Qos(
                1,     // prefetch count
                0,     // prefetch size
                false, // global
        )
        failOnError(err, "Failed to set QoS")

        msgs, err := ch.Consume(
                q.Name, // queue
                "",     // consumer
                false,  // auto-ack
                false,  // exclusive
                false,  // no-local
                false,  // no-wait
                nil,    // args
        )
        failOnError(err, "Failed to register a consumer")

        forever := make(chan bool)

        go func() {
                for d := range msgs {
                        n, err := strconv.Atoi(string(d.Body))
                        failOnError(err, "Failed to convert body to integer")

                        log.Printf(" [.] fib(%d)", n)
                        response := fib(n)

                        err = ch.Publish(
                                "",        // exchange
                                d.ReplyTo, // routing key
                                false,     // mandatory
                                false,     // immediate
                                amqp.Publishing{
                                        ContentType:   "text/plain",
                                        CorrelationId: d.CorrelationId,
                                        Body:          []byte(strconv.Itoa(response)),
                                })
                        failOnError(err, "Failed to publish a message")

                        d.Ack(false)
                }
        }()

        log.Printf(" [*] Awaiting RPC requests")
        <-forever
}
```



```go
 rpc_client.go
package main

import (
        "fmt"
        "log"
        "math/rand"
        "os"
        "strconv"
        "strings"
        "time"

        "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
        if err != nil {
                log.Fatalf("%s: %s", msg, err)
        }
}

func randomString(l int) string {
        bytes := make([]byte, l)
        for i := 0; i < l; i++ {
                bytes[i] = byte(randInt(65, 90))
        }
        return string(bytes)
}

func randInt(min int, max int) int {
        return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
        conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
        failOnError(err, "Failed to connect to RabbitMQ")
        defer conn.Close()

        ch, err := conn.Channel()
        failOnError(err, "Failed to open a channel")
        defer ch.Close()

        q, err := ch.QueueDeclare(
                "",    // name
                false, // durable
                false, // delete when usused
                true,  // exclusive
                false, // noWait
                nil,   // arguments
        )
        failOnError(err, "Failed to declare a queue")

        msgs, err := ch.Consume(
                q.Name, // queue
                "",     // consumer
                true,   // auto-ack
                false,  // exclusive
                false,  // no-local
                false,  // no-wait
                nil,    // args
        )
        failOnError(err, "Failed to register a consumer")

        corrId := randomString(32)

        err = ch.Publish(
                "",          // exchange
                "rpc_queue", // routing key
                false,       // mandatory
                false,       // immediate
                amqp.Publishing{
                        ContentType:   "text/plain",
                        CorrelationId: corrId,
                        ReplyTo:       q.Name,
                        Body:          []byte(strconv.Itoa(n)),
                })
        failOnError(err, "Failed to publish a message")

        for d := range msgs {
                if corrId == d.CorrelationId {
                        res, err = strconv.Atoi(string(d.Body))
                        failOnError(err, "Failed to convert body to integer")
                        break
                }
        }

        return
}

func main() {
        rand.Seed(time.Now().UTC().UnixNano())

        n := bodyFrom(os.Args)

        log.Printf(" [x] Requesting fib(%d)", n)
        res, err := fibonacciRPC(n)
        failOnError(err, "Failed to handle RPC request")

        log.Printf(" [.] Got %d", res)
}

func bodyFrom(args []string) int {
        var s string
        if (len(args) < 2) || os.Args[1] == "" {
                s = "30"
        } else {
                s = strings.Join(args[1:], " ")
        }
        n, err := strconv.Atoi(s)
        failOnError(err, "Failed to convert arg to integer")
        return n
}
```

