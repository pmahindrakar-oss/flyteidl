# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/plugins/waitable.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import execution_pb2 as flyteidl_dot_core_dot_execution__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/plugins/waitable.proto',
  package='flyteidl.plugins',
  syntax='proto3',
  serialized_pb=_b('\n\x1f\x66lyteidl/plugins/waitable.proto\x12\x10\x66lyteidl.plugins\x1a\x1d\x66lyteidl/core/execution.proto\x1a\x1e\x66lyteidl/core/identifier.proto\"\x96\x01\n\x08Waitable\x12>\n\nwf_exec_id\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifier\x12\x35\n\x05phase\x18\x02 \x01(\x0e\x32&.flyteidl.core.WorkflowExecution.Phase\x12\x13\n\x0bworkflow_id\x18\x03 \x01(\tB5Z3github.com/lyft/flyteidl/gen/pb-go/flyteidl/pluginsb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_execution__pb2.DESCRIPTOR,flyteidl_dot_core_dot_identifier__pb2.DESCRIPTOR,])




_WAITABLE = _descriptor.Descriptor(
  name='Waitable',
  full_name='flyteidl.plugins.Waitable',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='wf_exec_id', full_name='flyteidl.plugins.Waitable.wf_exec_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phase', full_name='flyteidl.plugins.Waitable.phase', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='workflow_id', full_name='flyteidl.plugins.Waitable.workflow_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=117,
  serialized_end=267,
)

_WAITABLE.fields_by_name['wf_exec_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._WORKFLOWEXECUTIONIDENTIFIER
_WAITABLE.fields_by_name['phase'].enum_type = flyteidl_dot_core_dot_execution__pb2._WORKFLOWEXECUTION_PHASE
DESCRIPTOR.message_types_by_name['Waitable'] = _WAITABLE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Waitable = _reflection.GeneratedProtocolMessageType('Waitable', (_message.Message,), dict(
  DESCRIPTOR = _WAITABLE,
  __module__ = 'flyteidl.plugins.waitable_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.plugins.Waitable)
  ))
_sym_db.RegisterMessage(Waitable)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z3github.com/lyft/flyteidl/gen/pb-go/flyteidl/plugins'))
# @@protoc_insertion_point(module_scope)
