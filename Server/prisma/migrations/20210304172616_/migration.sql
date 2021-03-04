/*
  Warnings:

  - Added the required column `ippeer` to the `Peer` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Peer" ADD COLUMN     "ippeer" VARCHAR(255) NOT NULL;
