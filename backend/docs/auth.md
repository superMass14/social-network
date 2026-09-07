// import { useState, useEffect } from 'react';
// // const url = 'ws://localhost:8080/ws'
// const useSocket = (url) => {

// // const [data, setData] = useState(null);
// const [socket, setSocket] = useState(null);
// useEffect(() => {
// const ws = new WebSocket(url);

// ws.onopen = (event) =>{
// }
// ws.onmessage = (event) => {
// const msg = JSON.parse(event.data);
// setData(msg);
// }
// ws.onerror = (error) => {
// }
// ws.onclose = (event) =>{
// }
// setSocket(ws);
// return () => {
// ws.close();
// }
// }, [url]);

// const sendMessage = (message) =>{
// if( socket && socket.readyState === WebSocket.OPEN ){
// socket.send(JSON.stringify(message));
// }
// }
// return {sendMessage};
// }

// export default useSocket;
