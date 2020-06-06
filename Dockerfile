FROM python:3.7

# Creating a special directory
RUN mkdir -p /solver

# Copying only the files we need
COPY main.py parser.py solver.py requirements.txt /solver/

# Working in solver
WORKDIR /solver

# Instaling the librairies (pip comes with the python image)
RUN pip install -r requirements.txt

# Solving ! Need arguments that will be stored in /mnt/s
ENTRYPOINT [ "python3", "main.py"]