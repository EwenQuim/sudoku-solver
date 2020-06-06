FROM python:latest

COPY sudoku-solver/ /

RUN pip install -r requirements.txt

CMD [ "python", "./main.py" ]