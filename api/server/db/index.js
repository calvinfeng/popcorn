const { Pool } = require('pg')

const pool = new Pool({
  user: 'popcorn',
  host: 'localhost',
  database: 'popcorn',
  password: 'popcorn',
  port: 5432,
})

pool.on('error', (err, client) => {
  console.error('Unexpected error on idle client', err)
  process.exit(-1)
})

pool.on('connect', () => {
  console.log('Connected to the db...');
});

exports.pool = pool;
