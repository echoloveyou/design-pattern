package facade_pattern

/*
=================外观模式===================
	1.定义：为多个复杂的子系统提供一个一致的接口，使这些子系统更加容易被访问。
	2.这个看了定义应该理解一半了。其实就是将一些复杂的方法包装成一个方法方便别人调用
		就拿简单的购物来说吧，要经过一系列复杂的过程，首先要查询库存是否有商品，然后计算金额，清空购物车，生成订单这只是简单的步骤
		这个例子略显粗糙，但是就是那么个意思，重点是让大家理解这个设计模式
		可以先看一下orderService 结构体继承了goodsService和carService 在SubOrder的方法中我们调用了各种方法一堆逻辑，搞得我晕头转向。
		但是最终我给出了只有一个方法，就是一个提交订单的操作，实际上相当的复杂
*/
