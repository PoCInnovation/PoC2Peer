import 'dart:developer';

import 'package:dashboard/Api_request.dart';
import 'package:flutter/material.dart';
import 'tracktype.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'Card.dart';

popPeerDialog(BuildContext context) {
  final idpeercontroler = TextEditingController();
  final ippeercontroler = TextEditingController();

  return showDialog(
    barrierDismissible: false,
    context: context,
    builder: (_) => AlertDialog(
      title: Text('Add Peer'),
      content: SingleChildScrollView(
        child: ListBody(
          children: <Widget>[
            TextField(
              controller: idpeercontroler,
              obscureText: true,
              decoration: InputDecoration(
                border: OutlineInputBorder(),
                labelText: 'IdPeer',
              ),
            ),
            Container(height: 10),
            TextField(
              decoration: InputDecoration(
                border: OutlineInputBorder(),
                labelText: 'IpPeer',
              ),
              controller: ippeercontroler,
            )
          ],
        ),
      ),
      actions: <Widget>[
        ElevatedButton(
          child: Text('Add'),
          onPressed: () {
            addPeer(idpeercontroler.text, ippeercontroler.text);
            Navigator.of(context).pop();
          },
        ),
      ],
    ),
  );
}

class ActionList extends StatelessWidget {
  final Function reloadPeer;
  final Function reloadTrack;
  ActionList({Key key, this.reloadPeer, this.reloadTrack}) : super(key: key);

  bool addpeer = false;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Container(
          height: 30,
          width: 100,
          child: ElevatedButton(
            onPressed: () => initDB(),
            child: Text('Init'),
          ),
        ),
        Container(height: 5),
        Container(
          height: 30,
          width: 100,
          child: ElevatedButton(
            onPressed: () => reloadPeer(),
            child: Text('fetchPeer'),
          ),
        ),
        Container(height: 5),
        Container(
          height: 30,
          width: 100,
          child: ElevatedButton(
            onPressed: () => reloadTrack(),
            child: Text('fetchTrack'),
          ),
        ),
        Container(height: 5),
        Container(
          height: 30,
          width: 100,
          child: ElevatedButton(
            onPressed: () => popPeerDialog(context),
            child: Text('AddPeer'),
          ),
        ),
      ],
    );
  }
}
