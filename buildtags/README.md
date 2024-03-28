### build <tags> 用法
```
//go:build ignore
* +build ignore

package build_tags
```

> go 编译注解（build tag）-注释里的编译语法
* 在Go中，build tag是添加到我们代码中的第一行，来表示编译相关信息。
* 其决定当前文件是否被当前 package 包含。
* 用于限制一整个文件是否应该被编译入最终的二进制文件，而不是一个文件中的部分代码。
> go 的标签（build tag）语法：`+build [tag]`
* build tags文件顶部附近，前面只能有空行和其他注释。
* 编译标记必须出现在package自居之前，并且为了与包文档区分开来，它必须后跟一个空行。
* 使用多个标签时的标签交互逻辑，标签之间使用bool类型交互（运算）
> 遵循三个原则：
* 以空格分割的标签将在OR逻辑进行解释。
* 以逗号分割的标签将在AND逻辑下进行解释。
* 每个属于都是一个字母数字单次，如果前面有 ! 符号它意味着被否定、
> 举例：
* +build tag1 tag2  * OR 语法，在执行build构建命令时存在tag1或tag2，则将包含此文件
* +build tag1, tag2 * AND 语法，在执行build构建命令时必须同时存在tag1和tag2，此文件才会被加入编译
* +build !tag1	   * ! 语法，编译命令中不包含tag1才会编译当前文件


> go:build 编译指令是1.17引入的新条件编译指令格式，它旨在替换 * +build 
* 为什么要采用新格式？
* go:build linux && amd64 || darwin
* +build linux,amd64 darwin
* 新的格式更加清楚的表示逻关系，开发更友好
* 与 go:embed 和 go:generate 风格上保持统一
