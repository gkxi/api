# api

## 1. 初始rst

```shell
cargo init --lib
```

## 2. 拷贝build.rs

## 3. make

## 4. 修改文件名

需要把文件名都修改了，改成一个文件不冲突的模样

## FAQ

### 1. prost生的结构体怎么json

build的时候通过添加宏来序列化，但是有个问题，google一些公共的类型不行，因为是直接引用的prost-types这个库预先生成的，不能自动添加