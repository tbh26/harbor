#! /usr/bin/env python3
###
from proto import simple_pb2, complex_pb2, enumerations_pb2, oneofs_pb2, maps_pb2
from google.protobuf.message import Message
from google.protobuf import json_format


def create_simple() -> simple_pb2.Simple:
    return simple_pb2.Simple(
        id=42,
        is_simple=True,
        name='My python3 simple proto name',
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
    message.one_dummy.name = 'First complex name!'
    message.multiple_dummies.add(id=61, name='My complex name 61.')
    message.multiple_dummies.add(id=73, name='My complex name 73.')
    message.multiple_dummies.add(id=84, name='My last complex name?')
    return message


def complex_demo():
    print('complex demo')
    complex_message = create_complex()
    print(f'{complex_message=}')
    print()


def create_enum() -> enumerations_pb2.Enumeration:
    return enumerations_pb2.Enumeration(
        # eye_color=1
        eye_color=enumerations_pb2.EYE_COLOR_GREEN,
    )


def enum_demo():
    print('enum demo')
    enum_message = create_enum()
    print(f'{enum_message=}')
    print()


def create_one_of_number(n: int) -> oneofs_pb2.Result:
    message = oneofs_pb2.Result()
    message.id = 42
    return message


def create_one_of_str(m: str) -> oneofs_pb2.Result:
    message = oneofs_pb2.Result()
    message.message = m
    return message


def one_of_demo():
    print('one of demo')
    one_of_message = create_one_of_number(42)
    print(f'{one_of_message=}')
    one_of_message = create_one_of_str('Hello one of proto world! 👋')
    print(f'{one_of_message=}')
    one_of_message.id = 24
    print(f'{one_of_message=}')
    print()


def create_map() -> maps_pb2.MapExample:
    message = maps_pb2.MapExample()
    message.ids["my id"].id = 42
    message.ids["next"].id = 61
    message.ids["another"].id = 73
    message.ids["last"].id = 84
    return message


def map_demo():
    print('map demo')
    map_message = create_map()
    print(f'{map_message=}')
    print()


def use_file(message: Message, file_path: str) -> None:
    print(f'use_file, {message=} ')
    print(f" - write to file ({file_path})")
    with open(file_path, "wb") as f:
        bytes_as_str = message.SerializeToString()
        f.write(bytes_as_str)

    print(f" - read from file ({file_path})")
    with open(file_path, "rb") as f:
        t = type(message)
        message_read = t().FromString(f.read())
    print(f'use_file, {message_read=} ')


def file_demo():
    print('file demo')
    sm = create_simple()
    use_file(sm, 'simple.data')
    cm = create_complex()
    use_file(cm, 'complex.data')
    print()


def to_json(message: Message) -> str:
    return json_format.MessageToJson(
        message,
        indent=4,
        preserving_proto_field_name=True
    )


def from_json(json_str: str, type) -> Message:
    return json_format.Parse(
        json_str,
        type(),
        ignore_unknown_fields=True
    )


def json_demo():
    print('file demo')

    sm = create_simple()
    json_str = to_json(sm)
    print(f' {json_str=}')
    message = from_json(json_str, simple_pb2.Simple)
    print(f' {message=}')

    cm = create_complex()
    json_str = to_json(cm)
    print(f' {json_str=}')
    message = from_json(json_str, complex_pb2.Complex)
    print(f' {message=}')

    json_str = '{"id": 42, "unknown field": "bla bla"}'
    print(f' {json_str=}')
    message = from_json(json_str, simple_pb2.Simple)
    print(f' {message=}')

    print()


def main():
    print('Hello python proto world! 🌠')
    print()

    simple_demo()
    complex_demo()
    enum_demo()

    one_of_demo()
    map_demo()

    file_demo()

    json_demo()


if __name__ == '__main__':
    main()
