const router = require('express').Router();
const messages = require('../../pb/movie/movie_pb');

router.get('/:user', (req, res) => {
  const cli = req.grpc_client;

  if (cli === undefined) {
    res.status(500);
    res.send('server failed to define gRPC client');
    return;
  }

  const request = new messages.RecRequest();
  request.setUserId(req.params.user);
  
  cli.fetch(request, (err, rpcRes) => {
    if (err !== null) {
      res.status(400);
      res.send(err.details);
      return;
    }
    
    result = [];
    movies = rpcRes.getMoviesList();
    movies.forEach((movie) => {
      result.push(movie.toObject());
    });

    res.status(200);
    res.send(result);
  });
});

module.exports = router;
