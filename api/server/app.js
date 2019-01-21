const express = require('express');
const morgan = require('morgan');
// const userAuthentication = require('./middleware/auth');
const newGRPCMiddleware = require('./middleware/grpc');
const { log, newLogMiddleware } = require('./middleware/logging');

function main() {
  const app = express(); 

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
  app.use(express.json())
  // app.use('/api', userAuthentication);
  app.use(newLogMiddleware(), newGRPCMiddleware(), require('./routes'));

  const port = process.env.PORT | "8080";
  app.listen(port, () => {
    log.info(`Node API server is serving and listening on port ${port}`)
  });
}

main();


