import 'package:flutter/material.dart';
import 'tracklist.dart';

class BodyPage extends StatefulWidget {
  BodyPage({Key key}) : super(key: key);

  @override
  _BodyPageState createState() => _BodyPageState();
}

class _BodyPageState extends State<BodyPage> {
  @override
  Widget build(BuildContext context) {
    return (ListView(
      scrollDirection: Axis.vertical,
      children: <Widget>[
        TrackList(),
        PeerList(),
      ],
    ));
  }
}

class PeerList extends StatelessWidget {
  const PeerList({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 400,
      color: Colors.blue,
    );
  }
}
