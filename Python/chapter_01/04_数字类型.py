# 1.数字类型
a = 10
b = 666
c = 3.1415926
print(a, type(a))  # 10
print(b, type(b))  # 666
print(c, type(c))  # 3.1415926

# 2.数字类型间转换
var1 = 3.14
var2 = 5

# 浮点型转整型
var3 = int(var1)
print(var3, type(var3))  # 3

# 整型转浮点型
var4 = float(var2)
print(var4, type(var4))  # 5.0

# 3.常用内置数学方法
# math模块是 内置的一个第三方模块，包含了很多数学计算相关的方法
import math

# 返回数字的上入整数
res = math.ceil(4.1)
print(res)  # 5

# 返回e的x次幂(ex)
res = math.exp(1)  # 2.718281828459045
print(res)

# 返回数字的绝对值
res = math.fabs(-10)
print(res)  # 10.0

# 返回数字的下舍整数
res = math.floor(4.9)
print(res)  # 4

# 计算对数
res = math.log(100, 10)
print(res)  # 2.0

# 计算以10为基数的x的对数
res = math.log10(100)
print(res)  # 2.0

# 返回x的整数部分与小数部分，两部分的数值符号与x相同，整数部分以浮点型表示
res = math.modf(3.14)
print(res)  # (0.14000000000000012, 3.0)

# 返回数字x的平方根
res = math.sqrt(4)
print(res)  # 2.0

# 返回数字的绝对值
res = abs(-10)
print(res)  # 10

# 返回给定参数的最大值，参数可以为序列。
res = max(10, 5)
print(res)  # 10

# 返回给定参数的最小值，参数可以为序列。
res = min(10, 5)
print(res)  # 5

# x**y 运算后的值。
res = pow(2, 3)
print(res)  # 8

# 返回浮点数x的四舍五入值，如给出n值，则代表舍入到小数点后的位数。
# round(x [,n])
res = round(3.14758, 2)
print(res)  # 3.15

# 计算两个数值之间的商和余数
res = divmod(7, 3)
print(res)  # (2, 1)
