# verb-rank
Find common verb used for coding variables

我发现80%的时间花在实现上，而另外的20%的时间则是在思考着如何用一个
合适的变量来清晰的表达意图。

而且时常在纠结所使用的变量是不是真的合适，于是思考着如何去借鉴别人的
想法。因此想着从github上分析最常用的变量，作为日后编程变量名的一个参考。

## 实现的主要功能包括但不局限于：

- 根据语言搜索git repository, ordering by stars
- 下载指定排名的repository, 分析文件中的变量

## 实现步骤

- 集成http功能, go request
- 使用go实现git clone
- 使用go正则或者linux shell实现文件分析，单词统计
- 将分析结果纪录到文件或者redis里面

