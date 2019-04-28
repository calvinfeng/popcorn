const services = require('../protojs/recommendation_grpc_pb');
const grpc = require('grpc');

function newGRPCMiddleware() {
  let gRPCaddress = 'localhost:8081';

  // Check if we are running in a dockerized environment.
  // Use port 8081 locally.
  if (process.env.GRPC_HOSTNAME) {
    gRPCaddress = `${process.env.GRPC_HOSTNAME}:8081`;
  }

  // Check if we are on Google cloud platform.
  // Use port 80 for deployment.
  if (process.env.GCP) {
    gRPCaddress = `${process.env.GPC_GRPC_HOSTNAME}:80`;
  }

  const cli = new services.RecommendationClient(gRPCaddress, grpc.credentials.createInsecure());
  
  return (req, res, next) => {
    res.locals.grpc_client = cli;
    next();
  }
}

module.exports = newGRPCMiddleware;