// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var calculator_pb = require('./calculator_pb.js');

function serialize_Input(arg) {
  if (!(arg instanceof calculator_pb.Input)) {
    throw new Error('Expected argument of type Input');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_Input(buffer_arg) {
  return calculator_pb.Input.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_Output(arg) {
  if (!(arg instanceof calculator_pb.Output)) {
    throw new Error('Expected argument of type Output');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_Output(buffer_arg) {
  return calculator_pb.Output.deserializeBinary(new Uint8Array(buffer_arg));
}


// Calculator is a dummy server that tests the implementation of gRPC server in Python.
var CaculatorService = exports.CaculatorService = {
  add: {
    path: '/Caculator/Add',
    requestStream: false,
    responseStream: false,
    requestType: calculator_pb.Input,
    responseType: calculator_pb.Output,
    requestSerialize: serialize_Input,
    requestDeserialize: deserialize_Input,
    responseSerialize: serialize_Output,
    responseDeserialize: deserialize_Output,
  },
};

exports.CaculatorClient = grpc.makeGenericClientConstructor(CaculatorService);
