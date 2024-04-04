#! /usr/bin/env python3
##

import time
import multiprocessing as mp
from multiprocessing import Process


def do_work(n: int = 42):
    stamp = int(time.time())
    m = 0
    print(f'Starting work {stamp=} at {n=}, {m=} cpu bound.')
    for _ in range(50_000_000):
        m += 1
    stamp = int(time.time())
    print(f'Finished work {stamp=} at {n=}, {m=} cpu bound (done).')


if __name__ == '__main__':
    print('hello world!   👋')
    mp.set_start_method('spawn')
    for i in range(5):
        p = Process(target=do_work, args=(i, ))
        p.start()
