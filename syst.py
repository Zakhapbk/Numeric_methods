import matplotlib.pyplot as plt
import numpy as np
from sympy import *
import math


#Построение графика
x,y = var('x,y')
plot1 = plot_implicit(Eq(x ** 2 / 4 + y ** 2, 1),show=False, line_color="r",adaptive = False)
plot2 = plot_implicit(Eq(2*y - exp(x) - x, 0),show=False,adaptive = False)
axr1 = plot_implicit(Eq(x,0.6),show=False,line_color='green')
axl1 = plot_implicit(Eq(x,0.2),show=False,line_color='green')
ayh1 = plot_implicit(Eq(y+0.000001*x,1.2),show=False,line_color='black')
ayl1 = plot_implicit(Eq(y+0.000001*x,0.8),show=False,line_color='black')
axr2 = plot_implicit(Eq(x,-1.4),show=False,line_color='black')
axl2 = plot_implicit(Eq(x,-1.6),show=False,line_color='black')
ayl2 = plot_implicit(Eq(y+0.000001*x,-0.8),show=False,line_color='green')
ayh2 = plot_implicit(Eq(y+0.000001*x,-0.5),show=False,line_color='green')
plot1.append(plot2[0])
plot1.append(axr1[0])
plot1.append(axl1[0])
plot1.append(ayh1[0])
plot1.append(ayl1[0])
plot1.append(axr2[0])
plot1.append(axl2[0])
plot1.append(ayh2[0])
plot1.append(ayl2[0])
plot1.show()

def det(a,b,c,d):
    return a*d-b*c
print("Метод Ньютона\n")
e = float(input("Введите точность:\n"))
k = int(input("Введите колличество решений:\n"))

for number in range(k):
    xl0 = float(input("Левая граница x:\n"))
    xr0 = float(input("Правая граница x:\n"))
    yl0 = float(input("Левая граница y:\n"))
    yr0 = float(input("Правая граница y:\n"))
    xseed = float((xl0 + xr0) / 2)
    yseed = float((yl0 + yr0) / 2)
    x = Symbol('x')
    y = Symbol('y')
    f1 = ((x**2)/4) + y**2 -1
    f2 = 2*y - exp(x) - x
    f1dx = simplify(f1.diff(x))
    f1dy = simplify(f1.diff(y))
    f2dx = simplify(f2.diff(x))
    f2dy = simplify(f2.diff(y))
    count = 0
    #print("| k  |   x1   |   x2   |   f1   |   f2   | f1dx1 | f2dx1 | f1dx2 | f2dx2 | detA1 | detA2 | detJ |")

    while True:
        #print("|{:d}|   {:.6f}   |   {:.6f}   |   {:.6f}   |   {:.6f}  | {:.6f} | {:.6f} | {:.6f} | {:.6f} | {:.6f} | {:.6f} | {:.6f} |".format(count, xseed, yseed, f1.subs([(x, xseed), (y, yseed)]), f2.subs([(x, xseed), (y, yseed)]),f1dx.subs([(x, xseed), (y, yseed)]), f2dx.subs([(x, xseed), (y, yseed)]),f1dy.subs([(x, xseed), (y, yseed)]), f2dy.subs([(x, xseed), (y, yseed)]),det(f1.subs([(x, xseed), (y, yseed)]), f1dy.subs([(x, xseed), (y, yseed)]),f2.subs([(x, xseed), (y, yseed)]), f2dy.subs([(x, xseed), (y, yseed)])),det(f1dx.subs([(x, xseed), (y, yseed)]), f1.subs([(x, xseed), (y, yseed)]),f2dx.subs([(x, xseed), (y, yseed)]), f2.subs([(x, xseed), (y, yseed)])),det(f1dx.subs([(x, xseed), (y, yseed)]), f1dy.subs([(x, xseed), (y, yseed)]),f2dx.subs([(x, xseed), (y, yseed)]), f2dy.subs([(x, xseed), (y, yseed)]))))
        xseedn = xseed - det(f1.subs([(x,xseed),(y,yseed)]),f1dy.subs([(x,xseed),(y,yseed)]),f2.subs([(x,xseed),(y,yseed)]),f2dy.subs([(x,xseed),(y,yseed)]))/det(f1dx.subs([(x,xseed),(y,yseed)]),f1dy.subs([(x,xseed),(y,yseed)]),f2dx.subs([(x,xseed),(y,yseed)]),f2dy.subs([(x,xseed),(y,yseed)]))
        yseedn = yseed - det(f1dx.subs([(x,xseed),(y,yseed)]),f1.subs([(x,xseed),(y,yseed)]),f2dx.subs([(x,xseed),(y,yseed)]),f2.subs([(x,xseed),(y,yseed)]))/det(f1dx.subs([(x,xseed),(y,yseed)]),f1dy.subs([(x,xseed),(y,yseed)]),f2dx.subs([(x,xseed),(y,yseed)]),f2dy.subs([(x,xseed),(y,yseed)]))
        #print("|{:d}|   {:.6f}   |   {:.6f}   |   {:.6f}   |   {:.6f}  | {:.6f} | {:.6f} | {:.6f} | {:.6f} | {:.6f} | {:.6f} | {:.6f} |".format(count,))
        if (max(abs(xseed - xseedn),abs(yseed - yseedn)) < e):
            break
        xseed = xseedn
        yseed = yseedn
        count = count + 1
    if(xseedn > 0 and yseedn > 0):
        print("Ответ: x1 = {:f} x2 = {:f}, число итераций = {:d}" .format(xseedn,yseedn,count))
        break

x = Symbol('x')
y = Symbol('y')
print("Метод простых итерации:\n")
e = float(input("Введите точность:\n"))
phi1 =  x**2/4*(-0.1) + y**2*(-0.1) - (-0.1) + 2*y*(0.1)-exp(x)*(0.1)-x*(0.1)+x         #    2*sqrt(1-y**2)                  sqrt(16-y**2) 2*sqrt(1-y**2)        log(2*y-x)
phi2 =  x**2/4*(-0.1) + y**2*(-0.1) - (-0.1) + 2*y*(-0.2)-exp(x)*(-0.2)-x*(-0.2)+y      #   (exp(x) + x)/2                 log(x + 4) (exp(x) + x)/2

ph1dx = simplify(phi1.diff(x))
ph1dy = simplify(phi1.diff(y))
ph2dx = simplify(phi2.diff(x))
ph2dy = simplify(phi2.diff(y))



sum1 = 0
max1 = abs(ph1dx.subs([(x,xl0+0.0001), (y,yl0+0.0001)])) + abs(ph1dy.subs([(x,xl0+0.0001 ), (y,yl0+0.0001)]))


for i in np.arange(xl0,xr0,0.01):
    for j in np.arange(yl0,yr0,0.01):
        sum1 = abs(ph1dx.subs([(x, i), (y,j)])) + abs(ph1dy.subs([(x, i), (y,j)]))
        if (sum1 > max1):
            max1 = sum1
        sum1 = 0
    if (max1 > 1):
        print("Процесс расходится")
        break

max2 = abs(ph2dx.subs([(x,xl0+0.0001), (y,yl0+0.0001)])) + abs(ph2dy.subs([(x,xl0+0.0001 ), (y,yl0+0.0001)]))
for i in np.arange(xl0,xr0,0.01):
    for j in np.arange(yl0,yr0,0.01):
        sum1 = abs(ph2dx.subs([(x, i), (y,j)])) + abs(ph2dy.subs([(x, i), (y,j)]))
        if (sum1 > max2):
            max2 = sum1
        sum1 = 0
    if (max2 > 1):
        print("Процесс расходится")
        break

q = max(max1,max2)
xseed = (xl0+xr0)/2
yseed = (yl0+yr0)/2
count = 0
while True:
    xseedn = phi1.subs([(x,xseed),(y,yseed)])
    yseedn = phi2.subs([(x,xseed),(y,yseed)])
    count = count + 1
    if (q*max(abs(xseedn - xseed), abs(yseed - yseedn))/(1-q) < e):
        print("Полученное приближение: x1 = {:.6f} x2 = {:.6f} число итераций = {:d}" .format(xseedn, yseedn, count))
        break
    xseed = float(xseedn)
    yseed = float(yseedn)
