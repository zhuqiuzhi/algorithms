def count_partations(n, m):
    """统计将 n 分区, 最大分区大小为 m, 计算出所有的分区方案数
    例如 n = 5, m = 3
    """
    if n == 0:
        return 1
    elif n < 0:
        return 0
    elif m == 0:
        return 0
    else:
        with_m = count_partations(n-m, m)
        without_m = count_partations(n, m-1)
        return with_m + without_m
    