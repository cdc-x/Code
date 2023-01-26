# 1的布尔值为True，0的布尔值为False
print(bool(1))  # True
print(bool(0))  # False

# True的布尔值为True，False的布尔值为False
print(bool(True))  # True
print(bool(False))  # False

# 语句条件成立的情况下布尔值为True,否则为False
print(4 > 2)  # True

# 变量为空时布尔值为False,变量有值的时候布尔值为True
s1 = ""
s2 = "aaa"
print(bool(s1))  # True
print(bool(s1))  # False

# 布尔值也可以直接进行数学运算，因为True相当于1,False相当于0
print(True + 1)  # 2

# 与或非操作：
bool(1 and 0)
bool(1 and 1)
bool(1 or 0)
bool(not 0)

# 布尔值经常用在条件判断中:
age = 18
if age > 18:  # 等价于bool(age > 18)
    print('old')
else:
    print('young')
