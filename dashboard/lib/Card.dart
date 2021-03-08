import 'dart:developer';

import 'package:flutter/material.dart';
import 'tracktype.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class CardItem extends StatelessWidget {
  final String id;
  final String url;
  final Function refresh;

  final String title;
  const CardItem({Key key, this.url, this.title, this.id, this.refresh}) : super(key: key);

  void deleteSong() async {
    final _url = new Uri.http("localhost:3000", "/deletSong", {'id': id});

    final http.Response response = await http.post(
      _url,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
      },
    );
    print(response);
  }

  @override
  Widget build(BuildContext context) {
    return Card(
        child: Container(
      decoration: BoxDecoration(
        image: DecorationImage(
          image: NetworkImage(url),
          fit: BoxFit.cover,
        ),
      ),
      padding: EdgeInsets.all(10.0),
      child: Column(
        children: [
          Text(
            title,
            style: TextStyle(
              fontWeight: FontWeight.bold,
            ),
          ),
          Expanded(
            child: Align(
              alignment: FractionalOffset.bottomRight,
              child: IconButton(
                icon: Icon(Icons.delete, color: Colors.red),
                tooltip: 'delete track',
                onPressed: () {
                  deleteSong();
                  refresh();
                },
              ),
            ),
          ),
        ],
      ),
      // child: Column(
      //   mainAxisSize: MainAxisSize.min,
      //   mainAxisAlignment: MainAxisAlignment.end,
      //   children: [
      //     Text(
      //       title,
      //       style: TextStyle(
      //         fontWeight: FontWeight.bold,
      //       ),
      //     ),
      //     Container(
      //       child: IconButton(
      //         icon: Icon(Icons.volume_up),
      //         tooltip: 'Increase volume by 10',
      //         onPressed: () {
      //           print('clic');
      //         },
      //       ),
      //     ),
      //   ],
      // )),
    ));
  }
}

// Container(
//                   color: Colors.grey[300],
//                   padding: EdgeInsets.all(10.0),
//                   child: Center(
//                     child: Text("GridView $index"),
//                   ),
//                 );

// "id": "1",
// "title": "Intro - The Way Of Waking Up (feat. Alan Watts)",
// "album": "Wake Up",
// "artist": "The Kyoto Connection",
// "genre": "Electronic",
// "source": "https://storage.googleapis.com/uamp/The_Kyoto_Connection_-_Wake_Up/01_-_Intro_-_The_Way_Of_Waking_Up_feat_Alan_Watts.mp3",
// "image": "https://storage.googleapis.com/uamp/The_Kyoto_Connection_-_Wake_Up/art.jpg",
// "trackNumber": 1,
// "totalTrackCount": 1,
// "duration": 90,
// "site": "http://freemusicarchive.org/music/The_Kyoto_Connection/Wake_Up_1957/"
