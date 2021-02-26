import express from 'express';
import bodyParser from 'body-parser';
import cookieParser from 'cookie-parser';
import httpStatus from 'http-status-codes';
import { PrismaClient } from '@prisma/client';
import * as jsonfile from './init.json';

const PORT = 3000;
const server = express();
const prisma = new PrismaClient();

server.use(bodyParser.json());
server.use(cookieParser());

async function main() {
  const allPost = await prisma.post.findMany();
  if (allPost.length === 0) {
    jsonfile.music.forEach(async (element) => {
      // console.log('loop');
      // console.log(element);
      await prisma.post.create({
        data: {
          title: element.title,
          album: element.album,
          artist: element.artist,
          genre: element.genre,
          source: element.source,
          image: element.image,
          trackNumber: element.trackNumber,
          totalTrackCount: element.totalTrackCount,
          duration: element.duration,
          site: element.site,
        },
      });
    });
  }
  // const allPost2 = await prisma.post.findMany();
  // console.log(allPost2);
}

server.get('/init', (req, res) => {
  main()
    .catch((e) => {
      console.log('erreur');
      res.status(httpStatus.EXPECTATION_FAILED);
      throw e;
    })
    .finally(async () => {
      console.log('fin');
      res.status(httpStatus.OK).send('Done');
      await prisma.$disconnect();
    });
  res.status(httpStatus.OK).send('Done');
  console.log('fin sans catch');

  res.status(httpStatus.OK);
});
server.get('/getSong', (req, res) => {
  prisma.post.findMany().then((data) => {
    res.status(httpStatus.OK).send(data);
    // console.log(data);
  });
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
