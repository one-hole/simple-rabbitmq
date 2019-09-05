### RabbitMQ AMQP 相关的概念

1. Broker

    1. Exchange
        1. direct
        2. topic
        3. fanout
        4. headers

    2. Bindings
    3. Queue

2. Message
    1. Producer
      1. Delivery

    2. Consumer

3. Connection

4. Channel

### RabbitMQ 的工作模式

![RabbitMQWorkFlow](http://ww2.sinaimg.cn/large/006tNc79gy1g619oqhlc5j30vx0hemz8.jpg)

1. Work Queues
    * 对应着上图中的 【3】- 只需要 Q 的名字即可实现

2. Publish / Subscribe （Fanout） - BroadCast
    * 对应这上图中的 【2】- 并且 binding 失效 必须有 X 的名字

3. Routing
    * 【1】需要有 RoutingKey 【2】需要有 BindingKey

4. Topic

5. RPC

### RabbitMQ MQTT 相关概念