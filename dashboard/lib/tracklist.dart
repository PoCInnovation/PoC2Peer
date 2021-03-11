import 'dart:developer';

import 'package:flutter/material.dart';

import 'Card.dart';
import 'Api_request.dart';

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

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<List<Map<String, dynamic>>>(
      future: fetchTracks(),
      builder: (context, AsyncSnapshot<List<Map<String, dynamic>>> snapshot) {
        if (snapshot.hasData) {
          final tmp = snapshot.data;
          final tmpLenght = tmp.length;
          return Container(
            height: 700,
            child: GridView.builder(
              padding: EdgeInsets.all(5.0),
              gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 3,
                crossAxisSpacing: 5.0,
                mainAxisSpacing: 12.0,
              ),
              itemCount: tmpLenght,
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
