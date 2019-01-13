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
    res.locals.grpc_client = cli;
    next();
  }
}

function newLogMiddleware(log) {
  return (req, res, next) => {
    res.locals.log = log;
    next();
  }
}

function main() {
  const app = express(); 
 
  const port = 8080;
  
  // Configure a logger that we will use throughout the application.
  const myFormat = printf(info => (
    `${info.timestamp} [${info.label}] ${info.level}: ${info.message}`
  ));

  const log = createLogger({
    level: 'info',
    format: combine(label({ label: 'api'}), timestamp(), myFormat),
    transports: [
      new transports.Console({colorize: true})
    ]
  });

  // Use redirect to HTTPS logic if the current environment is GCP (Google Cloud Platform.)
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
  app.use(newLogMiddleware(log), newGRPCMiddleware(), require('./routes'));
  app.use(express.static('public'));
  app.listen(port, () => {
    log.info(`Node API server is serving and listening on port ${port}`)
  })
}

main();


