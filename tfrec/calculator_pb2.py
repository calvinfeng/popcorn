# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: calculator.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='calculator.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('Z\007protogo'),
  serialized_pb=_b('\n\x10\x63\x61lculator.proto\"$\n\x05Input\x12\x0c\n\x04left\x18\x01 \x01(\x03\x12\r\n\x05right\x18\x02 \x01(\x03\"\x17\n\x06Output\x12\r\n\x05value\x18\x01 \x01(\x03\x32#\n\tCaculator\x12\x16\n\x03\x41\x64\x64\x12\x06.Input\x1a\x07.OutputB\tZ\x07protogob\x06proto3')
)




_INPUT = _descriptor.Descriptor(
  name='Input',
  full_name='Input',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='left', full_name='Input.left', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='right', full_name='Input.right', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=20,
  serialized_end=56,
)


_OUTPUT = _descriptor.Descriptor(
  name='Output',
  full_name='Output',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='Output.value', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=58,
  serialized_end=81,
)

DESCRIPTOR.message_types_by_name['Input'] = _INPUT
DESCRIPTOR.message_types_by_name['Output'] = _OUTPUT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Input = _reflection.GeneratedProtocolMessageType('Input', (_message.Message,), dict(
  DESCRIPTOR = _INPUT,
  __module__ = 'calculator_pb2'
  # @@protoc_insertion_point(class_scope:Input)
  ))
_sym_db.RegisterMessage(Input)

Output = _reflection.GeneratedProtocolMessageType('Output', (_message.Message,), dict(
  DESCRIPTOR = _OUTPUT,
  __module__ = 'calculator_pb2'
  # @@protoc_insertion_point(class_scope:Output)
  ))
_sym_db.RegisterMessage(Output)


DESCRIPTOR._options = None

_CACULATOR = _descriptor.ServiceDescriptor(
  name='Caculator',
  full_name='Caculator',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=83,
  serialized_end=118,
  methods=[
  _descriptor.MethodDescriptor(
    name='Add',
    full_name='Caculator.Add',
    index=0,
    containing_service=None,
    input_type=_INPUT,
    output_type=_OUTPUT,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CACULATOR)

DESCRIPTOR.services_by_name['Caculator'] = _CACULATOR

# @@protoc_insertion_point(module_scope)
