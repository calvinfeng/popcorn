// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var recommendation_pb = require('./recommendation_pb.js');

function serialize_RecommendRequest(arg) {
  if (!(arg instanceof recommendation_pb.RecommendRequest)) {
    throw new Error('Expected argument of type RecommendRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_RecommendRequest(buffer_arg) {
  return recommendation_pb.RecommendRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_RecommendResponse(arg) {
  if (!(arg instanceof recommendation_pb.RecommendResponse)) {
    throw new Error('Expected argument of type RecommendResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_RecommendResponse(buffer_arg) {
  return recommendation_pb.RecommendResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateRequest(arg) {
  if (!(arg instanceof recommendation_pb.UpdateRequest)) {
    throw new Error('Expected argument of type UpdateRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_UpdateRequest(buffer_arg) {
  return recommendation_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateResponse(arg) {
  if (!(arg instanceof recommendation_pb.UpdateResponse)) {
    throw new Error('Expected argument of type UpdateResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_UpdateResponse(buffer_arg) {
  return recommendation_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var RecommendationService = exports.RecommendationService = {
  fetch: {
    path: '/Recommendation/Fetch',
    requestStream: false,
    responseStream: false,
    requestType: recommendation_pb.RecommendRequest,
    responseType: recommendation_pb.RecommendResponse,
    requestSerialize: serialize_RecommendRequest,
    requestDeserialize: deserialize_RecommendRequest,
    responseSerialize: serialize_RecommendResponse,
    responseDeserialize: deserialize_RecommendResponse,
  },
  updateUserPreference: {
    path: '/Recommendation/UpdateUserPreference',
    requestStream: false,
    responseStream: false,
    requestType: recommendation_pb.UpdateRequest,
    responseType: recommendation_pb.UpdateResponse,
    requestSerialize: serialize_UpdateRequest,
    requestDeserialize: deserialize_UpdateRequest,
    responseSerialize: serialize_UpdateResponse,
    responseDeserialize: deserialize_UpdateResponse,
  },
};

exports.RecommendationClient = grpc.makeGenericClientConstructor(RecommendationService);
