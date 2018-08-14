# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mhubapi.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='mhubapi.proto',
  package='mhubapi',
  syntax='proto3',
  serialized_pb=_b('\n\rmhubapi.proto\x12\x07mhubapi\"\x1a\n\x18\x41vailableDSKFStreamInput\"?\n\x19\x41vailableDSKFStreamOutput\x12\x0f\n\x07task_id\x18\x01 \x01(\t\x12\x11\n\tdskf_file\x18\x02 \x01(\t\"[\n\x12SetTaskStatusInput\x12\x0f\n\x07task_id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12#\n\x06status\x18\x03 \x01(\x0e\x32\x13.mhubapi.TaskStatus\"\x1f\n\x0e\x44\x65\x65pPingOutput\x12\r\n\x05value\x18\x01 \x01(\t\"\x06\n\x04Void*A\n\nTaskStatus\x12\x0c\n\x08NOT_SENT\x10\x00\x12\x08\n\x04SENT\x10\x01\x12\x0c\n\x08RECEIVED\x10\x02\x12\r\n\tCOMPLETED\x10\x03\x32\xe9\x01\n\x10MicrologProxyHub\x12\x34\n\x08\x44\x65\x65pPing\x12\r.mhubapi.Void\x1a\x17.mhubapi.DeepPingOutput\"\x00\x12=\n\rSetTaskStatus\x12\x1b.mhubapi.SetTaskStatusInput\x1a\r.mhubapi.Void\"\x00\x12`\n\x13\x41vailableDSKFStream\x12!.mhubapi.AvailableDSKFStreamInput\x1a\".mhubapi.AvailableDSKFStreamOutput\"\x00\x30\x01\x42\x1c\xaa\x02\x19SKF.Enlight.API.MProxyHubb\x06proto3')
)

_TASKSTATUS = _descriptor.EnumDescriptor(
  name='TaskStatus',
  full_name='mhubapi.TaskStatus',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NOT_SENT', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SENT', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RECEIVED', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='COMPLETED', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=253,
  serialized_end=318,
)
_sym_db.RegisterEnumDescriptor(_TASKSTATUS)

TaskStatus = enum_type_wrapper.EnumTypeWrapper(_TASKSTATUS)
NOT_SENT = 0
SENT = 1
RECEIVED = 2
COMPLETED = 3



_AVAILABLEDSKFSTREAMINPUT = _descriptor.Descriptor(
  name='AvailableDSKFStreamInput',
  full_name='mhubapi.AvailableDSKFStreamInput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=26,
  serialized_end=52,
)


_AVAILABLEDSKFSTREAMOUTPUT = _descriptor.Descriptor(
  name='AvailableDSKFStreamOutput',
  full_name='mhubapi.AvailableDSKFStreamOutput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='task_id', full_name='mhubapi.AvailableDSKFStreamOutput.task_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dskf_file', full_name='mhubapi.AvailableDSKFStreamOutput.dskf_file', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=54,
  serialized_end=117,
)


_SETTASKSTATUSINPUT = _descriptor.Descriptor(
  name='SetTaskStatusInput',
  full_name='mhubapi.SetTaskStatusInput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='task_id', full_name='mhubapi.SetTaskStatusInput.task_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='mhubapi.SetTaskStatusInput.user_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='mhubapi.SetTaskStatusInput.status', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=119,
  serialized_end=210,
)


_DEEPPINGOUTPUT = _descriptor.Descriptor(
  name='DeepPingOutput',
  full_name='mhubapi.DeepPingOutput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='mhubapi.DeepPingOutput.value', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=212,
  serialized_end=243,
)


_VOID = _descriptor.Descriptor(
  name='Void',
  full_name='mhubapi.Void',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=245,
  serialized_end=251,
)

_SETTASKSTATUSINPUT.fields_by_name['status'].enum_type = _TASKSTATUS
DESCRIPTOR.message_types_by_name['AvailableDSKFStreamInput'] = _AVAILABLEDSKFSTREAMINPUT
DESCRIPTOR.message_types_by_name['AvailableDSKFStreamOutput'] = _AVAILABLEDSKFSTREAMOUTPUT
DESCRIPTOR.message_types_by_name['SetTaskStatusInput'] = _SETTASKSTATUSINPUT
DESCRIPTOR.message_types_by_name['DeepPingOutput'] = _DEEPPINGOUTPUT
DESCRIPTOR.message_types_by_name['Void'] = _VOID
DESCRIPTOR.enum_types_by_name['TaskStatus'] = _TASKSTATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AvailableDSKFStreamInput = _reflection.GeneratedProtocolMessageType('AvailableDSKFStreamInput', (_message.Message,), dict(
  DESCRIPTOR = _AVAILABLEDSKFSTREAMINPUT,
  __module__ = 'mhubapi_pb2'
  # @@protoc_insertion_point(class_scope:mhubapi.AvailableDSKFStreamInput)
  ))
_sym_db.RegisterMessage(AvailableDSKFStreamInput)

AvailableDSKFStreamOutput = _reflection.GeneratedProtocolMessageType('AvailableDSKFStreamOutput', (_message.Message,), dict(
  DESCRIPTOR = _AVAILABLEDSKFSTREAMOUTPUT,
  __module__ = 'mhubapi_pb2'
  # @@protoc_insertion_point(class_scope:mhubapi.AvailableDSKFStreamOutput)
  ))
_sym_db.RegisterMessage(AvailableDSKFStreamOutput)

SetTaskStatusInput = _reflection.GeneratedProtocolMessageType('SetTaskStatusInput', (_message.Message,), dict(
  DESCRIPTOR = _SETTASKSTATUSINPUT,
  __module__ = 'mhubapi_pb2'
  # @@protoc_insertion_point(class_scope:mhubapi.SetTaskStatusInput)
  ))
_sym_db.RegisterMessage(SetTaskStatusInput)

DeepPingOutput = _reflection.GeneratedProtocolMessageType('DeepPingOutput', (_message.Message,), dict(
  DESCRIPTOR = _DEEPPINGOUTPUT,
  __module__ = 'mhubapi_pb2'
  # @@protoc_insertion_point(class_scope:mhubapi.DeepPingOutput)
  ))
_sym_db.RegisterMessage(DeepPingOutput)

Void = _reflection.GeneratedProtocolMessageType('Void', (_message.Message,), dict(
  DESCRIPTOR = _VOID,
  __module__ = 'mhubapi_pb2'
  # @@protoc_insertion_point(class_scope:mhubapi.Void)
  ))
_sym_db.RegisterMessage(Void)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\252\002\031SKF.Enlight.API.MProxyHub'))

_MICROLOGPROXYHUB = _descriptor.ServiceDescriptor(
  name='MicrologProxyHub',
  full_name='mhubapi.MicrologProxyHub',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=321,
  serialized_end=554,
  methods=[
  _descriptor.MethodDescriptor(
    name='DeepPing',
    full_name='mhubapi.MicrologProxyHub.DeepPing',
    index=0,
    containing_service=None,
    input_type=_VOID,
    output_type=_DEEPPINGOUTPUT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetTaskStatus',
    full_name='mhubapi.MicrologProxyHub.SetTaskStatus',
    index=1,
    containing_service=None,
    input_type=_SETTASKSTATUSINPUT,
    output_type=_VOID,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='AvailableDSKFStream',
    full_name='mhubapi.MicrologProxyHub.AvailableDSKFStream',
    index=2,
    containing_service=None,
    input_type=_AVAILABLEDSKFSTREAMINPUT,
    output_type=_AVAILABLEDSKFSTREAMOUTPUT,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_MICROLOGPROXYHUB)

DESCRIPTOR.services_by_name['MicrologProxyHub'] = _MICROLOGPROXYHUB

# @@protoc_insertion_point(module_scope)
