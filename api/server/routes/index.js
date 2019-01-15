const router = require('express').Router();

router.use('/api', require('./api'));
router.use('/mock', require('./mock'));

module.exports = router;