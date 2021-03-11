import 'dart:developer';

import 'package:flutter/material.dart';
import 'tracktype.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'Card.dart';

class PeerList extends StatefulWidget {
  PeerList({Key key}) : super(key: key);

  @override
  _PeerListState createState() => _PeerListState();
}

class _PeerListState extends State<PeerList> {
  List<Map<String, dynamic>> peerlist = [];
  List<DataRow> tmprow = [];
  Future<List<DataRow>> fetchTracks() async {
    final _url = new Uri.http("localhost:3000", "/peerList");

    final http.Response response = await http.get(
      _url,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
      },
    );

    final value = jsonDecode(response.body);
    for (var item in value) {
      final tmp = {
        "id": item['id'],
        "ippeer": item['ippeer'],
        "idpeer": item['idpeer'],
      };
      peerlist.add(tmp);
    }
    for (var item in peerlist) {
      tmprow.add(DataRow(cells: <DataCell>[
        DataCell(Text(item['idpeer'])),
        DataCell(Text(item['ippeer'])),
      ]));
    }
    inspect(tmprow);
    return (tmprow);
  }

  @override
  void initState() {
    super.initState();
    fetchTracks();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<List<DataRow>>(
      future: fetchTracks(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
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
    // return DataTable(
    //   columns: const <DataColumn>[
    //     DataColumn(
    //       label: Text(
    //         'IdPeer',
    //         style: TextStyle(fontStyle: FontStyle.italic),
    //       ),
    //     ),
    //     DataColumn(
    //       label: Text(
    //         'IP',
    //         style: TextStyle(fontStyle: FontStyle.italic),
    //       ),
    //     ),
    //   ],
    //   rows: tmprow,
    // );
  }
}
