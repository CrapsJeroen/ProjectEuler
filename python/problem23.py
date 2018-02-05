def sum_divisors(n):
    s = 0
    for i in range(1,n):
        if n % i == 0: s += i
    return s

ab = list(filter(lambda x: x != 0, [i if sum_divisors(i) > i else 0 for i in range(1,28124)]))

lst = [0]*28124
for x in range (0, len(ab)):
    for y in range (x, len(ab)):
            tmp = ab[x]+ab[y]
            if (tmp <= 28123):
                if (lst[tmp] == 0):
                    lst[tmp] = tmp

result = 0
for i, x in enumerate(lst):
    if (x == 0):
        result += i

print(result)
