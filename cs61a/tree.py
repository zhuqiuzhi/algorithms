class Tree:
    def __init__(self, label, branches=[]) -> None:
        self.label = label
        for branch in branches:
            assert isinstance(branch, Tree)
        # create new list
        self.branches = list(branches)

def prun(t, n):
    """Prun all subtres whose lable is n.
    """
    t.branches = [b for b in t.branches if b.label != n]
    for branch in t.branches:
        prun(branch, n)