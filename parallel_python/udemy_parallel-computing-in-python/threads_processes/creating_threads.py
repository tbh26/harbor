#! /usr/bin/env python3
##

import time
from threading import Thread


def do_work(n: int = 42):
    stamp = int(time.time())
    print(f'Starting work {stamp=} at {n}')
    time.sleep(1)
    stamp = int(time.time())
    print(f'Finished work {stamp=} at {n}')


def main():
    print('hello world!   👋')

    for n in range(5):
        do_work(n)

    for n in range(5):
        t = Thread(target=do_work, args=(n, ))
        t.start()


if __name__ == '__main__':
    main()
