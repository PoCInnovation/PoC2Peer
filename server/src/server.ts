import express from 'express';
import bodyParser from 'body-parser';
import cookieParser from 'cookie-parser';
import httpStatus from 'http-status-codes';
import { Post, PrismaClient } from '@prisma/client';
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
async function addSongInDB(value: Array<any>) {
  value.forEach(async (element) => {
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
  const tmpvalue: {
    // eslint-disable-next-line max-len
    id: string; title: string; album: string; artist: string; genre: string; source: string; image: string; trackNumber: number; totalTrackCount: number; duration: number; site: string;
  }[] = [];

  prisma.post.findMany().then((data) => {
    data.forEach((element) => {
      tmpvalue.push({
        // eslint-disable-next-line max-len
        id: element.id.toString(), title: element.title, album: element.album, artist: element.artist, genre: element.genre, source: element.source, image: element.image, trackNumber: element.trackNumber, totalTrackCount: element.trackNumber, duration: element.duration, site: element.site,
      });
    });
    res.status(httpStatus.OK).send({ music: tmpvalue });
    // console.log(data);
  });
});

server.post('/addSong', (req, res) => {
  if (!req.body.songs) {
    res.status(httpStatus.BAD_REQUEST).send('Bad Request');
  } else {
    console.log(req.body.songs);
    addSongInDB(req.body.songs);
    res.status(httpStatus.OK).send('great');
  }
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
