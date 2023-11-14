def make_withdraw_list(balance):
    # 注意这里只能使用 list 或者 dict
    # 否则会报错: cannot access local variable 'b' where it is not associated with a value
    b = [balance]
    def withdraw(amount):
        if amount > b[0]:
            return 'Insufficient funds'
        b[0] = b[0] - amount
        return b[0]
    return withdraw
