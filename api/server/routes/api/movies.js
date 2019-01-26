const config = require('config');
const router = require('express').Router();
const { schema } = require('../schema/movies');
const validate = require('express-validation');
const { pool } = require('../../db');
const axios = require('axios');

router.param('imdbId', validate(schema.getMovie));

router.get('/details/:imdbId', async (req, res) => {
  const client = await pool.connect();
  const imdbId = req.params.imdbId;
  const apiKey = config.get('MovieDB.apiKey');
  try {
    // pool.query returns an object with a key called row. The value is an array of requested sql row(s).
    const { rows } = await pool.query('SELECT * From movie_details WHERE imdb_id = $1', [imdbId]);
    let movieDetail = rows[0];

    if ( !movieDetail ) {
      // If record is not found, make a request to The Movie Database and retrieve data.
      try {
        const tmdbRes = await axios.get(`https://api.themoviedb.org/3/movie/${imdbId}?api_key=${apiKey}`);
        const date = new Date();
        await pool.query('INSERT INTO movie_details (imdb_id, created_at, updated_at, detail) VALUES ($1, $2, $3, $4)', [imdbId, date, date, res.data]);
        res.status(200).send(tmdbRes.data);
      } catch (error) {
        console.log(error)
        const status = error.response.status;
        const message = `There is no movie associated with the provided IMDB ID: ${imdbId}.`;
        res.status(status).send(message);
      }
    }

    res.setHeader('Content-Type', 'application/json');
    res.status(200).send(movieDetail.detail);
  } 
  catch(err) {
    console.log('error ', err);    
  }
  finally {
    client.release();
  }
});

module.exports = router;
