// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var movie_movie_pb = require('../movie/movie_pb.js');

function serialize_movie_RecommendRequest(arg) {
  if (!(arg instanceof movie_movie_pb.RecommendRequest)) {
    throw new Error('Expected argument of type movie.RecommendRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_movie_RecommendRequest(buffer_arg) {
  return movie_movie_pb.RecommendRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_movie_RecommendResponse(arg) {
  if (!(arg instanceof movie_movie_pb.RecommendResponse)) {
    throw new Error('Expected argument of type movie.RecommendResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_movie_RecommendResponse(buffer_arg) {
  return movie_movie_pb.RecommendResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_movie_UpdateRequest(arg) {
  if (!(arg instanceof movie_movie_pb.UpdateRequest)) {
    throw new Error('Expected argument of type movie.UpdateRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_movie_UpdateRequest(buffer_arg) {
  return movie_movie_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_movie_UpdateResponse(arg) {
  if (!(arg instanceof movie_movie_pb.UpdateResponse)) {
    throw new Error('Expected argument of type movie.UpdateResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_movie_UpdateResponse(buffer_arg) {
  return movie_movie_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var RecommendationService = exports.RecommendationService = {
  fetch: {
    path: '/movie.Recommendation/Fetch',
    requestStream: false,
    responseStream: false,
    requestType: movie_movie_pb.RecommendRequest,
    responseType: movie_movie_pb.RecommendResponse,
    requestSerialize: serialize_movie_RecommendRequest,
    requestDeserialize: deserialize_movie_RecommendRequest,
    responseSerialize: serialize_movie_RecommendResponse,
    responseDeserialize: deserialize_movie_RecommendResponse,
  },
  updateUserPreference: {
    path: '/movie.Recommendation/UpdateUserPreference',
    requestStream: false,
    responseStream: false,
    requestType: movie_movie_pb.UpdateRequest,
    responseType: movie_movie_pb.UpdateResponse,
    requestSerialize: serialize_movie_UpdateRequest,
    requestDeserialize: deserialize_movie_UpdateRequest,
    responseSerialize: serialize_movie_UpdateResponse,
    responseDeserialize: deserialize_movie_UpdateResponse,
  },
};

exports.RecommendationClient = grpc.makeGenericClientConstructor(RecommendationService);
