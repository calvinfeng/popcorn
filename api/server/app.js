const services = require('./pb/movie/movie_grpc_pb');
const grpc = require('grpc');

const morgan = require('morgan');
const { createLogger, format, transports } = require('winston');
const { combine, timestamp, label, printf } = format;

const express = require('express');
const userAuthentication = require('./middleware/auth');

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

function newLogMiddleware(log) {
  return (req, res, next) => {
    res.locals.log = log;
    next();
  }
}

function main() {
  const app = express(); 
 
  let port = "8080";
  if (process.env.PORT) {
    port = process.env.PORT;
  }
  
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
  app.use(express.static('public'));
  
  app.use('/api', userAuthentication);
  app.use(newLogMiddleware(log), newGRPCMiddleware(), require('./routes'));
  app.listen(port, () => {
    log.info(`Node API server is serving and listening on port ${port}`)
  })
}

main();


