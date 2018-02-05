import math
import cProfile

from sympy.ntheory import factorint

lst = [0] * ((10000 * 10000) + 1)


def d(n):
    dct = factorint(n)
    res = 1
    for key, value in dct.items():
        res *= int(((key ** (value + 1)) - 1) / (key - 1))
    return res


def s(n):
    res = 0
    for i in range(1, n + 1):
        for j in range(1, n + 1):
            res = res % (10 ** 9)
            tmp = i * j
            if tmp <= 1000000:
                if lst[tmp] == 0:
                    lst[tmp] = d(tmp)
                res += lst[tmp]
            else:
                res += d(tmp)
    return res


cProfile.run('s(int(10**4))')  # -> 563576517282
# print(s(int(10**5))%10**9) #-> 215766508
