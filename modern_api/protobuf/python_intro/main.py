#! /usr/bin/env python3
###
from proto import simple_pb2, complex_pb2, enumerations_pb2


def create_simple() -> simple_pb2.Simple:
    return simple_pb2.Simple(
        id=42,
        is_simple=True,
        name="My python3 simple proto name",
        sample_lists=[3, 4, 5]
    )


def simple_demo():
    print('simple demo')
    simple = create_simple()
    print(f'{simple=}')
    print()


def create_complex() -> complex_pb2.Complex:
    message = complex_pb2.Complex()
    message.one_dummy.id = 42
    message.one_dummy.name = "First complex name!"
    message.multiple_dummies.add(id=61, name="My complex name 61.")
    message.multiple_dummies.add(id=73, name="My complex name 73.")
    message.multiple_dummies.add(id=84, name="My last complex name?")
    return message


def complex_demo():
    print('complex demo')
    complex_message = create_complex()
    print(f'{complex_message=}')
    print()


def create_enum():
    return enumerations_pb2.Enumeration(
        # eye_color=1
        eye_color=enumerations_pb2.EYE_COLOR_GREEN,
    )


def enum_demo():
    print('enum demo')
    enum_message = create_enum()
    print(f'{enum_message=}')
    print()


def main():
    print('Hello python proto world!')
    print()

    simple_demo()
    print()

    complex_demo()
    print()

    enum_demo()
    print()


if __name__ == '__main__':
    main()
