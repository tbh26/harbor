
####

ARG ?=
#SHELL := /bin/bash
bash = /bin/bash

py3 ?= python3.11
venv = ./venv
req_txt = ./requirements.txt

####


default:
	@printf "\nPlease use 'make' here with an 'argument'.   ( current target == '$@' ) \n\n"
	@printf "    examples:  \n\n"
	@printf "             $$ make venv \n\n"
	@printf "             $$ make venv py3=python3.10 \n\n"
	@printf "             $$ make pip_list \n\n"
	@printf "             $$ make pip_add ARG=requests \n\n"
	@printf "             $$ make install_reqs \n\n"
	@printf "             $$ make store_reqs \n\n"
	@printf "             $$ make clean   # danger! \n\n"

venv:
	$(py3) -m venv $(venv)
	$(bash) -c "source $(venv)/bin/activate ; pip3 install --upgrade pip ; pip3 --version "

py_ver: venv
	$(SHELL) -c "source $(venv)/bin/activate ; type python3 "
	$(bash) -c "source $(venv)/bin/activate ; python3 --version "

pip_list: venv
	$(bash) -c "source $(venv)/bin/activate ; pip3 freeze "

pip_add:
	@printf "\n    arg: $(ARG) \n"
	$(bash) -c "source $(venv)/bin/activate ; pip3 install $(ARG) "

install_reqs: venv
	$(bash) -c "source $(venv)/bin/activate ; pip3 install --requirement $(req_txt) "


store_reqs: venv
	$(bash) -c "source $(venv)/bin/activate ; pip3 freeze > $(req_txt) "

clean:
	rm -f $(req_txt)
	rm -rf $(venv)


#
.PHONY: py_ver pip_list pip_add install_reqs store_reqs default


