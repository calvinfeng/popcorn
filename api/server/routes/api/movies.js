const config = require('config');
const router = require('express').Router();
const { schema } = require('../schema/movies');
const validate = require('express-validation');
const { pool } = require('../../db');
const axios = require('axios');
const cache = require('memory-cache');

router.param('imdbId', validate(schema.getMovie));

/**
 * Get a movie detail
 */
router.get('/details/:imdbId', async (req, res) => {
  let client; 
  const imdbId = req.params.imdbId;
  const apiKey = config.get('MovieDB.apiKey');
  try {
    client = await pool.connect();
    // pool.query returns an object with a key called row. The value is an array of requested sql row(s).
    const { rows } = await pool.query('SELECT * From movie_details WHERE imdb_id = $1', [imdbId]);
    let movieDetail = rows[0];

    if ( !movieDetail ) {
      // If record is not found, make a request to The Movie Database and retrieve data.
      try {
        const tmdbRes = await axios.get(`https://api.themoviedb.org/3/movie/${imdbId}?api_key=${apiKey}`);
        const date = new Date();
        await pool.query('INSERT INTO movie_details (imdb_id, created_at, updated_at, detail) VALUES ($1, $2, $3, $4)', [imdbId, date, date, tmdbRes.data]);
        
        res.setHeader('Content-Type', 'application/json');
        res.status(200).send(tmdbRes.data);
      } 
      catch (err) {
        const status = error.response.status;
        const message = `The movie with the given IMDB ID ${imdbId} was not found: ${err}`;
        
        res.status(status).send(message);
      }
    } else {
      res.setHeader('Content-Type', 'application/json');
      res.status(200).send(movieDetail.detail);
    }
  } 
  catch(err) {
    res.status(404).send(`The movie with the given IMDB ID ${imdbId} was not found: ${err}`);    
  }
  finally {
    client.release();
  }
});

/**
 * Get a movie
 */
router.get('/:imdbId', async (req, res) => {
  let client;
  const imdbId = req.params.imdbId;
  try {
    client = await pool.connect();
    const { rows } = await pool.query('SELECT * From movies WHERE imdb_id = $1', [imdbId]);
    let movie = rows[0];
    
    res.setHeader('Content-Type', 'application/json');
    res.status(200).send(movie);
  }
  catch(err) {
    res.status(404).send(`The movie with the given IMDB ID ${imdbId} was not found: ${err}`)
  }
  finally {
    client.release();
  }
});

/**
 * Get movies by page number
 */
// TODO: validation for query?
router.get('/list', async (req, res) => {
  let client;
  const pageNumber = parseInt(req.query.page);

  if (pageNumber < 0 || pageNumber === 0) {
    res.status(404).send(`The page was not found.`)
  }
  const startIdx = 20 * (pageNumber - 1);
  const endIdx = (20 * pageNumber) - 1;
  const movieIdList = cache.get('movieList')
                           .slice(startIdx, endIdx)
                           .join(', ');

  try {
    client = await pool.connect();
    const { rows } = await pool.query(`SELECT * from movies WHERE imdb_id in (${movieIdList})`);
    const movieList = rows;

    res.setHeader('Content-Type', 'application/json');
    res.status(200).send(movieList); // TODO: returns an array do we want that?
  }
  catch(err) {
    res.status(404).send(`The page was not found: ${err}`)
  }
  finally {
    client.release();
  }
})

module.exports = router;
