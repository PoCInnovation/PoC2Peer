/*
  Warnings:

  - The migration will add a unique constraint covering the columns `[idpeer]` on the table `Peer`. If there are existing duplicate values, the migration will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "Peer.idpeer_unique" ON "Peer"("idpeer");
