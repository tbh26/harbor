#! /usr/bin/env python3
###
from proto import simple_pb2


def create_simple() -> simple_pb2.Simple:
    return simple_pb2.Simple(
        id=42,
        is_simple=True,
        name="My python3 simple proto name",
        sample_lists=[3, 4, 5]
    )


def simple_demo():
    simple = create_simple()
    print(f'{simple=}')
    print()


def main():
    print('Hello python proto world!')
    print()
    simple_demo()


if __name__ == '__main__':
    main()
