const express = require('express');
const morgan = require('morgan');
// const userAuthentication = require('./middleware/auth');
const newGRPCMiddleware = require('./middleware/grpc');
const { log, newLogMiddleware } = require('./middleware/logging');
const cache = require('memory-cache');
const { pool } = require('./db');

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
  } catch(err) {
    process.exit(1);
  }

  setInterval(() => { 
    const previousMovieList = cache.get('movieList');
    const newMovieList = shuffle(previousMovieList);
    cache.put('movieList', newMovieList);
  }, 10000);

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

async function shuffleMovieList() {
  const client = await pool.connect();
  try {
    const result = await pool.query('SELECT id FROM movies');

    const movieIds = [];
    for (let idx in result.rows){
      movieIds.push(result.rows[idx].id);
    }

    cache.put('movieList', shuffle(movieIds));
  } catch(err) {
    console.log(`Failed to get movies from database: ${err}`);
  } finally {
    client.release();
  }
}

function shuffle(array) {
  let currentIndex = array.length;
  let temporaryValue = 0;
  let randomIndex = 0;

  // While there remain elements to shuffle...
  while (0 !== currentIndex) {

    // Pick a remaining element...
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex -= 1;

    // And swap it with the current element.
    temporaryValue = array[currentIndex];
    array[currentIndex] = array[randomIndex];
    array[randomIndex] = temporaryValue;
  }

  return array;
}
