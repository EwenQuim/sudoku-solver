# -*- coding: utf-8 -*-
"""
Created on Sun May 28 15:30:17 2017

@author: Ewen
"""
import numpy as np
import time

S0 =[[0,0,0,  0,0,0,  0,0,0],
     [0,0,0,  0,0,0,  0,0,0],
     [0,0,0,  0,0,0,  0,0,0],
     
     [0,0,0,  0,0,0,  0,0,0],
     [0,0,0,  0,0,0,  0,0,0],
     [0,0,0,  0,0,0,  0,0,0],
     
     [0,0,0,  0,0,0,  0,0,0],
     [0,0,0,  0,0,0,  0,0,0],
     [0,0,0,  0,0,0,  0,0,0]]
     
S = np.array(   [[0,0,3, 0,0,0, 5,0,7],
                 [8,4,5, 0,0,0, 2,0,0],
                 [0,0,0, 0,1,9, 0,0,0],
                 
                 [3,2,1, 0,7,0, 0,0,0],
                 [4,0,0, 0,0,0, 0,5,3],
                 [0,0,0, 0,4,1, 0,0,0],
                 
                 [0,5,0, 0,0,0, 8,4,0],
                 [2,9,0, 0,0,6, 0,0,0],
                 [1,0,0, 2,5,4, 0,0,0]])

T = np.array(   [[1,0,0,  0,0,7,  0,9,0],
                 [0,3,0,  0,2,0,  0,0,8],
                 [0,0,9,  6,0,0,  5,0,0],
                 
                 [0,0,5,  3,0,0,  9,0,0],
                 [0,1,0,  0,8,0,  0,0,2],
                 [6,0,0,  0,0,4,  0,0,0], #Sudoku "le plus dur du monde" !
                 
                 [3,0,0,  0,0,0,  0,1,0],
                 [0,4,1,  0,0,0,  0,0,7],
                 [0,0,7,  0,0,0,  3,0,0]])

R = np.array(   [[0,1,0,  0,0,0,   0,0,8],
                 [0,0,6,  0,0,5,  9,0,7],
                 [0,0,2,  3,8,0,  6,0,1],
                 
                 [0,0,0,  9,0,0,  0,0,0],
                 [0,0,0,  8,0,0,  5,0,0],
                 [9,8,0,  5,6,7,  1,0,0],
                 
                 [0,0,0,  7,0,0,  0,0,0],
                 [0,0,9,  0,0,0,  0,0,0],
                 [0,0,8,  1,9,0,  2,0,6]])

"""
S = np.array(   [[0,7,5,  0,9,0,  0,0,6],
                 [0,2,3,  0,8,0,  0,4,0],
                 [8,0,0,  0,0,3,  0,0,1],
                 
                 [5,0,0,  7,0,2,  0,0,0],
                 [0,4,0,  8,0,6,  0,2,0],
                 [0,0,0,  9,0,1,  0,0,3],
                 
                 [9,3,0,  4,0,0,  0,0,7],
                 [0,6,0,  0,7,0,  5,8,0],
                 [7,0,0,  0,1,0,  3,9,0]])

S = np.array(   [[8,3,0,  9,0,5,  0,1,7],
                 [0,0,0,  0,0,0,  0,8,3],
                 [7,0,1,  8,4,0,  6,0,0],
                 
                 [4,0,0,  1,0,6,  0,0,0],
                 [1,0,0,  0,0,0,  9,0,0],
                 [0,0,6,  0,0,0,  8,7,0],
                 
                 [0,9,7,  0,0,0,  0,2,0],
                 [6,4,0,  0,0,7,  0,5,0],
                 [0,0,0,  0,0,9,  0,4,8]])"""


#####################   Initialisation   ###################################
# Appel unique: complexité constante
   
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

#####################   Résolution   ###################################
# Appel multiples


# Check si possible
def is_in_bloc(S, i, j, n):
    for k in range(3*int(i/3), 3*int(i/3)+3):
        for l in range(3*int(j/3), 3*int(j/3)+3):
            if k != i and S[k][l]==n:
                return False
    return True
# En complexite constante 9*3
def is_available(S, i, j, n):
    for k in range(9):
        if S[i][k] == n and k != j:
            return False
    for k in range(9):
        if S[k][j] == n and k != i:
            return False
    return is_in_bloc(S, i, j, n)


# Prends une case, cherche le plus petit 
# Complexité pire des cas : 10*(9*3)
def assigne_valeur(S, i, j, mini):
    for n in range(mini, 10):
        if is_available(S, i, j, n):
            S[i][j] = n
            return n
    S[i][j] = 0
    return 0
            
        
def resoudre(S):
    tmps1=time.clock()
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
    tmps2=time.clock()
    tmps = int((tmps2-tmps1)*1000)/1000
    print("Résolution du sudoku : ", tmps, " secondes\n")
    return S
    











