import numpy as np
from rich import print


def parse(s):
    matrix = np.zeros((9,9), dtype=int)
    lines=s.replace("\n\n", "\n").split("\n") # Avoiding empty lines
    for i in range(9):
        for j in range(9):
            try:
                matrix[i][j] = int(lines[i][j])
            except:
                pass
    return matrix

def unparse(s_array):
    s = ""
    for i in range(9):
        for j in range(9):
            s += str(s_array[i][j])
        s += "\n"
    return s

def pretty_print(S, changed=None, liste_init=[]):
    if not changed:
        changed = (-1, -1)
    for i in range(len(S)):
        for j in range(len(S[0])):
            if j%3 == 0:
                print("|", end="")
            else:
                print(" ", end="")
            if (i, j) in liste_init:
                print("[bold yellow]"+str(S[i, j])+"[/bold yellow]", end="")
            elif (i, j) == changed:
                print("[bold red]"+str(S[i, j])+"[/bold red]", end="")
            elif S[i, j]:
                print("[cyan]"+str(S[i, j])+"[/cyan]", end="")
            else:
                print(".", end="")
        print("")
        if i % 3 == 2:
            print("------------------")
