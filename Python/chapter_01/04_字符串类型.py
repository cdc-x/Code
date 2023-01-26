# 1. 创建字符串
# 单行字符串，单双引号都可以使用
var1 = 'Hello World!'
var2 = "Python RAlvin"

# 多行字符串
var3 = """
aaa
bbb
ccc
"""

var4 = '''
aaa
bbb
ccc
'''

print(var1, type(var1))
print(var3, type(var3))

# 2.字符串常用方法
# 2.1 * 重复输出字符串
print('hello' * 2)

# 2.2 [] ,[:] 通过索引获取字符串中字符,这里和列表的切片操作是相同的,具体内容见列表
print('helloworld'[2:])

# 2.3 in 成员运算符 如果字符串中包含给定的字符返回 True
print('el' in 'hello')

# 2.4 格式化输出字符串
name = "cdc"
age = 18

# 方式一 %
print('my name is %s' % name)
print("my name is %s and I'm %d years old" % (name, age))

# 方式二 format
print("my name is {} and I'm {} years old".format(name, age))
# 根据索引取值
print("my name is {1} and I'm {0} years old".format(name, age))

# 方式三 f
print(f"my name is {name} and I'm {age} years old")

# 2.5 + 字符串拼接
a = '123'
b = 'abc'
c = '789'
d1 = a + b + c
print(d1)

# + 效率低,建议使用join
d2 = ''.join([a, b, c])
print(d2)
