import express from 'express';
import bodyParser from 'body-parser';
import cookieParser from 'cookie-parser';
import httpStatus from 'http-status-codes';

const PORT = 3000;
const server = express();

server.use(bodyParser.json());
server.use(cookieParser());

server.get('/health', (req, res) => {
  res.sendStatus(httpStatus.OK);
});

server.listen(PORT, () => {
  console.log(`server is listening on ${PORT}`);
});

server.get('/repeat-my-query', (req, res) => {
  if (!req.query.message) {
    res.status(httpStatus.BAD_REQUEST).send('Bad Request');
  } else {
    res.status(httpStatus.OK).send(req.query.message);
  }
});

server.get('/repeat-my-param/:message', (req, res) => {
  res.status(httpStatus.OK).send(req.params.message);
});

server.post('/repeat-my-body', (req, res) => {
  if (!req.body.message) {
    res.status(httpStatus.BAD_REQUEST).send('Bad Request');
  } else {
    res.status(httpStatus.OK).send(req.body.message);
  }
});

server.get('/repeat-my-header', (req, res) => {
  if (!req.header('X-Message')) {
    res.status(httpStatus.BAD_REQUEST).send('Bad Request');
  } else {
    res.status(httpStatus.OK).send(req.header('X-Message'));
  }
});

server.get('/repeat-my-cookie', (req, res) => {
  if (!req.cookies.message) {
    res.status(httpStatus.BAD_REQUEST).send('Bad Request');
  } else {
    res.status(httpStatus.OK).send(req.cookies.message);
  }
});
