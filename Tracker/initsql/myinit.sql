

create table Post (
  id SERIAL PRIMARY KEY NOT NULL,
  title VARCHAR(255) NOT NULL,
  album VARCHAR(255) NOT NULL,
  autor VARCHAR(255) NOT NULL,
  genre VARCHAR(255) NOT NULL,
  source VARCHAR(255) NOT NULL,
  image VARCHAR(255) NOT NULL,
  trackNumber INTEGER NOT NULL,
  totalTrackCount INTEGER NOT NULL,
  duration INTEGER NOT NULL,
  site VARCHAR(255) NOT NULL
);

INSERT INTO Post (id, title, album, autor, genre, source, image, trackNumber, totalTrackCount,duration, site ) VALUES (1, 'Intro - The Way Of Waking Up (feat. Alan Watts)', 'Wake Up', 'The Kyoto Connection', 'Electronic', 'https://storage.googleapis.com/uamp/The_Kyoto_Connection_-_Wake_Up/01_-_Intro_-_The_Way_Of_Waking_Up_feat_Alan_Watts.mp3', 'https://storage.googleapis.com/uamp/The_Kyoto_Connection_-_Wake_Up/art.jpg',1,13,90, 'http://freemusicarchive.org/music/The_Kyoto_Connection/Wake_Up_1957/');
INSERT INTO Post (id, title, album, autor, genre, source, image, trackNumber, totalTrackCount,duration, site ) VALUES (2, 'Geisha', 'Wake Up', 'The Kyoto Connection', 'Electronic', 'https://storage.googleapis.com/uamp/The_Kyoto_Connection_-_Wake_Up/02_-_Geisha.mp3', 'https://storage.googleapis.com/uamp/The_Kyoto_Connection_-_Wake_Up/art.jpg',2,13,267, 'http://freemusicarchive.org/music/The_Kyoto_Connection/Wake_Up_1957/');