FROM --platform=linux/amd64 python:3.12

RUN mkdir -p /tests/pytest
COPY ./pytest /tests/pytest
COPY ./dotest.sh /tests
RUN pip install -r /tests/pytest/requirements.txt

ENTRYPOINT [ "/tests/dotest.sh" ]