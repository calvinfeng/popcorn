const router = require('express').Router();
const pool = require('../../db');

router.get('/details/:id', async (req, res) => {
  // add validation using joi library
  const client = await pool.connect();
  const id = parseInt(req.params.id);
  try {
    const res = await pool.query('SELECT * FROM movie_details WHERE id = $1', [id]);
    console.log(res);
    res.send(res.rows[0]);
  } 
  catch(err) {
    res.status(404).send(`The movie with the given ID ${id} was not found.`);
  }
  finally {
    client.release();
  }
})