///////////////////////////////消息序列化///////////////////////////////
1，先写入CLSID int
2.写入HEAD
    1）serial int
    2）seq_or_ack int
3.写入CustomAttributes（从子类-父类，获取CustomAttributes，然后根据CustomAttributes获取顺序读取value写入）
    1)写入CustomAttributes的类型（ushort类型）
    2)写入值
    
///////////////////////////////涉及内容///////////////////////////////
1.数据类型，语法
2.interface struct tag对应的json、xml读取
3.goroutine channel
4.package

        