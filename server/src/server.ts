import express from 'express';
import bodyParser from 'body-parser';
import cookieParser from 'cookie-parser';
import httpStatus from 'http-status-codes';
import { PrismaClient } from '@prisma/client';

const PORT = 3000;
const server = express();
const prisma = new PrismaClient();

server.use(bodyParser.json());
server.use(cookieParser());

server.get('/health', (req, res) => {
  prisma.post.findMany().then((data) => {
    res.status(httpStatus.OK).send(data);
  });
  // console.log(allPost);
});

async function main() {
  await prisma.post.create({
    data: {
      title: 'Intro - The Way Of Waking Up (feat. Alan Watts)',
      album: 'Wake Up',
      autor: 'The Kyoto Connection',
      genre: 'Electronic',
      source: 'https://storage.googleapis.com/uamp/The_Kyoto_Connection_-_Wake_Up/01_-_Intro_-_The_Way_Of_Waking_Up_feat_Alan_Watts.mp3',
      image: 'https://storage.googleapis.com/uamp/The_Kyoto_Connection_-_Wake_Up/art.jpg',
      trackNumber: 1,
      totalTrackCount: 13,
      duration: 90,
      site: 'http://freemusicarchive.org/music/The_Kyoto_Connection/Wake_Up_1957/',
    },
  });
  const allPost = await prisma.post.findMany();
  console.log(allPost);
}

server.listen(PORT, () => {
  main()
    .catch((e) => {
      throw e;
    })
    .finally(async () => {
      await prisma.$disconnect();
    });
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
