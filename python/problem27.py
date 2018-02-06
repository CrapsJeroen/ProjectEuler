def is_prime(n):
    i = 2
    while i * i <= n:
        if n % i == 0:
            return False
        i = i + 1
    return True


def generate_prime_from(n):
    for i in range(n, 1000, 1):
        if is_prime(i):
            yield i


def generate_numbers(aa, bb):
    for n in range(0, 1000, 1):
        yield n**2 + aa*n + bb


longest = (0, 0, 0)
for a in range(-1000, 1001):
        for b in range(1, 1001):
            if is_prime(b):
                tmp = 0
                while True:
                    if is_prime(n**2 + aa*n + bb):
                        tmp += 1
                        continue

                    if tmp > longest[0]:
                        longest = (tmp, a, b)
                    break

print(longest[1]*longest[2])
