# 1.算术运算符
a = 2
b = 3

# 加法
print(a + b)

# 减法
print(a - b)

# 乘法
print(a * b)

# 除法
print(a / b)

# 取余
print(a % b)

# 取整
print(a // b)

# 幂运算
print(a ** b)

# 2.比较运算符
a = 2
b = 3

# 等于
print(a == b)

# 不等于
print(a != b)

# 大于
print(a > b)

# 小于
print(a < b)

# 大于等于
print(a >= b)

# 小于等于
print(a <= b)

# 3.逻辑运算符
print(0 and 3)  # 0
print(1 and 3)  # 3
print(2 and 0)  # 0

# 4.赋值运算符
a = 10
b = 5

a = b  # 普通赋值
a += b  # 加法赋值，等价于 a = a + b
a -= b  # 减法赋值，等价于 a = a - b
a *= b  # 乘法赋值，等价于 a = a * b
a /= b  # 除法赋值，等价于 a = a / b
a **= b  # 幂赋值，等价于 a = a ** b
a %= b  # 取余赋值，等价于 a = a % b
a //= b  # 取整赋值，等价于 a = a // b

# 5.成员运算符
a = "aaa"
b = "bbb"
c = "aaaccc"

print(a in c)  # True
print(b in c)  # False

print(a in c)  # False
print(b in c)  # True
