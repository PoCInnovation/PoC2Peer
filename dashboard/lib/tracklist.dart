import 'dart:developer';

import 'package:flutter/material.dart';
import 'tracktype.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'Card.dart';

class TrackList extends StatefulWidget {
  TrackList({Key key}) : super(key: key);

  @override
  _TrackListState createState() => _TrackListState();
}

class _TrackListState extends State<TrackList> {
  actualize() {
    setState(() {
      print('refrech');
    });
  }

  Future<List<Map<String, dynamic>>> fetchTracks() async {
    final _url = new Uri.http("localhost:3000", "/getSong");

    final http.Response response = await http.get(
      _url,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
      },
    );

    List<Map<String, dynamic>> hello = [];

    final value = jsonDecode(response.body);
    for (var item in value['music']) {
      final tmp = {
        "id": item['id'],
        "title": item['title'],
        "album": item['album'],
        "artist": item['artist'],
        "genre": item['genre'],
        "source": item['source'],
        "image": item['image'],
        "trackNumber": item['trackNumber'],
        "totalTrackCount": item['totalTrackCount'],
        "duration": item['duration'],
        "site": item['site'],
      };
      hello.add(tmp);
    }
    // inspect(hello);
    return hello;
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<List<Map<String, dynamic>>>(
      future: fetchTracks(),
      builder: (context, AsyncSnapshot<List<Map<String, dynamic>>> snapshot) {
        if (snapshot.hasData) {
          final tmp = snapshot.data;
          final tmp_lenght = tmp.length;
          print(tmp_lenght);
          return Container(
            height: 700,
            child: GridView.builder(
              padding: EdgeInsets.all(5.0),
              gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 3,
                crossAxisSpacing: 5.0,
                mainAxisSpacing: 12.0,
              ),
              itemCount: tmp_lenght,
              scrollDirection: Axis.horizontal,
              itemBuilder: (context, int index) {
                return CardItem(
                  id: snapshot.data[index]['id'],
                  title: snapshot.data[index]['title'],
                  url: snapshot.data[index]['image'],
                  refresh: () => actualize(),
                );
              },
            ),
          );
        } else if (snapshot.hasError) {
          return Text("${snapshot.error}");
        }
        return CircularProgressIndicator();
      },
    );
  }
}
