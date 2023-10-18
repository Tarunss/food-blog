//Our API to connecto to our backend
var socket = new WebSocket("ws://localhost:8080/ws");
let connect = () => {
    console.log("Attempting connection...")
    socket.onopen = () => {
        console.log("Successfully Connected!");
    };
    socket.onmessage = msg => {
        console.log(msg);
    };
    socket.onclose = event => {
        console.log("socket has closed connection...", event);
    };
    socket.onerrro = error => {
        console.log("Ran into error:", error);
    };
};
let sendMsg = msg => {
    console.log("Sending message:", msg);
    socket.send(msg);
};

export { connect, sendMsg };