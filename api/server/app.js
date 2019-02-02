const express = require('express');
const morgan = require('morgan');
const { timedMovieShuffler, shuffleMovieList } = require('./services/movies');
const userAuthentication = require('./middleware/auth');
const newGRPCMiddleware = require('./middleware/grpc');
const { log, newLogMiddleware } = require('./middleware/logging');
const ev = require('express-validation');

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

  try {
    shuffleMovieList();
    timedMovieShuffler();
  } catch(err) {
    process.exit(1);
  }

  app.use(morgan(':date[iso] :http-version :method :url => :response-time ms'));
  app.use(express.static('public'));
  app.use(express.json())
  // app.use('/api', userAuthentication);
  app.use(newLogMiddleware(), newGRPCMiddleware(), require('./routes'));
  app.use(function (err, req, res, next) {
    // specific for validation errors
    if (err instanceof ev.ValidationError) return res.status(err.status).json(err);

    if (process.env.NODE_ENV !== 'production') {
      return res.status(500).send(err.stack);
    } else {
      return res.status(500);
    }
  });

  const port = process.env.PORT | "8080";
  app.listen(port, () => {
    log.info(`Node API server is serving and listening on port ${port}`)
  });
}

main();

