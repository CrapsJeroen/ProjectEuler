def is_prime(n):
    i = 2
    while i * i <= n:
        if n % i == 0:
            return False
        i = i + 1
    return True


def generate_prime_below(n):
    for i in range(n, 1, -1):
        if is_prime(i):
            yield i


# According to Fermat's theorem, length of the repeated part is given as follows:
# 10^n mod P = 1
# now, this n can be either P-1 or factor of P-1
for p in generate_prime_below(1000):
    c = 1
    while pow(10, c, p) != 1:
        c = c + 1

    if p - c == 1:
        print(p)
        break
