def sum_divisors(n):
    s = 0
    for i in range(1,n):
        if n % i == 0: s += i
    return s
 
def amicable_pairs(low,high):
    L = [sum_divisors(i) for i in range(low,high + 1)]
    pairs = []
    for i in range(high - low + 1):
        ind = L[i]
        if i + low < ind and low <= ind and ind <= high and L[ind - low] == i + low:
            pairs.append([i+low,ind])
    return pairs
 
def sum_pairs(pairs):
    return sum([sum(pair) for pair in pairs])
 
print(sum_pairs(amicable_pairs(1,10000)))