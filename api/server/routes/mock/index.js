const router = require('express').Router();

router.use('/grpc', require('./grpc'));

module.exports = router;