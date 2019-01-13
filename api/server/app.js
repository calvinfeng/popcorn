const services = require('./pb/movie/movie_grpc_pb');
const grpc = require('grpc');

const morgan = require('morgan');
const { createLogger, format, transports } = require('winston');
const { combine, timestamp, label, printf } = format;

const express = require('express');

function newGRPCMiddleware() {
  let hostname = 'localhost';
  if (process.env.BACKEND_HOST) {
    hostname = process.env.BACKEND_HOST
  }

  const gRPCaddress = `${hostname}:8081`;
  const cli = new services.RecommendationClient(gRPCaddress, grpc.credentials.createInsecure());
  
  return (req, res, next) => {
    // Ghetto way of setting context for a request
    req.grpc_client = cli;
    next();
  }
}

function newLogger() {
  const myFormat = printf(info => (
    `${info.timestamp} [${info.label}] ${info.level}: ${info.message}`
  ));

  return createLogger({
    level: 'info',
    format: combine(label({ label: 'api'}), timestamp(), myFormat),
    transports: [
      new transports.Console({colorize: true})
    ]
  });
}

function main() {
  let port = 8080;
  const log = newLogger();
  const app = express(); 
  
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

  app.use(morgan(':date[iso] :http-version :method :url => :response-time ms'));
  app.use(newGRPCMiddleware(), require('./routes'));
  app.use(express.static('public'));
  app.listen(port, () => {
    log.info(`Node API server is serving and listening on port ${port}`)
  })
}

main();


