const { resolveSoa } = require('dns');
var express = require('express');
var axios = require("axios").default;
var app = express();
var fs = require("fs");
// const myJson = require("./filename.json");

app.use(express.json())

var PORT = 3000;




var options = {
    method: 'GET',
    url: 'https://shazam.p.rapidapi.com/songs/list-artist-top-tracks',
    params: {id: '40008598', locale: 'en-US'},
    headers: {
      'x-rapidapi-key': 'd02ba540c2msh02980eca12e1a86p1e0979jsn9dcb64089edd',
      'x-rapidapi-host': 'shazam.p.rapidapi.com'
    }
  };

  app.get('/', function(req, res) {
        res.send(pong);
  });

app.get('/api', function(req, res) {
    axios.request(options).then(function (response) {
        // console.log(response.data['tracks'][0]);
       
        buildjsonfile(response.data);
        
        res.status(200).send(response.data);
    }).catch(function (error) {
        console.error(error);
    });
    
});

app.listen(PORT, function() {
    console.log('Server is running on PORT:',PORT);
});

function buildjsonfile(value) {
  // const myjson = fs.readFile('filename.json', (err, data) => {
  //   if (err) throw err;
  //   return JSON.parse(data);
  // });
  // console.log('json:', myjson);
  // console.log('value:', value);
  // fs.writeFile( "filename.json", JSON.stringify( value ), "utf8", 
  //         function (err) {
  //           if (err) return console.log(err);
  //           console.log('SUCESS');
  //         }
  //       );
  // console.log(value);
  var tmp_value = [];

  value['tracks'].forEach(index => {
    console.log('loooooop');
    // console.log(index);
    tmp_value.push({id: "1", title:"hello", desc:"world"});
  });
  console.log(tmp_value);
}



