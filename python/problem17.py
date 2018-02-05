import inflect
import re

p = inflect.engine()
pattern = '\w'

result = 0

for i in range(1,1001):
	word = p.number_to_words(i)
	result += len(re.findall(pattern, word))

print(result)