import 'dart:async';

import 'package:flutter/material.dart';
import 'tracklist.dart';
import 'PeerTab.dart';
import 'List_action.dart';

class BodyPage extends StatefulWidget {
  BodyPage({Key key}) : super(key: key);

  @override
  _BodyPageState createState() => _BodyPageState();
}

class _BodyPageState extends State<BodyPage> {
  Timer timer;

  void tmp() {
    setState(() {
      print('loop');
    });
  }

  @override
  void initState() {
    super.initState();
    timer = Timer.periodic(Duration(seconds: 5), (Timer t) => {tmp()});
  }

  @override
  Widget build(BuildContext context) {
    return (ListView(
      scrollDirection: Axis.vertical,
      children: <Widget>[
        TrackList(),
        Row(
          children: [
            Container(
              width: 800,
              child: PeerList(),
            ),
            Container(width: 10),
            Flexible(
              child: ActionList(),
              flex: 1,
            ),
          ],
        ),
      ],
    ));
  }
}
