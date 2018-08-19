const CommentView = props => {
    console.log(props);
    return <li>
        <div>{props.content}</div>
    </li>
};

var websocket = null;

function ws_open() {
    if (websocket == null) {
        websocket = new WebSocket("ws://localhost:8080/echo");
        websocket.onopen = onOpen;
        websocket.onmessage = onMessage;
        websocket.onclose = onClose;
        websocket.onerror = onError;
    }
};

function onOpen(event) {
    console.log(event);
};

function onMessage(event) {
    const content = event.data;
    console.log(content);
    console.log(CommentView);
    /*
    ReactDOM.render(
        <CommentView content={content} />,
        document.getElementById('comments-list'),
    );
    */
};

function onError(event) {
};

function onClose(event) {
};

document.getElementById("ws-connect").onclick = function() {
    ws_open();
};

document.getElementById("comment").onclick = function() {
    const content = document.getElementById("content").value;
    if (content.length > 0 && websocket) {
        websocket.send("" + content);
    }
};

ReactDOM.render(
    <h1>aaa</h1>,
    document.getElementById('comments-list'),
);

