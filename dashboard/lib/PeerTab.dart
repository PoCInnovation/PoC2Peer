import 'dart:developer';

import 'package:flutter/material.dart';
import 'tracktype.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'Card.dart';
import 'Api_request.dart';

class PeerList extends StatefulWidget {
  PeerList({Key key}) : super(key: key);

  @override
  _PeerListState createState() => _PeerListState();
}

class _PeerListState extends State<PeerList> {
  @override
  Widget build(BuildContext context) {
    return FutureBuilder<List<DataRow>>(
      future: fetchPeer(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return DataTable(
            columns: const <DataColumn>[
              DataColumn(
                label: Text(
                  'delete',
                  style: TextStyle(fontStyle: FontStyle.italic),
                ),
              ),
              DataColumn(
                label: Text(
                  'IdPeer',
                  style: TextStyle(fontStyle: FontStyle.italic),
                ),
              ),
              DataColumn(
                label: Text(
                  'IP',
                  style: TextStyle(fontStyle: FontStyle.italic),
                ),
              ),
            ],
            rows: snapshot.data,
          );
        }
        if (snapshot.hasError) {
          return (Text('erreur'));
        }
        return DataTable(
          columns: const <DataColumn>[
            DataColumn(
              label: Text(
                'IdPeer',
                style: TextStyle(fontStyle: FontStyle.italic),
              ),
            ),
            DataColumn(
              label: Text(
                'IP',
                style: TextStyle(fontStyle: FontStyle.italic),
              ),
            ),
          ],
          rows: [],
        );
      },
    );
  }
}
