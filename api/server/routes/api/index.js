const router = require('express').Router();

router.use('/recommend', require('./recommend'));
router.use('/movies', require('./movies'));
router.use('/movies', require('./user'));

module.exports = router;