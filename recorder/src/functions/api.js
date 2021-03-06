import Axios from 'axios'

export function Upload(file, userID, uploadID){
  var fd = new FormData();
  fd.set('uploadID', uploadID)
  fd.set('userID', userID)
  fd.set('file', file,  userID+'.webm')

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
