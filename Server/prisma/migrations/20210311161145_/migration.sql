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

-- CreateTable
CREATE TABLE "Peer" (
    "id" SERIAL NOT NULL,
    "idpeer" VARCHAR(255) NOT NULL,
    "ippeer" VARCHAR(255) NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Peer.idpeer_unique" ON "Peer"("idpeer");
