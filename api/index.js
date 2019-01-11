const messages = require('./pb/movie/movie_pb');
const services = require('./pb/movie/movie_grpc_pb');
const grpc = require('grpc');

function main() {
  const client = new services.RecommendationClient('localhost:3000', grpc.credentials.createInsecure());
  const request = new messages.RecRequest();

  request.setUserId(1);

  client.fetch(request, function(err, res) {
    if (err !== null) {
      console.log(err.details);
      return;
    }

    console.log('Movies:', res.getMoviesList());
  });
}

main();


