def exp(b, n):
    if n == 0:
        return 1
    else:
        return b * exp(b, n-1)

def sqaure(x):
    return x * x

def exp_fast(b, n):
    if n == 0:
        return 1
    elif n % 2 == 0:
        return sqaure(exp_fast(b, n//2))
    else:
        return b * exp_fast(b, n-1)
