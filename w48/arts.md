# Week 49 (2019-11-25)

## Algorithm
Basic sorts:

1. Insert 
2. Binary insert
3. Select
4. Bubble
5. Shell 
6. Heap
7. Merge
8. Fast

## Reading
Changes of JsonSchema: [https://json-schema.org/draft/2019-09/release-notes.html]

## Tip
Json schema 2019 makes `format` optional and encourage to do format validation on application level. The default validator supplied by an implementation lib is switched off. 

## Share
通常我们用大O方式表达算法效率，排序算法主要有两个类别(另外还有一个线性时间的如基数排序，本文不讨论)

1. O(N<sup>2</sup>)
	1. 插入排序
	2. 选择排序
	3. 冒泡排序


2. N*log<sub>2</sub>N
	1. 快速排序
	2. 归并排序
	3. 堆排序
	4. 希尔排序
	
3. 下表分别选择了小数集 和 大数据集排序，可以看到一些有趣的现象：
	1. 即使是时间复杂度级别相同的排序方法，效率也会有数倍的差距，绝不可忽略。 
	2. 插入排序对**部分有序**的数据很友好，因为它能够在确认有序的时候停下来。 而选择排序，冒泡排序没法探测到部分有序，不能减少工作量。 
	3. 对小数据集，插入排序有碾压式的优式。尽管它是 O(N<sup>2</sup>)的，比那些 Nlog<sub>2</sub>N还要好。因此在有些分治策略的排序中，分到比较小时即改为插入排序（一般选择64为界限）

	4. 对于大数据集，理论上快速排序在有的情况下会恶化(最坏到（O(N<sup>2</sup>))，但是在平均情况下，插入排序总是优于其它算法。
	5. 归并排序和堆排序在最坏情况下也是对数时间，不会退化。对大数据集，归并排序性能高于堆排序，但是需要消耗1倍的内存。
	

	5. 实际中如果需要快速一般使用 改进的快速排序（分治较小时改为插入排序。）如果需要稳定算法，一般选择归并排序。希尔排序很有趣（而且据说是最早实现的**对数复杂度**的排序算法），有学习价值，但实际中基本不用了。

4.  上一条也不全对，现在实际使用多是Tim-Sort(Python和Java标准库），Tim-Sort是稳定排序，大多数情况下在各数排序算法中是最快的。它本质上是归并排序，但是利用了很多工程实践中观察到的优势。主要有
	1. 对部分有序友好，主动发现要邻的升序数据段，和严格降序数据段（理论上也可以不严格，严格是为了保证排序稳定性）
	2. 分治生成 一系列 **run** （数据段）， 在合适时间合并 **run**， **合适** 一般指的是两个 **run**的长长接近。
	3. 充分利用缓存特性，发现并优先处列相邻数据
	4. 嗯，Tim-Sort的实现复杂度比较大，这里先学习一下它的特性，以后再练手、实现。





1. Sort on an small-size array (32 elements)

|Type| Speed Ratio | Stable|Time Complexity|friendly to partly ordered data|friendly to cache|
|----|-------------|-------|----------------|-----------|--------------|
|insert(ordered)| 1 |Y|O(N^2)| Y | Y |
|insert|16 |Y|O(N^2)|Y | Y |
|fast-insert|16|N|Nlog<sub>2</sub>N|Y | Y |
|fast|25 |N|Nlog<sub>2</sub>N|Y | Y |
|binary insert|32 |Y|O(N^2)|Y | N |
|shell| 36|N|Nlog<sub>2</sub>N(roughly)|Y|Y|
|heap| 37|N|Nlog<sub>2</sub>N|N|N|
|bubble| 44|Y|O(N^2)|N|Y|
|merge|46 |Y|Nlog<sub>2</sub>N|N|Y|
|select| 53|N|O(N^2)|N|Y|

2. Sort on an big-size array (32000 elements, 1000 times of the small array)



|Type| Speed Ratio |
|----|-------------|
|insert(ordered)| 1 |
|fast-insert|69
|fast| 78|
|merge|124 |
|shell|155 |
|heap| 183|
|binary insert| 10439|
|insert|16100 |
|select| 17245|
|bubble| 53234|




