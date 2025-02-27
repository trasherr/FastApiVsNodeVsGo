import express from 'express';
import bodyParser from 'body-parser';
import { Create, Delete, GetAll, Login, Update } from './controller.js';

const app = express();
app.use(bodyParser.json());

const Auth = (req, res, next) => {
  if(req.headers.authorization === 'Bearer test@gmail.com') {
    next();
  } else {
    res.status(401).send({"message": "Unauthorized"});
  }
}

app.get('/', GetAll);
app.post('/user', Auth,  Create);
app.put('/user', Auth, Update);
app.delete('/user', Auth, Delete);
app.get('/user', Auth, Login);

app.listen(3000, () => {
  console.log('Server is running on http://localhost:3000');
});

