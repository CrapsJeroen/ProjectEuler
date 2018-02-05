import math

day = 366 % 7

months = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

sunday_count = 0

for year in range(1901,2001):
	for i, m in enumerate(months):
		if day % 7 == 0:
			sunday_count += 1
		day += m
		if i == 1 and year % 4 == 0 and not (year % 400 == 0 != year % 100 == 0):
			day += 1
        

print(sunday_count)
