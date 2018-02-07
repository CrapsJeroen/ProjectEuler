tmp = 1
t = 1

for i in range(3, 1002, 2):
    for j in range(1, 5):
        t += (i-1)
        tmp += t

print(tmp)
