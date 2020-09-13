# -*- coding: utf-8 -*-
"""
Created on Sun May 28 15:30:17 2017

@author: Ewen
"""
import argparse
import sys
import time
from parser import parse, pretty_print, unparse

import numpy as np
import rich.traceback
from matplotlib import pyplot as plt

from solver import solve


def solve_and_display(path="sudoku.txt", graph_display=False, iterate=False, fastest=False):
    timestamp_unparsed=time.time()

    # Opening file
    to_solve = open(path, "r")
    S_init = ""
    for line in to_solve:
        S_init += line
    to_solve.close() 


    #Parsing
    time_message = "\nSudoku solving speed:\n"
    S_init = parse(S_init)
    timestamp_parsed=time.time()
    time_parsing = int((timestamp_parsed-timestamp_unparsed)*10000)/10000
    time_message+="Parsing... {} seconds\n".format(time_parsing)

    S = np.copy(S_init)
    
    #Solving
    timestamp_unsolved=time.time()
    S, meta, init = solve(S, display_iterations=iterate, stats=not(fastest))
    timestamp_solved=time.time()
    time_solving = int((timestamp_solved-timestamp_unsolved)*1000)/1000

    pretty_print(S_init, liste_init=init)
    pretty_print(S, liste_init=init)

    # Writing
    # Just checking if there is a .txt extension or not (can disturb Windows dumb users)
    if len(path) > 4:
        if path[-4:] == ".txt":
            path = path[:-4]
    solved  = open(path+"_solved.txt", "w")
    s = unparse(S)
    solved.write(s)
    solved.close()

    # Displaying
    time_message+="Solving... {} seconds\n".format(time_solving)
    print(time_message)
    print("Find the answer above or in the "+path+" file!")

    # Pretty things
    if graph_display:
        print("Watch the figure of the solving evolution.")
        plt.plot([_ for _ in map(lambda x: max(meta)-x, meta)])
        plt.show()


# Appel de la fonction
if __name__ == '__main__':
    rich.traceback.install()

    parser = argparse.ArgumentParser()
    parser.add_argument('--iterate', type=bool, default=False)
    parser.add_argument('--graph', type=bool, default=False)
    parser.add_argument('--inputs', type=str, nargs="+")
    
    args = parser.parse_args()

    print(args)

    for path in args.inputs:

        print("\n=== Reading file "+path+" ===")
        try:
            solve_and_display(path, args.graph, args.iterate)
        except (FileNotFoundError, IOError): 
            print("File doesn't exist or isn't a sudoku")
