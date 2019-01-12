const router = require('express').Router();
const {OAuth2Client} = require('google-auth-library');

const CLIENT_ID = '554659982367-k2ijq5aqf1pnqhat0scq918akbtoib7a.apps.googleusercontent.com';

async function verify(token) {
  const client = new OAuth2Client(CLIENT_ID);
  const ticket = await client.verifyIdToken({
      idToken: token,
      audience: CLIENT_ID,
  });
  
  return ticket.getPayload();;
}

router.get('/:token', (req, res) => {
  verify(req.params.token).then((user) => {
    res.status(200);
    res.send({
      "status": 200,
      "email": user.email,
      "name": user.name,
      "message": `Welcome to Popcorn, ${user.given_name}`
    });
  }).catch((err) => {
    console.error(err);
    res.status(400);
    res.send({
      "status": 400,
      "message": "unexpected error with Google sign-in"
    });
  });
});

module.exports = router;