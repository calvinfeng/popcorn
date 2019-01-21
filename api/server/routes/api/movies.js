const router = require('express').Router();
const { pool } = require('../../db');

router.get('/details/:imdb_id', async (req, res) => {
  // add validation using joi library
  const client = await pool.connect();
  const imdb_id = req.params.imdb_id;
  try {
    const result = await pool.query('SELECT * From movie_details WHERE imdb_id = $1', [imdb_id]);
    result.rows.forEach((row) => {
      console.log(row.detail.toString());
    })
    res.send(result.rows);
  } 
  catch(err) {
    res.status(404).send(`The movie with the given ID ${id} was not found.`);
  }
  finally {
    client.release();
  }
})

module.exports = router;
