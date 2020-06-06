# Ewen Quimerc'h project

Here is the docker image of my project: <https://hub.docker.com/repository/docker/feust/sudoku-solver>

## What is it?

A simple sudoku solver. You can find the source code [here](https://github.com/EwenQuim/sudoku-solver)

It is shipped in a **Docker image**, so you'll have to run a container (don't worry, the command is juste below). It **mounts** a folder from your computer to the corresponding container's folder containing the sudokus. I could have done it other ways (using just a string for example) but I did it for the challenge and explore a bit Docker's functionalities.

## How to use it

You should:

- put all the sudoku you want to solve in a folder `/data`
- open a terminal
- go to the `/data` parent folder (so you can see the /data folder by typing `ls`)
- run the following command : `docker run --rm -v "$(pwd)/data:/mnt/data" feust/sudoku-solver:2.1 mon_sudoku.txt`

If you have more skills and you want to use it a more convenient way, you can make [aliases](https://dev.ewen.quimerch.com/articles/2-linux-aliases.html) (BEWARE it depends on your computer and config):
  
```bash
echo "alias solve="docker run --rm -v "$(pwd)/data:/mnt/data" feust/sudoku-solver:2.1"" >> ~/.bashrc
bash
solve  <name_of_your_sudoku.txt> <name_of-sudoku2.txt>
```

The results are displayed in your terminal AND in some files created on your /data folder :)

## Handled cases

I tried to handle as many cases I could. You can solve several sudokus in a row, or even all the sudokus in your folder by typing `$(ls)` instead of the names

## Format

The sudokus should be saved in plain text files and look like this :

```bash
...546..9
.2......7
..39....4
9.5....7.
7......2.
....93...
.56..8...
.1..39...
......8.6
```

## From command line instead of files

Just type something like this:

```bash
echo "...546..9
.2......7
..39....4
9.5....7.
7......2.
....93...
.56..8...
.1..39...
......8.6" > data/test && solve test
```

## Dockerfile

```dockerfile
FROM python:3.7

# Creating a special directory
RUN mkdir -p /solver

# Copying only the files we need
COPY main.py parser.py solver.py requirements.txt /solver/

# Working in solver
WORKDIR /solver

# Instaling the librairies (pip comes with the python image)
RUN pip install -r requirements.txt

# Solving ! Need arguments that will be stored in /mnt/data
ENTRYPOINT [ "python3", "main.py"]
```

## Examples

For one file:

![Photo](Capture%20d’écran%20de%202020-06-06%2018-11-23.png)

For several files:

![Photo2](Capture%20d’écran%20de%202020-06-06%2018-12-54.png)
