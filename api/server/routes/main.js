const router = require('express').Router();

router.get('/', (req, res) => {
  res.status(200);
  res.send({
    "status": 200,
    "message": `Welcome to Popcorn`
  });
});

module.exports = router;
