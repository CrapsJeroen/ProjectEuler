lst = []

for i in range(2, 354294):
        total = 0
        for j in str(i):
                total += int(j) ** 5

        if total == i:
                lst.append(i)

print(sum(lst))