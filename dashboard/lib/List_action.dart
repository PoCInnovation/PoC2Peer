import 'dart:developer';

import 'package:flutter/material.dart';
import 'tracktype.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'Card.dart';

class ActionList extends StatelessWidget {
  const ActionList({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        ElevatedButton(
          onPressed: () => print('hello'),
          child: Text('hello'),
        ),
      ],
    );
  }
}
