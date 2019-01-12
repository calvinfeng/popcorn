// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var movie_movie_pb = require('../movie/movie_pb.js');

function serialize_movie_RecRequest(arg) {
  if (!(arg instanceof movie_movie_pb.RecRequest)) {
    throw new Error('Expected argument of type movie.RecRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_movie_RecRequest(buffer_arg) {
  return movie_movie_pb.RecRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_movie_RecResponse(arg) {
  if (!(arg instanceof movie_movie_pb.RecResponse)) {
    throw new Error('Expected argument of type movie.RecResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_movie_RecResponse(buffer_arg) {
  return movie_movie_pb.RecResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var RecommendationService = exports.RecommendationService = {
  fetch: {
    path: '/movie.Recommendation/Fetch',
    requestStream: false,
    responseStream: false,
    requestType: movie_movie_pb.RecRequest,
    responseType: movie_movie_pb.RecResponse,
    requestSerialize: serialize_movie_RecRequest,
    requestDeserialize: deserialize_movie_RecRequest,
    responseSerialize: serialize_movie_RecResponse,
    responseDeserialize: deserialize_movie_RecResponse,
  },
};

exports.RecommendationClient = grpc.makeGenericClientConstructor(RecommendationService);
