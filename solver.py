import time
from os import system
from parser import pretty_print

import numpy as np

################################
######   Initialisation   ######
################################

# Liste des couples de base à ne pas changer
# En complexite constante 9²
def liste_init(S):              
    l = []
    for i in range(9):
        for j in range(9):
            if S[i][j] != 0:
                l.append((i, j))
    return l

# n° of solutions available for S at (i, j)
def nb_possible(S, i, j):
    if (i, j) in liste_init(S):
        return 0
    
    c = 0
    for n in range(1, 10):
        if is_available(S, i, j, n):
            c += 1
    return c

def tableau_possibilites(S):
    tab = np.array(S)
    for i in range(9):
        for j in range(9):
            n =  nb_possible(S, i, j)
            tab[i][j] = n
    return tab
    
# tableau des points dans l'ordre croissant des possibilités  
def tableau_ordre(S):
    tab = np.array(tableau_possibilites(S))
    liste = []
    for k in range (1, 10):
        for i in range(9):
            for j in range(9):
                if tab[i][j] == k:
                    liste.append((i, j))
    return liste

############################
######   Résolution   ######
############################

# Appel multiples

# Check if possible to fill case (i, j) with n inside a block
def is_in_bloc(S, i, j, n):
    for k in range(3*int(i/3), 3*int(i/3)+3):
        if k != i:
            for l in range(3*int(j/3), 3*int(j/3)+3):
                if l != j and S[k][l]==n:
                        return False
    return True

# Check if possible to fill case (i, j) with n (lines + rows + block)
def is_available(S, i, j, n):
    for k in range(9):
        if (S[i][k] == n and k != j) or (S[k][j] == n and k != i):
            return False
    return is_in_bloc(S, i, j, n)


# Gives the smallest number available 
# Maximal complexity: 10*(9*3)
def assigne_valeur(S, i, j, mini):
    for n in range(mini, 10):
        if is_available(S, i, j, n):
            S[i][j] = n
            return n
    S[i][j] = 0
    return 0
            
def solve(S, display_iterations=False, stats=True):
    sudoku_initial = liste_init(S)
    tab_ordre = tableau_ordre(S)
    n_to_change = len(tab_ordre)
    tab_mini = S
    rank = 0

    # Display
    ranks = []
    iterate_display = 0

    while rank < n_to_change:

        # Computing
        (i, j) = tab_ordre[rank]
        replaced_value = assigne_valeur(S, i, j, tab_mini[i][j] +1) # assigne une valeur à la case (i, j) si possible et récupère la clé de l'op
        if replaced_value != 0: #on continue
            rank += 1
        else: #on recule
            rank -= 1
        tab_mini[i][j] = replaced_value

        # Just displaying stuff
        if display_iterations:
            iterate_display += 1
            if (iterate_display < 50):
                pretty_print(S, changed=(i,j))
                time.sleep(.3)
                system('clear')
        if stats:
            ranks.append(rank)

    return S, ranks, sudoku_initial
