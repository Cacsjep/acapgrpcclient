# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cctv.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\ncctv.proto\"\x1d\n\x0c\x46rameMessage\x12\r\n\x05\x66rame\x18\x01 \x01(\x0c\"\"\n\x0f\x41\x63knowledgement\x12\x0f\n\x07message\x18\x01 \x01(\t2A\n\x0b\x43\x43TVService\x12\x32\n\x0bStreamVideo\x12\r.FrameMessage\x1a\x10.Acknowledgement(\x01\x30\x01\x42(Z&github.com/Cacsjep/acapgrpcclient/cctvb\x06proto3')



_FRAMEMESSAGE = DESCRIPTOR.message_types_by_name['FrameMessage']
_ACKNOWLEDGEMENT = DESCRIPTOR.message_types_by_name['Acknowledgement']
FrameMessage = _reflection.GeneratedProtocolMessageType('FrameMessage', (_message.Message,), {
  'DESCRIPTOR' : _FRAMEMESSAGE,
  '__module__' : 'cctv_pb2'
  # @@protoc_insertion_point(class_scope:FrameMessage)
  })
_sym_db.RegisterMessage(FrameMessage)

Acknowledgement = _reflection.GeneratedProtocolMessageType('Acknowledgement', (_message.Message,), {
  'DESCRIPTOR' : _ACKNOWLEDGEMENT,
  '__module__' : 'cctv_pb2'
  # @@protoc_insertion_point(class_scope:Acknowledgement)
  })
_sym_db.RegisterMessage(Acknowledgement)

_CCTVSERVICE = DESCRIPTOR.services_by_name['CCTVService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z&github.com/Cacsjep/acapgrpcclient/cctv'
  _FRAMEMESSAGE._serialized_start=14
  _FRAMEMESSAGE._serialized_end=43
  _ACKNOWLEDGEMENT._serialized_start=45
  _ACKNOWLEDGEMENT._serialized_end=79
  _CCTVSERVICE._serialized_start=81
  _CCTVSERVICE._serialized_end=146
# @@protoc_insertion_point(module_scope)
