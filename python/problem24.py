import math

noi = 10**6

digits = [str(i) for i in range(10)]

result = 1
order = []

while result != noi:
    tmp = math.factorial(len(digits)-1)
    index = (noi-result) // tmp
    digit = digits.pop(index)
    order.append(digit)
    result += index*tmp

order.append(digits[0])

print(int(''.join(order)))
