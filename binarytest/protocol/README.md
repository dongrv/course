###

```
protoc --go_out=./ --go_opt=paths=source_relative  ./*.proto

--go_out=./ 表示所使用的插件名为protoc-gen-xxxx.exe。等号后面表示插件执行后生成到的目标路径。
--go_opt=paths=source_relative 表示生成的.pb.go文件路径不依赖于.proto文件中的option go_package配置项，直接在go_out指定的目录下生成.pb.go文件（.pb.go文件的package名还是由option go_package决定）。
```