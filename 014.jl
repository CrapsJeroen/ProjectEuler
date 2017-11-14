function collatz(n::Int)
    if (n%2 != 1)
        trunc(Int, n/2)
    else
        3n+1
    end
end

function getlength(x::Int)
    i = 1
    while x > 1
        x = collatz(x)
        i += 1
    end
    i
end

cache = zeros(Int32, 1000000)
cache[1] = 1
max = [1,1]

for x in 1:1000000
    n,s = x,0
    
    TO_ADD = []
    while n > 999999 || cache[n] < 1
        push!(TO_ADD,n)
        if n % 2 == 0
            n = Int(n/2) 
        else
            n = 3*n + 1 
        end
        s += 1
    end
    
    for j in 1:s
        m = TO_ADD[j]
        if m < 1000000
            new_length = cache[n] + s - j + 1
            cache[m] = new_length
            if new_length >= max[2]
                max = [x,new_length]
            end
        end
    end
end
print(max)
