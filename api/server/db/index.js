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

// router.post('/api/movie_details/:id', (req, res, next) => {
//   const results = [];
//   // Grab data from http request
//   const data = {text: req.body.text, complete: false};
//   // Get a Postgres client from the connection pool
//   pg.connect(connectionString, (err, client, done) => {
//     // Handle connection errors
//     if(err) {
//       done();
//       console.log(err);
//       return res.status(500).json({success: false, data: err});
//     }
//     // SQL Query > Insert Data
//     client.query('INSERT INTO items(text, complete) values($1, $2)',
//     [data.text, data.complete]);
//     // SQL Query > Select Data
//     const query = client.query('SELECT * FROM items ORDER BY id ASC');
//     // Stream results back one row at a time
//     query.on('row', (row) => {
//       results.push(row);
//     });
//     // After all data is returned, close connection and return results
//     query.on('end', () => {
//       done();
//       return res.json(results);
//     });
//   });
// });

// app.get('/student', (req, res) => {
//   pool.connect((err, client, done) => {
//       const query = 'SELECT * FROM students';
//       client.query(query, (error, result) => {
//         done();
//         if (error) {
//           res.status(400).json({error})
//         } 
//         if(result.rows < '1') {
//           res.status(404).send({
//           status: 'Failed',
//           message: 'No student information found',
//           });
//         } else {
//           res.status(200).send({
//           status: 'Successful',
//           message: 'Students Information retrieved',
//           students: result.rows,
//           });
//         }
//       });
//     });
//   });
  
  
//   app.post('/student', (req, res) => {
//     const data = {
//       name : req.body.studentName,
//       age : req.body.studentAge,
//       classroom : req.body.studentClass,
//       parents : req.body.parentContact,
//       admission : req.body.admissionDate,
//     }
  
//     pool.connect((err, client, done) => {
//       const query = 'INSERT INTO students(student_name,student_age, student_class, parent_contact, admission_date) VALUES($1,$2,$3,$4,$5) RETURNING *';
//       const values = [data.name, data.age, data.classroom, data.parents, data.admission];
  
//       client.query(query, values, (error, result) => {
//         done();
//         if (error) {
//           res.status(400).json({error});
//         }
//         res.status(202).send({
//           status: 'SUccessful',
//           result: result.rows[0],
//         });
//       });
//     });
//   });