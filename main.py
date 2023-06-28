import matplotlib.pyplot as plt
import numpy as np
from sympy import *
import math


#Построение графика
x = np.linspace(0, 1, 500)
X = np.linspace(-0.5, 1, 500)
Y = np.linspace(0, 0, 500)
y = np.log(x+1)
z = 2*x - 0.5
ax = plt.axes()
ax.arrow(0, 0, 1, 0, head_width=0.05, color='black')

plt.plot(x, y, label='log')
plt.plot(x, z, label='linear')
plt.plot(x, Y, color='black')
plt.plot(Y+0.6, X, color='red', linestyle='dashed', label='limitations')
plt.plot(Y+0.4, X, color='red', linestyle='dashed')
plt.legend()
plt.show()

a = float(input("Левая граница:\n"))
b = float(input("Правая граница:\n"))
e = float(input("Введите точность:\n"))
seed = b
seedn = 0
print("Метод Ньютона\n")
def func(x):
    return math.log(x+1) - 2*x + 0.5
def fderiv(x):
    return -2 + 1/(x + 1)
def sderiv(x):
    return -1/(x + 1)**2
print("f(x) = ln(x+1) – 2x + 0.5") #
print("f'(x) = 1/(x+1) - 2")
print("f''(x) = -1/(x+1)^2")
print("f(a)f(b)= {:.6f} < 0".format(func(a)*func(b)))
print("f(x)f''(x)= {:.6f} > 0".format(func(b)*sderiv(b)))

for i in np.arange(b, a, -0.1):
    if (func(i)*sderiv(i) > 0):
        seed = i
        break
print("| k |     Xk     |   f(Xk)   |   f'(Xk)  |-f(Xk)/f'(Xk)|")
count = 1
print("| {:d} |  {:.6f}  | {:.6f} | {:.6f} | {:.6f} |".format(0,seed,func(seed),fderiv(seed),-func(seed)/fderiv(seed)))
while True:
    seedn = seed - func(seed)/fderiv(seed)
    print("| {:d} |  {:.6f}  | {:.6f} | {:.6f} | {:.6f} |".format(count,seedn,func(seedn),fderiv(seedn),-func(seedn)/fderiv(seedn)))
    count = count + 1
    if (abs(seedn - seed) < e):
        print("Полученное приближение:\n", seedn)
        break
    seed = seedn

print("Метод простых итерации:\n")
e = float(input("Введите точность:\n"))

def f1(x):
    return (math.log(x+1)+0.5)/2 # Correct

def f2(x):
    return math.exp(2*x-0.5) -1

def df(x):
    return 0.5/(x+1)

for i in np.arange(a,b,0.01):
    if (f1(i) < a or f1(i) > b):
        f = f2
        break
    else:
        f = f1
        break
max = math.fabs(df(a+0.00000001))
for i in np.arange(a+0.00000001,b-0.00000001,0.01):
    if (max > math.fabs(df(i))):
        max = math.fabs(df(i))
    if (max > 1):
        print("Условия не выполнены:")
        break
print("|phi'(x)| <= {:.6f} < 1".format(max))
k = 0
seed = (a+b)/2
print("|    k     |      x(k)      |   phi(x(k))  |")

while True:
    seedn = f(seed)
    print("|    {:d}     |    {:.6f}    |   {:.6f}   |" .format(k,seed,seedn))
    if (max*math.fabs(seed-seedn)/(1-max) <= e):
        print("Полученное приближение:\n", seedn)
        break
    k = k+1
    seed = seedn
