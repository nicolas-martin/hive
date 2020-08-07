import Axios from 'axios'

export function Upload(file, id){
  var fd = new FormData();
  fd.set('id', id)
  fd.set('file', file, id+'.webm')

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
