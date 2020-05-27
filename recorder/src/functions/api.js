import Axios from 'axios'

export function Upload(file){
  var fd = new FormData();
  fd.set('user', 'user_here')
  fd.set('tenant', '1')
  fd.set('file', file, 'ok.webm')

  Axios.post('http://localhost:8080/upload', fd )
  .then((response) => {
    console.log(response);
  }, (error) => {
    console.log(error);
  })

};

export function Ping(){
  Axios.get('http://localhost:8080/ping')
  .then((response) => {
    console.log(response);
  }, (error) => {
    console.log(error);
  })

};
