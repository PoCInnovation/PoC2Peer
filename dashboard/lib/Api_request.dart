import 'tracktype.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

void initDB() async {
  final _url = new Uri.http("localhost:3000", "/init");
  await http.get(
    _url,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    },
  );
}

void addPeer(String id, String ip) async {
  //http://localhost:3000/addPeer?idpeer=hello3&ippeer=world3
  final _url = new Uri.http("localhost:3000", "/addPeer", {"idpeer": id, "ippeer": ip});
  final http.Response response = await http.get(
    _url,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    },
  );
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

  return hello;
}
