###  [Protobuf](https://protobuf.dev/)
<details>
<summary>Protocol Buffers（Protobuf）是 Google 推出的一种数据序列化协议</summary>
<p><strong>高效：</strong>Protobuf 用二进制格式搞定数据序列化，体积小、编解码速度快。</p>
<p><strong>紧凑：</strong>Protobuf 用可变长度编码，压缩数据无压力，节省带宽和存储空间。</p>
<p><strong>跨平台：</strong>Protobuf 支持多种编程语言，如 Golang、Java、C++、Python 等。</p>
<p><strong>易维护：</strong>Protobuf 用自描述的数据结构，理解和维护轻松无负担。</p>
</details>

##### 拓展阅读：
- [编码方式](https://protobuf.dev/programming-guides/encoding/)
- [最佳实践](https://protobuf.dev/programming-guides/dos-donts/)

---

### 开发要点
> ##### 兼容性
>> - 无符号整数同类型由小改大，兼容，举例：uint32 -> uint64
>> - 无符号整数同类型由大改小，不兼容，举例：uint64 -> uint32
>> - 有符号整数同类型由小改大，兼容，举例：int32 -> int64
>> - 有符号整数同类型由大改小，不兼容，举例：int64 -> int32
>> - 无符号和有符号类型互相转换时，如果通讯中存在负数值，不兼容，举例：int32 -> uint32
>> - string类型改整数，不兼容，字段值失效，反之亦然；同理bytes类型也不能修改为其他类型
>> - message中的字段绑定序号修改，不兼容，含义和数值都不匹配
>> - 修改message中的字段名称，兼容，不影响proto序列化和反序列化
>> - 只修改message中的字段行的位置，不修改绑定的序号，兼容
>>> ```text
>>> // 修改前
>>> message Foo {
>>>   int32 id    = 2;
>>>   string text = 1;
>>> }
>>> // 修改后
>>> message Foo {
>>>   string text = 1;
>>>   int32 id    = 2;
>>> }
>>> ```

---

> ##### 注意事项
>> - 已经上线的协议：message中的字段序号、字段类型、枚举类型都不应该再修改
>> - 整型同类型修改也要谨慎，尽量不修改，除非你能承担错误的代价
>> - 对已经弃用的字段标记`[deprecated=true]`，比如：
>>> ```text
>>> string text = 1 [deprecated=true];
>>> ```
>> - 对已经弃协议，在头部标记`deprecated`，比如：
>>> ```text
>>> // deprecated 
>>> message Foo {
>>>     string text = 1;
>>> }
>>> ```
>> - 如果确实需要修改字段类型，使用添加新类型字段过渡的形式，处理字段类型修改


