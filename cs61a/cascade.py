def cascade(n):
    """print casade(级联)
    >>> cascade(123)
    123
    12
    1
    12
    123
    >>> cascade(123456)
    123456
    12345
    1234
    123
    12
    1
    12
    123
    1234
    12345
    123456
    """
    if n < 10:
        print(n)
    else: 
        print(n) 
        cascade(n//10)
        print(n)

def reverse_cascade(n):
    """print reverse cascade
    >>> reverse_cascade(1234)
    1
    12
    123
    1234
    123
    12
    1
    """
    grow(n)
    print(n)
    shrink(n)

def f_then_g(f, g, n):
    if n:
        f(n)
        g(n)

grow=lambda n: f_then_g(grow, print, n//10)
shrink=lambda n: f_then_g(print, shrink, n//10)