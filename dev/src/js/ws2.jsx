const CommentList = props => {
    return <ul className="item-list">
        {props.comments.map(item => <CommentDetail item={item} />)}
    </ul>
};

const CommentDetail = props => {
    const item = props.item;
    return <li>
        <div>{item.Timestamp}</div>
        <div>{item.Content}</div>
    </li>
};

class Comment extends React.Component {
    constructor(props, context, updater) {
        super(props, context, updater)
        // 以下修正点
        this.state = { list: [] }
        this.addToList = this.addToList.bind(this)

        console.log(this.props);
    }

    addToList(todo) {
        this.state.list.push(todo)
        this.setState({ list: this.state.list })
      }

    render() {
        return (
            <div>
                <h1>COMMENT</h1>
                <CommentList comments={this.props.comments} />
            </div>
        )
    }
};

//<TodoForm addToList={this.addToList} />
//<TodoList list={this.state.list} />

fetch('/api/comments').then(res => res.json()).then(data => {
    console.log(data);
    ReactDOM.render(
        <Comment comments={data.Comments} />,
        document.getElementById('board')
    );
});


var websocket = null;
let data = {};

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

