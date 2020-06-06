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


def nb_possible(S, i, j):
    if (i, j) in liste_init(S):
        return 0
    else:
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
        for l in range(3*int(j/3), 3*int(j/3)+3):
            if k != i and S[k][l]==n:
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
            
def solve(S):
    tab_ordre = tableau_ordre(S)
    nombre_a_changer = len(tab_ordre)
    tab_mini = S
    rang = 0
    valeur_remplacee = 1
    while rang < nombre_a_changer:
        (i, j) = tab_ordre[rang]
        valeur_remplacee = assigne_valeur(S, i, j, tab_mini[i][j] +1) # assigne une valeur à la case (i, j) si possible et récupère la clé de l'op
        if valeur_remplacee != 0: #on continue
            rang += 1
        else: #on recule
            rang -= 1
        tab_mini[i][j] = valeur_remplacee
    return S