# -*- coding: utf-8 -*-
"""
Created on Sun May 28 15:30:17 2017

@author: Ewen
"""
import numpy as np
import time
#from sudokus import S
from parser import parse, unparse
from solver import solve
import sys

def solve_and_display(path="sudoku.txt", display_speed = True):
    timestamp_unparsed=time.time()

    # Opening file
    to_solve = open(path, "r")
    S = ""
    for line in to_solve:
        S += line
    to_solve.close() 

    #Parsing
    time_message = "\nSudoku solving speed:\n"
    if type(S) is str:
        S = parse(S)
        timestamp_parsed=time.time()
        time_parsing = int((timestamp_parsed-timestamp_unparsed)*10000)/10000
        time_message+="Parsing... {} seconds\n".format(time_parsing)
    else:
        time_message
    
    #Solving
    timestamp_unsolved=time.time()
    solve(S)
    timestamp_solved=time.time()
    time_solving = int((timestamp_solved-timestamp_unsolved)*1000)/1000


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
    print(S) # Solved!
    if display_speed: print(time_message)
    print("Find the answer above or in the "+path+" file!")

    return

# Appel de la fonction
if __name__ == '__main__':
    if len(sys.argv) > 1:
        for path in (sys.argv[1:]):
            print("\n=== Reading file "+path+" ===")
            try:
                solve_and_display(path)
            except: 
                print("File doesn't exist or isn't a sudoku")
    else:
        print("No sudoku given")
