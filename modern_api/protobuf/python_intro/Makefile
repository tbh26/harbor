BIN_NAME = python3
PROTO_DIR = proto

py3 ?= python3.12
bash = /bin/bash
venv = ./venv
req_txt = ./requirements.txt

##

ifeq ($(OS), Windows_NT)
	BIN = ${BIN_NAME}.exe
else
	BIN = ${BIN_NAME}
endif

run:	generate
	${BIN} main.py

generate:
	protoc -I${PROTO_DIR} --python_out=${PROTO_DIR} ${PROTO_DIR}/*.proto

bump:
	sed -i -e 's/==/>=/g' requirements.txt
	./venv/bin/python -m pip install -r requirements.txt --upgrade
	./venv/bin/python -m pip freeze > requirements.txt

clean:
	rm ${PROTO_DIR}/*_pb2.py

venv:
	$(py3) -m venv $(venv)
	$(bash) -c "source $(venv)/bin/activate ; pip3 install --upgrade pip ; pip3 --version "

py_ver: venv
	$(SHELL) -c "source $(venv)/bin/activate ; type python3 "
	$(bash) -c "source $(venv)/bin/activate ; python3 --version "

install_reqs: venv
	$(bash) -c "source $(venv)/bin/activate ; pip3 install --requirement $(req_txt) "

#
.PHONY: py_ver install_reqs clean

