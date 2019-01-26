const config = require('config');
const { Pool } = require('pg')

const dbConfig = config.get('Postgres.dbConfig');
const pool = new Pool(dbConfig)

pool.on('error', (err, client) => {
  console.error('Unexpected error on idle client', err)
  process.exit(-1)
})

pool.on('connect', () => {
  console.log('Connected to the db...');
});

exports.pool = pool;
