const router = require('express').Router();

router.use('/api', require('./api'));
router.use('/main', require('./main'));

module.exports = router;