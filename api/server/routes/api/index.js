const router = require('express').Router();

router.use('/recommend', require('./recommend'));

module.exports = router;