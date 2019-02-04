const router = require('express').Router();
const validate = require('express-validation');
const { pool } = require('../../db');
const { usersSchema } = require('../schema/users');

router.param('email', validate(usersSchema.validateEmail));

/**
 * Get all rated movies by user
 */
router.get('/:email/ratings', async (req, res) => {
  let client;
  const userEmail = req.params.email;

  try {
    client = await pool.connect();
    const { rows } = await pool.query('SELECT * from ratings WHERE user_email = $1', [userEmail]);

    const ratings = rows[0];
    if (!ratings) throw new Error(`Unable to fetch ${userEmail} ratings from DB.`);

    res.setHeader('Content-Type', 'application/json');
    res.status(200).send(rows);
  }
  catch(err) {
    res.status(404).send(err.message);
  }
  finally {
    client.release();
  }
})

/**
 * Post a movie rating
 */
router.post('/:email/ratings', validate(usersSchema.validateRatings), async (req, res) => {
  let client;
  const userEmail = req.params.email;
  const movieId = req.body.movieId; 
  const rating = req.body.rating;
  const date = new Date();

  try {
    client = await pool.connect();
    await pool.query('INSERT INTO ratings (movie_id, user_email, value, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)',
                      [movieId, userEmail, rating, date, date]);

    res.status(200).send(`Movie rating was successfully saved into the DB.`);
  }
  catch(err) {
    res.status(404).send(`Unable to save rating to db: ${err}`)
  }
  finally {
    client.release();
  }
})

/**
 * Edit a movie rating
 */
router.put('/:email/ratings', validate(usersSchema.validateRatings), async (req, res) => {
  let client;
  const userEmail = req.params.email;
  const movieId = req.body.movieId;
  const rating = req.body.rating;
  const date = new Date();

  try {
    client = await pool.connect();
    await pool.query('UPDATE ratings SET value = $1, updated_at = $2 WHERE movie_id = $3 AND user_email = $4',
                      [rating, date, movieId, userEmail]);

    res.status(200).send(`Movie rating was successfully updated into the DB.`);
  }
  catch(err) {
    res.status(404).send(`Unable to update ratings to db: ${err}`)
  }
  finally {
    client.release();
  }
})

module.exports = router;
