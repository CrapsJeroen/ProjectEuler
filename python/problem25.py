import math


def fibonacciGen():
    f1, f2 = 0, 1
    while True:
        yield f1
        f1, f2 = f2, f1 + f2


def nDigits(n):
    if n > 0:
        return int(math.log10(n)) + 1
    elif n == 0:
        return 1


fGen = fibonacciGen()
for i, f in enumerate(fGen):
    if nDigits(f) == 1000:
        print(i)
        break
