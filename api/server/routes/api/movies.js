const axios = require('axios');
const cache = require('memory-cache');
const config = require('config');
const router = require('express').Router();
const validate = require('express-validation');
const { moviesSchema } = require('../schema/movies');
const { pool } = require('../../db');

/**
 * Get a movie detail
 */
router.get('/details/:imdbId', validate(moviesSchema.validateImdbId), async (req, res) => {
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
        await pool.query('INSERT INTO movie_details (imdb_id, created_at, updated_at, detail) VALUES ($1, $2, $3, $4)', 
                          [imdbId, date, date, tmdbRes.data]);
        
        res.setHeader('Content-Type', 'application/json');
        res.status(200).send(tmdbRes.data);
      } 
      catch(err) {
        const status = err.response.status;
        const message = `The movie with the given IMDB ID ${imdbId} was not found. The Movie DB responded with "${err}"`;
        
        res.status(status).send(message);
      }
    } else {
      res.setHeader('Content-Type', 'application/json');
      res.status(200).send(movieDetail.detail);
    }
  } 
  catch(err) {
    res.status(404).send(`The movie with the given IMDB ID ${imdbId} was not found.`);    
  }
  finally {
    client.release();
  }
});

/**
 * Get a movie
 */
router.get('/title/:imdbId', validate(moviesSchema.validateImdbId), async (req, res) => {
  let client;
  const imdbId = req.params.imdbId;

  try {
    client = await pool.connect();
    const { rows } = await pool.query('SELECT * From movies WHERE imdb_id = $1', [imdbId]);

    let movie = rows[0];
    if (!movie) throw new Error(`The movie with the given IMDB ID ${imdbId} was not found.`);
    
    res.setHeader('Content-Type', 'application/json');
    res.status(200).send(movie);
  }
  catch(err) {
    res.status(404).send(err.message);
  }
  finally {
    client.release();
  }
});

/**
 * Get movies by page number
 */
router.get('/list', validate(moviesSchema.validatePage), async (req, res) => {
  let client;
  const pageNumber = parseInt(req.query.page);
  let movieIdList = cache.get('movieList');

  if (pageNumber <= 0 || pageNumber > Math.round(movieIdList.length / 20)) {
    res.status(404).send(`The page was not found.`);
    return;
  }
  
  const startIdx = 20 * (pageNumber - 1);
  const endIdx = (20 * pageNumber);
  movieIdList = movieIdList.slice(startIdx, endIdx)
                           .join(',');

  try {
    client = await pool.connect();
    const { rows } = await pool.query(`SELECT * from movies WHERE id in (${movieIdList})`);
   
    let movie = rows[0];
    if (!movie) throw new Error(`Unable to fetch movies from DB.`);

    res.setHeader('Content-Type', 'application/json');
    res.status(200).send(rows); 
  }
  catch(err) {
    res.status(404).send(`The page was not found: ${err.message}`)
  }
  finally {
    client.release();
  }
})

module.exports = router;
