str = string(BigInt(2) ^ 1000)
sum([Int(str[i] - '0') for i = 1:length(str)])
