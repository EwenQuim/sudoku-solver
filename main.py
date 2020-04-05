# -*- coding: utf-8 -*-
"""
Created on Sun May 28 15:30:17 2017

@author: Ewen
"""
import numpy as np
import time
from sudokus import S
from parser import parse, unparse
from solver import solve

def solve_and_display(path="to_solve.txt", display_speed = True):
    # Opening file
    to_solve = open(path, "r")
    S = ""
    for line in to_solve:
        S += line
    to_solve.close() 

    #Parsing
    time_message = "\nSudoku solving speed:\n"
    if type(S) is str:
        timestamp_unparsed=time.clock()
        S = parse(S)
        timestamp_parsed=time.clock()
        time_parsing = int((timestamp_parsed-timestamp_unparsed)*1000)/1000
        time_message+="Parsing... {} seconds\n".format(time_parsing)
    else:
        time_message
    
    #Solving
    timestamp_unsolved=time.clock()
    solve(S)
    timestamp_solved=time.clock()
    time_solving = int((timestamp_solved-timestamp_unsolved)*1000)/1000

    # Displaying
    time_message+="Solving... {} seconds\n".format(time_solving)
    print(S) # Solved!
    if display_speed: print(time_message)
    print("Find the answer in the solved.txt file or above!")

    # Writing
    solved  = open("solved.txt", "w")
    s = unparse(S)
    solved.write(s)
    solved.close()
    return

# Appel de la fonction
if __name__ == '__main__':
    solve_and_display()
