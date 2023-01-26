# 变量
# 1.变量赋值
a = 123
print(a)  # 123

# 2.变量地址
a = 10
b = a
a = 666
print(a)  # 5
print(b)  # 10

# 3.命令规则
# 驼峰式命名
MyName = "cdc"

# 下划线方式命名(蛇形命名)
my_name = "cdc"

__ = "cdc"
# print(__)   # cdc

# 查看变量的数据类型
s = "aaaa"
print(type(s))  # <class 'str'>
n = 123
print(type(n))  # <class 'int'>

# 5.常量
PI = 3.141526
# 实际上常量也是变量
PI = 23333
print(PI)
