-- CreateTable
CREATE TABLE "Post" (
    "id" SERIAL NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "album" VARCHAR(255) NOT NULL,
    "artist" VARCHAR(255) NOT NULL,
    "genre" VARCHAR(255) NOT NULL,
    "source" VARCHAR(255) NOT NULL,
    "image" VARCHAR(255) NOT NULL,
    "trackNumber" INTEGER NOT NULL,
    "totalTrackCount" INTEGER NOT NULL,
    "duration" INTEGER NOT NULL,
    "site" VARCHAR(255) NOT NULL,

    PRIMARY KEY ("id")
);
