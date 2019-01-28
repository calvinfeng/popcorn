const config = require('config');
const cache = require('memory-cache');
const { pool } = require('../db');

exports.timedMovieShuffler = function timedMovieShuffler() {
  setInterval(() => { 
    const previousMovieList = cache.get('movieList');
    const newMovieList = shuffle(previousMovieList);
    cache.put('movieList', newMovieList);
  }, config.get('MovieList.RandomizeTimer'));
}

exports.shuffleMovieList = async function shuffleMovieList() {
  let client;
  try {
    client = await pool.connect();
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

  while (0 !== currentIndex) {
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex -= 1;

    temporaryValue = array[currentIndex];
    array[currentIndex] = array[randomIndex];
    array[randomIndex] = temporaryValue;
  }
  
  return array;
}
