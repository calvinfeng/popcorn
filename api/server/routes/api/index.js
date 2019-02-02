const router = require('express').Router();

router.use('/recommend', require('./recommend'));
router.use('/movies', require('./movies'));
router.use('/users', require('./users'));

module.exports = router;