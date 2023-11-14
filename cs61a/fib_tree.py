def tree(label, branches=[]):
    for branch in branches:
        assert is_tree(branch), 'brancher must is tree'
    return [label] + list(branches)

def label(tree):
    return tree[0]

def branches(tree):
    return tree[1:]

def is_tree(tree):
    if type(tree) != list or len(tree) < 1:
        return False
    for branch in branches(tree):
        if not is_tree(branch):
            return False
    return True

def is_leaf(tree):
    # branches is None
    return not branches(tree)

def fib_tree(n):
    """
    构造以第 n 个 fib 数的 fib 树 , 0 是 第 0 个
    0, 1, 1, 2, 3, 5
    """
    if n <= 1:
        # 0, 1
        return tree(n)
    else:
        left, right = fib_tree(n-2), fib_tree(n-1)
        root = label(left) + label(right)
        return tree(root, [left, right])