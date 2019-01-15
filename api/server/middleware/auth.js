const {OAuth2Client} = require('google-auth-library');

// OAuth Client ID
const CLIENT_ID = '159011759683-01llm5cirgtboge88g73342bl9nn1ihb.apps.googleusercontent.com';

async function verify(token) {
  const client = new OAuth2Client(CLIENT_ID);
  const ticket = await client.verifyIdToken({
      idToken: token,
      audience: CLIENT_ID,
  });
  
  return ticket.getPayload();;
}

function userAuthentication(req, res, next) {
  console.log('Authenticating user...')
  
  if (!req.get('token')) {
    res.status(400);
    res.send({
      "message": "token is not found in header"
    });
    return;
  }

  verify(req.get('token')).then((user) => {
    console.log('yayyy! passed!')
    next()
  }).catch((err) => {
    res.status(400);
    res.send({
      "status": 400,
      "message": "unexpected error with Google sign-in"
    });
  });
}

module.exports = userAuthentication;