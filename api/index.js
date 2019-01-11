const messages = require('./pb/movie/movie_pb');
const services = require('./pb/movie/movie_grpc_pb');
const grpc = require('grpc');
const express = require('express');


function newRecommendationFetchHandler(cli) {
  return (req, res) => {
    const request = new messages.RecRequest();
    request.setUserId(1);

    cli.fetch(request, (err, rpcRes) => {
      if (err !== null) {
        res.status(400);
        res.send(err.details);
        return;
      }
      
      result = [];
      movies = rpcRes.getMoviesList();
      movies.forEach((movie) => {
        result.push(movie.toObject());
      });

      res.status(200);
      res.send(result);
    });
  };
}


function main() {
  const client = new services.RecommendationClient('localhost:8080', grpc.credentials.createInsecure());
  const app = express();
  const port = 8000;

  app.get('/', newRecommendationFetchHandler(client));
  app.listen(port, () => console.log(`gRPC client app listening on port ${port}!`))
}

main();


