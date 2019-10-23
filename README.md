TODO-List
1. 解析
   - 支持解析go文件的代码结构，解析其定义变量/类型，并生成json文件 
2. 分析
   - 可以读取上一步生成的json文件，支持在线编辑代码一级结构
   - 支持的拆分/合并go文件
3. 生成
   - 将指定文件的可以拆分成多个文件，或将文件进行合并
   - 将文件合并成包，并替换所有文件其引用以及的包的引用
   - 支持版本管理，可以回退

依赖
1. gofmt
2. goimports


todo
1. ~~map嵌套定义~~
2. ~~func解析变量~~
3. ~~interface解析~~
4. ~~函数体解析变量~~
5. type区分指针
6. 拓扑排序
7. 解析类如`type Token int`的type定义
8. ~~解析called函数~~
9. 替换变量可以用gofmt
10. 生成代码还是得手写
11. 可以解析出一个未定义类型，是未定义类型的其实都是其他文件的，然后根据这个生成图
12. ~~链式调用的时候注意区分属性还是func~~
13. 其实可以加个value字段然后在最后重新根据value构建的type
14. 函数调用参数为定义函数
