# 1.方式一 %s
# 按照位置和占位符一一对应，按照传值顺序进行替换，多了或少了都不行
name = "cdc"
age = 18

# 单个值
info1 = "my name is %s" % name

# 多个值
info2 = "my name is %s and I'm %d years old" % (name, age)
info3 = "my name is %s and I'm %s years old" % (name, age)

# 2.方式二 format
# 按照位置传值
name = "cdc"
age = 18
info = "my name is {} and I'm {} years old".format(name, age)

# 按照索引传值
info = "I'm {1} years old and my name is {0}".format(name, age)

# 使用关键字或者字典方式传值，打破位置限制
info = "I'm {name} years old and my name is {age}".format(name="cdc", age=18)

kwargs = {'name': 'cdc', 'age': 18}
info = "my name is {name} and I'm {age} years old".format(**kwargs)  # 使用**进行解包操作


# 方式三 f
# 该方法是从 Python 3.6 版本才开始引入使用的
name = "cdc"
age = 18
info = f"my name is {name} and I'm {age} years old"