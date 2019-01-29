const config = require('config');
const router = require('express').Router();
const { schema } = require('../schema/movies');
const validate = require('express-validation');
const { pool } = require('../../db');
const axios = require('axios');

router.param('imdbId', validate(schema.getMovie));

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
        
        res.setHeader('Content-Type', 'application/json');
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
  catch {
    res.setHeader('Content-Type', 'application/json');
    res.status(404).send(`The movie with the given IMDB ID ${imdbId} was not found: ${err}`)
  }
  finally {
    client.release();
  }
});

router.get('/page/', async (req, res) => {
  let client;
  try {
    client = await pool.connect();
  }
  catch {
    res.setHeader('Content-Type', 'application/json');
    res.status(404).send(`The page was not found: ${err}`
  }
  finally {
    client.release();
  }
})

router.post('/rate/:imdbId', async (req, res) => {
  let client;
  const imdbId = req.params.imdbId;
  try {
    client = await pool.connect();
  }
  catch {
    res.setHeader('Content-Type', 'application/json');
    res.status(404).send(`The movie with the given IMDB ID ${imdbId} was not found: ${err}`)
  }
  finally {
    client.release();
  }
})

router.put('/rate/:imdbId', async (req, res) => {
  let client;
  const imdbId = req.params.imdbId;
  try {
    client = await pool.connect();
  }
  catch {
    res.setHeader('Content-Type', 'application/json');
    res.status(404).send(`The movie with the given IMDB ID ${imdbId} was not found: ${err}`
  }
  finally {
    client.release();
  }
})

module.exports = router;
