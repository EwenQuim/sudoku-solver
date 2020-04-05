import numpy as np

def parse(s):
    matrix = np.zeros((9,9), dtype=int)
    lines=s.split("\n")
    for i in range(9):
        for j in range(9):
            if lines[i][j] != ".":
                matrix[i][j] = lines[i][j]
    return matrix

def unparse(s_array):
    s = ""
    for i in range(9):
        for j in range(9):
            s += str(s_array[i][j])
        s += "\n"
    return s