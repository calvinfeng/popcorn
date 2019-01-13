const services = require('./pb/movie/movie_grpc_pb');
const grpc = require('grpc');
const express = require('express');

function newGRPCMiddleware(cli) {
  return (req, res, next) => {
    req.grpc_client = cli;
    next();
  }
}

function main() {
  let hostname = 'recommender';
  if (process.env.BACKEND_HOST) {
    hostname = process.env.BACKEND_HOST
  }

  // TODO: Move gRPC client into respective route.
  const address = `${hostname}:8081`;
  const client = new services.RecommendationClient(address, grpc.credentials.createInsecure());
  const app = express(); 
  const port = 8080;
  
  // Use redirect to HTTPS logic if the current environment is GCP
  if (process.env.GCP) {
    app.all('/', (req, res, next) => {
      if (req.get('X-Forwarded-Proto') === 'https') {
        next();
      } else {
        res.redirect('https://' + req.headers.host + req.url);
      }
    });
  }

  app.use(newGRPCMiddleware(client), require('./routes'));
  app.use(express.static('public'));
  app.listen(port, () => {
    console.log(`Node API server is serving and listening on port ${port}`)
  })
}

main();


