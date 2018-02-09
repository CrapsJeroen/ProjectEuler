st = set()  # new version with a set

for a in range(2, 101):
    for b in range(2, 101):
        c = a ** b
        if c in st: pass
        st.add(c)


print(len(st))